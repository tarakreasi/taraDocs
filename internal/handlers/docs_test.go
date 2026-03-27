package handlers

import (
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// setupTestApp returns a Fiber app with docs routes pointing to the given docs directory.
func setupTestApp(t *testing.T, docsPath string) *fiber.App {
	t.Helper()
	os.Setenv("DOCS_PATH", docsPath)
	t.Cleanup(func() { os.Unsetenv("DOCS_PATH") })

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		// Use JSON-only error handler in tests (no view engine available)
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})
	app.Get("/docs/*", DocsView)
	app.Get("/docs", DocsView)
	app.Post("/docs/save", SaveDoc)
	app.Post("/docs/api/create-file", CreateFile)
	app.Get("/docs/api/navigation", GetDocsNavigation)
	return app
}

// createTempDocs sets up a temporary docs directory with sample files.
func createTempDocs(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	// Create an md file
	os.WriteFile(filepath.Join(dir, "hello.md"), []byte("# Hello\n\nWorld."), 0644)

	// Create an html file
	os.WriteFile(filepath.Join(dir, "page.html"), []byte("<h1>Hello HTML</h1>"), 0644)

	// Create a subfolder with a file
	subDir := filepath.Join(dir, "guides")
	os.MkdirAll(subDir, 0755)
	os.WriteFile(filepath.Join(subDir, "start.md"), []byte("# Start\n\nGuide."), 0644)

	return dir
}

// ─────────────────────────────────────────────
// isPathSafe Tests
// ─────────────────────────────────────────────

func TestIsPathSafe_ValidPath(t *testing.T) {
	root := "/docs"
	target := "/docs/file.md"
	if !isPathSafe(root, target) {
		t.Error("expected valid path to be safe")
	}
}

func TestIsPathSafe_TraversalAttack(t *testing.T) {
	root := "/docs"
	target := "/docs/../etc/passwd"
	if isPathSafe(root, target) {
		t.Error("expected traversal path to be blocked")
	}
}

func TestIsPathSafe_ExactRoot(t *testing.T) {
	root := "/docs"
	if !isPathSafe(root, root) {
		t.Error("expected exact root to be safe")
	}
}

// ─────────────────────────────────────────────
// DocsView Tests
// ─────────────────────────────────────────────

func TestDocsView_ExistingMarkdownFile(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	req := httptest.NewRequest("GET", "/docs/hello", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestDocsView_ExistingHtmlFile(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	req := httptest.NewRequest("GET", "/docs/page.html", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestDocsView_FileNotFound(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	req := httptest.NewRequest("GET", "/docs/nonexistent", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}

func TestDocsView_PathTraversalBlocked(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	req := httptest.NewRequest("GET", "/docs/../etc/passwd", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestDocsView_SubfolderFile(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	req := httptest.NewRequest("GET", "/docs/guides/start", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

// ─────────────────────────────────────────────
// GetDocsNavigation Tests
// ─────────────────────────────────────────────

func TestGetDocsNavigation_ReturnsFiles(t *testing.T) {
	dir := createTempDocs(t)

	// Request navigation endpoint using a minimal fiber app
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Get("/docs/api/navigation", GetDocsNavigation)

	os.Setenv("DOCS_PATH", dir)
	t.Cleanup(func() { os.Unsetenv("DOCS_PATH") })

	req := httptest.NewRequest("GET", "/docs/api/navigation", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGetDocsNavigation_EmptyDir(t *testing.T) {
	dir := t.TempDir()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Get("/docs/api/navigation", GetDocsNavigation)

	os.Setenv("DOCS_PATH", dir)
	t.Cleanup(func() { os.Unsetenv("DOCS_PATH") })

	req := httptest.NewRequest("GET", "/docs/api/navigation", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

// ─────────────────────────────────────────────
// SaveDoc Tests
// ─────────────────────────────────────────────

func TestSaveDoc_EmptyPath(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	body := strings.NewReader(`{"path":"","content":"test","extension":".md"}`)
	req := httptest.NewRequest("POST", "/docs/save", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestSaveDoc_PathTraversalBlocked(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	body := strings.NewReader(`{"path":"../../etc/passwd","content":"evil","extension":".md"}`)
	req := httptest.NewRequest("POST", "/docs/save", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestSaveDoc_ValidMarkdown(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	body := strings.NewReader(`{"path":"hello","content":"# Updated","extension":".md"}`)
	req := httptest.NewRequest("POST", "/docs/save", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

// ─────────────────────────────────────────────
// CreateFile Tests
// ─────────────────────────────────────────────

func TestCreateFile_DefaultsToMd(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	body := strings.NewReader(`{"folderPath":"","fileName":"newfile"}`)
	req := httptest.NewRequest("POST", "/docs/api/create-file", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if _, err := os.Stat(filepath.Join(dir, "newfile.md")); os.IsNotExist(err) {
		t.Error("expected newfile.md to be created")
	}
}

func TestCreateFile_HtmlExtension(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	body := strings.NewReader(`{"folderPath":"","fileName":"newpage.html"}`)
	req := httptest.NewRequest("POST", "/docs/api/create-file", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if _, err := os.Stat(filepath.Join(dir, "newpage.html")); os.IsNotExist(err) {
		t.Error("expected newpage.html to be created")
	}
}

func TestCreateFile_AlreadyExists(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	body := strings.NewReader(`{"folderPath":"","fileName":"hello.md"}`)
	req := httptest.NewRequest("POST", "/docs/api/create-file", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestCreateFile_TraversalBlocked(t *testing.T) {
	dir := createTempDocs(t)
	app := setupTestApp(t, dir)

	body := strings.NewReader(`{"folderPath":"../../","fileName":"evil.md"}`)
	req := httptest.NewRequest("POST", "/docs/api/create-file", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}
