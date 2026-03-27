package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// DocsView renders the documentation page
func DocsView(c *fiber.Ctx) error {
	path := c.Params("*")
	if path == "" {
		path = c.Params("path") // Fallback if using :path in some other context
	}
	if path == "" {
		path = "INDEX" // Default page
	}

	// 1. Sanitize Path (Prevent Directory Traversal)
	// Allow alphanumeric, underscores, hyphens, and slashes for subdirectories
	// But strictly prevent '..'
	if strings.Contains(path, "..") {
		return c.Status(400).SendString("Invalid path")
	}

	// 2. Construct absolute path to docs/
	// Check for DOCS_PATH environment variable
	docsRoot := os.Getenv("DOCS_PATH")
	if docsRoot == "" {
		// Fallback to local ./docs directory
		cwd, _ := os.Getwd()
		docsRoot = filepath.Join(cwd, "docs")
	}

	// Determine if path has an explicit extension
	hasExplicitExt := strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".html")
	var docPath string
	var ext string

	if hasExplicitExt {
		docPath = filepath.Join(docsRoot, path)
		if strings.HasSuffix(path, ".md") {
			ext = ".md"
		} else {
			ext = ".html"
		}
	} else {
		docPath = filepath.Join(docsRoot, path+".md")
		ext = ".md"
	}

	// 3. Check if file exists
	if _, err := os.Stat(docPath); os.IsNotExist(err) && !hasExplicitExt {
		docPath = filepath.Join(docsRoot, path+".html")
		ext = ".html"
		if _, err := os.Stat(docPath); os.IsNotExist(err) {
			// Try checking without extension or index for folders
			docPathIndex := filepath.Join(docsRoot, path, "README.md")
			ext = ".md"
			if _, err := os.Stat(docPathIndex); err == nil {
				docPath = docPathIndex
			} else {
				docPathIndexHtml := filepath.Join(docsRoot, path, "index.html")
				ext = ".html"
				if _, err := os.Stat(docPathIndexHtml); err == nil {
					docPath = docPathIndexHtml
				} else {
					// Force HTML content type for 404 error page to avoid quirks mode
					c.Set("Content-Type", "text/html; charset=utf-8")
					c.Status(404)
					return utils.RenderInertia(c, "Docs", fiber.Map{
						"content":     "# 404 Not Found\n\nThe requested documentation page could not be found.",
						"currentPath": path,
						"displayName": "Not Found",
						"extension":   ".md",
					})
				}
			}
		}
	} else if err != nil && hasExplicitExt {
		// If explicit extension was provided but file not found, 404 immediately
		c.Set("Content-Type", "text/html; charset=utf-8")
		c.Status(404)
		return utils.RenderInertia(c, "Docs", fiber.Map{
			"content":     "# 404 Not Found\n\nThe requested documentation page could not be found.",
			"currentPath": path,
			"displayName": "Not Found",
			"extension":   ".md",
		})
	}

	// 4. Read File Content
	contentBytes, err := os.ReadFile(docPath)
	if err != nil {
		return c.Status(500).SendString("Error reading documentation file")
	}

	content := string(contentBytes)

	// 5. Determine Display Name
	displayName := filepath.Base(path)
	displayName = strings.ReplaceAll(displayName, "_", " ")
	caser := cases.Title(language.English)
	displayName = caser.String(strings.ToLower(displayName))

	props := fiber.Map{
		"content":     content,
		"currentPath": path,
		"displayName": displayName,
		"extension":   ext,
	}

	return utils.RenderInertia(c, "Docs/index", props)
}

