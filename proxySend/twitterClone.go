package proxySend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func GetAllTweetsPath() string {
	return "http://localhost/api/tweets"
}

func GetOneTweetPath(c *fiber.Ctx) error {
	url := "http://localhost/api/tweets/" + c.Params("id")
	if err := proxy.Do(c, url); err != nil {
		return err
	}
	// Remove Server header from response
	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}
