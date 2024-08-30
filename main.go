package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	api := fiber.New()

	api.Get("/", func(c *fiber.Ctx) error {

		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println("Error, Cannot Get Host Name")
			return c.Status(500).JSON(err)
		}

		ip := c.IP()

		fmt.Println("Send Request From Path / Success")
		return c.JSON(fiber.Map{
			"Hostname":      hostname,
			"IP":            ip,
			"Message":       "Test CICD From Go",
			"Version Image": "v0.1",
		})

	})

	api.Get("/p1", func(c *fiber.Ctx) error {

		fmt.Println("Send Request From Path /p1 Success")
		return c.JSON(fiber.Map{
			"Message":       "Path /p1",
			"Version Image": "v0.1",
		})

	})

	api.Get("/status", func(c *fiber.Ctx) error {
		fmt.Println("OK")
		return c.SendStatus(200)
	})

	api.Listen(":3000")
}
