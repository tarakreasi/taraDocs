package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/tarakreasi/taraNote_go/internal/handlers"
)

// SetupWeb configures all application routes.
// writeLimiter is applied to all write (POST) endpoints.
func SetupWeb(app *fiber.App, writeLimiter fiber.Handler) {
	// Redirect Root to Docs
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/init")
	}).Name("home")

	// Read-only API (no limiter needed)
	app.Get("/docs/api/navigation", handlers.GetDocsNavigation).Name("docs.navigation")
	app.Get("/docs", handlers.DocsView).Name("docs.index")
	app.Get("/docs/*", handlers.DocsView).Name("docs.show")

	// Write endpoints - rate limited
	app.Post("/docs/save", writeLimiter, handlers.SaveDoc).Name("docs.save")
	app.Post("/docs/api/create-folder", writeLimiter, handlers.CreateFolder).Name("docs.create-folder")
	app.Post("/docs/api/create-file", writeLimiter, handlers.CreateFile).Name("docs.create-file")
}

// NoOpLimiter returns a pass-through handler for testing
func NoOpLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error { return c.Next() }
}

// compile-time assertion: limiter.New returns a fiber.Handler
var _ fiber.Handler = limiter.New()
