package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"

	"github.com/BDavid57/go-api-proxy/proxySend"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/tweets/get_all", proxy.Forward(proxySend.GetAllTweetsPath()))
	app.Get("/tweets/get_one/:id", proxySend.GetOneTweetPath)

	port := 3001
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
