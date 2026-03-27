package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/tarakreasi/taraNote_go/internal/routes"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Validate required environment (fail fast if port is missing)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize View Engine
	engine := html.New("./views", ".html")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "taraDocs",
		Views:   engine,
		// Do not expose internal error details to the client
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// For Inertia / API requests, return JSON with a safe message
			if c.Get("X-Inertia") == "true" || c.XHR() {
				msg := "An internal error occurred"
				if code < 500 {
					msg = err.Error()
				}
				return c.Status(code).JSON(fiber.Map{"error": msg})
			}

			// For browser requests, return a sanitized HTML page
			c.Set("Content-Type", "text/html; charset=utf-8")
			safeMsg := "An internal error occurred. Please try again."
			if code == 404 {
				safeMsg = "Page not found."
			} else if code == 400 {
				safeMsg = "Bad request."
			}
			return c.Status(code).SendString(
				"<!DOCTYPE html><html><head><title>Error</title></head><body>" +
					"<h1>Error</h1>" +
					"<p>" + safeMsg + "</p></body></html>",
			)
		},
	})

	// CORS: restrict to same-origin in production
	appEnv := os.Getenv("APP_ENV")
	corsOrigins := "http://localhost:3000,http://localhost:5173"
	if appEnv == "production" {
		corsOrigins = os.Getenv("APP_URL")
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: corsOrigins,
		AllowMethods: "GET,POST",
		AllowHeaders: "Origin, Content-Type, Accept, X-Inertia, X-Inertia-Version, X-XSRF-TOKEN",
	}))

	// Rate limiter for write endpoints
	writeLimiter := limiter.New(limiter.Config{
		Max:        30,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests. Please slow down.",
			})
		},
	})

	// Custom Method Override Middleware
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodPost {
			method := c.FormValue("_method")
			if method == "" {
				method = c.Get("X-HTTP-Method-Override")
			}
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				c.Method(method)
			}
		}
		return c.Next()
	})

	// Static Assets
	app.Static("/resources", "./resources")
	app.Static("/public", "./public")
	app.Static("/images", "./public/images")

	// Routes (rate limiter applied on write routes inside routes package)
	routes.SetupWeb(app, writeLimiter)

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
