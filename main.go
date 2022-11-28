package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/tweets/all", proxy.Forward("http://localhost/api/tweets"))

	port := 3001
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
