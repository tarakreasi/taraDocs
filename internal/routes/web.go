package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/handlers"
)

// SetupWeb routes
func SetupWeb(app *fiber.App) {
	// Redirect Root to Docs
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/init")
	}).Name("home")

	// Docs Routes
	app.Post("/docs/save", handlers.SaveDoc).Name("docs.save")
	app.Post("/docs/api/create-folder", handlers.CreateFolder).Name("docs.create-folder")
	app.Post("/docs/api/create-file", handlers.CreateFile).Name("docs.create-file")
	app.Get("/docs/api/navigation", handlers.GetDocsNavigation).Name("docs.navigation")
	app.Get("/docs", handlers.DocsView).Name("docs.index")
	app.Get("/docs/*", handlers.DocsView).Name("docs.show")
}
