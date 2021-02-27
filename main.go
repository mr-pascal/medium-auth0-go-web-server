package main

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"log"
)

func main() {
	// Create new Go Fiber application web server
	app := fiber.New()

	// Use YOUR 'HS256' signing key here
	signingKey := "YOUR_SIGNING_KEY"

	// A unrestricted 'ping' route
	unrestrictedPingPath := "/ping/unrestricted"

	// List of unrestricted paths
	unrestrictedPaths := map[string]bool{
		unrestrictedPingPath: true,
	}

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:    []byte(signingKey),
		SigningMethod: "HS256",
		Filter: func(ctx *fiber.Ctx) bool {
			currentPath := ctx.Path()
			v, _ := unrestrictedPaths[currentPath]
			// If 'true' the JWT middleware is ignored
			return v
		},
	}))

	// Add route handlers
	app.Get("/ping", MyHandler)
	app.Get(unrestrictedPingPath, MyHandler)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func MyHandler(c *fiber.Ctx) error {
	return c.SendString("Works")
}
