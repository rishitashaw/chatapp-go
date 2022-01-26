package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
    app := fiber.New()

	//enable browser to request frontend
	app.Use(cors.New())

	pusherClient := pusher.Client{
    AppID: "1338217",
    Key: "2b07a193e1ada1684bb4",
    Secret: "8ea7afeb125dd5254438",
    Cluster: "ap2",
    Secure: true,
  }
    app.Post("/api/messages", func (c *fiber.Ctx) error {
		var data map[string]string
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		pusherClient.Trigger("chat", "message", data)
        return c.JSON([]string{})
    })
	
    log.Fatal(app.Listen(":8000"))
}