// SaveDoc handles saving the edited documentation page
func SaveDoc(c *fiber.Ctx) error {
	var req struct {
		Path      string `json:"path"`
		Content   string `json:"content"`
		Extension string `json:"extension"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.Path == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Path is required"})
	}

	// 1. Sanitize Path (Prevent Directory Traversal)
	if strings.Contains(req.Path, "..") {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid path"})
	}

	// 2. Construct absolute path to docs/
	docsRoot := os.Getenv("DOCS_PATH")
	if docsRoot == "" {
		cwd, _ := os.Getwd()
		docsRoot = filepath.Join(cwd, "docs")
	}

	ext := req.Extension
	if ext == "" {
		ext = ".md"
	}

	docPath := filepath.Join(docsRoot, req.Path+ext)

	// 3. Check if file exists (or at least the directory exists)
	// We want to allow creating new files or editing existing ones
	dir := filepath.Dir(docPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Attempt to create the directory if it doesn't exist
		if err := os.MkdirAll(dir, 0755); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create directory"})
		}
	}

	// 4. Write File Content
	if err := os.WriteFile(docPath, []byte(req.Content), 0644); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Documentation saved successfully",
	})
}

// DocItem represents a documentation file or folder
type DocItem struct {
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Path      string    `json:"path"`
	Category  string    `json:"category"`
	Type      string    `json:"type"` // "file" or "folder"
	Extension string    `json:"extension"`
	Children  []DocItem `json:"children,omitempty"`
}

// GetDocsNavigation returns a JSON structure of the docs directory
func GetDocsNavigation(c *fiber.Ctx) error {
	docsRoot := os.Getenv("DOCS_PATH")
	if docsRoot == "" {
		cwd, _ := os.Getwd()
		docsRoot = filepath.Join(cwd, "docs")
	}

	allDocs := make([]DocItem, 0)

	err := filepath.Walk(docsRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip errors
		}

		// Skip hidden files/folders and specific directories
		if strings.HasPrefix(info.Name(), ".") || info.Name() == "node_modules" {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Only process .md or .html files
		isMd := !info.IsDir() && strings.HasSuffix(info.Name(), ".md")
		isHtml := !info.IsDir() && strings.HasSuffix(info.Name(), ".html")

		if isMd || isHtml {
			relPath, _ := filepath.Rel(docsRoot, path)

			// Determine Category based on parent folder
			// By default, if file is in root docs folder, use the root folder name
			rootFolderName := filepath.Base(docsRoot)
			caserTitleForRoot := cases.Title(language.English)
			category := caserTitleForRoot.String(rootFolderName)
			
			dirName := filepath.Dir(relPath)
			if dirName != "." {
				// Convert "planning/architecture" to "Planning / Architecture"
				parts := strings.Split(dirName, string(os.PathSeparator))
				caser := cases.Title(language.English)
				for i, p := range parts {
					parts[i] = caser.String(strings.ReplaceAll(p, "_", " "))
				}
				category = strings.Join(parts, " / ")
			}

			// Determine Extension and logical Name (path without extension)
			ext := ".md"
			if isHtml {
				ext = ".html"
			}
			logicalName := strings.TrimSuffix(relPath, ext)
			// Normalize slashes for web API
			logicalName = strings.ReplaceAll(logicalName, "\\", "/")

			// Create a readable Title from filename
			title := info.Name()
			title = strings.TrimSuffix(title, ext)
			title = strings.ReplaceAll(title, "_", " ")
			caserTitle := cases.Title(language.English)
			title = caserTitle.String(title)

			allDocs = append(allDocs, DocItem{
				Name:      logicalName,
				Title:     title,
				Path:      logicalName,
				Category:  category,
				Type:      "file",
				Extension: ext,
			})
		}
		return nil
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to read documentation structure"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    allDocs,
	})
}

// CreateFolder handles the creation of a new subdirectory in the docs folder
func CreateFolder(c *fiber.Ctx) error {
	var req struct {
		ParentPath string `json:"parentPath"`
		FolderName string `json:"folderName"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.FolderName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Folder name is required"})
	}

	docsRoot := os.Getenv("DOCS_PATH")
	if docsRoot == "" {
		cwd, _ := os.Getwd()
		docsRoot = filepath.Join(cwd, "docs")
	}

	// Sanitize folder name
	cleanName := strings.ReplaceAll(req.FolderName, "/", "")
	cleanName = strings.ReplaceAll(cleanName, "\\", "")
	cleanName = strings.ReplaceAll(cleanName, "..", "")
	cleanName = strings.TrimSpace(cleanName)

	if cleanName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid folder name"})
	}

	// Validate parent path
	if strings.Contains(req.ParentPath, "..") {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid parent path"})
	}

	folderPath := filepath.Join(docsRoot, req.ParentPath, cleanName)

	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create folder"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Folder created successfully",
	})
}

// CreateFile handles the creation of a new empty .md file
func CreateFile(c *fiber.Ctx) error {
	var req struct {
		FolderPath string `json:"folderPath"`
		FileName   string `json:"fileName"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.FileName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "File name is required"})
	}

	docsRoot := os.Getenv("DOCS_PATH")
	if docsRoot == "" {
		cwd, _ := os.Getwd()
		docsRoot = filepath.Join(cwd, "docs")
	}

	// Sanitize file name
	cleanName := strings.ReplaceAll(req.FileName, "/", "")
	cleanName = strings.ReplaceAll(cleanName, "\\", "")
	cleanName = strings.ReplaceAll(cleanName, "..", "")
	cleanName = strings.TrimSpace(cleanName)
	if !strings.HasSuffix(cleanName, ".md") && !strings.HasSuffix(cleanName, ".html") {
		cleanName += ".md" // default to .md
	}

	// Validate folder path
	if strings.Contains(req.FolderPath, "..") {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid folder path"})
	}

	filePath := filepath.Join(docsRoot, req.FolderPath, cleanName)

	// Check if already exists
	if _, err := os.Stat(filePath); err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "File already exists"})
	}

	initialContent := ""
	if strings.HasSuffix(cleanName, ".md") {
		initialContent = "# " + strings.TrimSuffix(cleanName, ".md") + "\n"
	} else if strings.HasSuffix(cleanName, ".html") {
		initialContent = "<h1>" + strings.TrimSuffix(cleanName, ".html") + "</h1>\n"
	}

	if err := os.WriteFile(filePath, []byte(initialContent), 0644); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create file"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "File created successfully",
	})
}
