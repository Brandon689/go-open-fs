package main

import (
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	FilePath string `json:"filePath"`
}

func main() {
	app := fiber.New()

	app.Post("/openFile", func(c *fiber.Ctx) error {
		request := new(RequestBody)

		if err := c.BodyParser(request); err != nil {
			log.Println("Failed to parse body:", err)
			// Send the error as the response
			return c.Status(400).SendString("Invalid request body")
		}

		err := OpenFileWithMpvfilePath(request.FilePath)
		if err != nil {
			log.Println("Failed to open file:", err)
			// Send the error as the response
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString("File successfully opened")
	})

	log.Fatal(app.Listen(":3000"))
}

func OpenFileWithMpvfilePath string) error {
	cmd := exec.Command("mpv.exe", filePath)
	return cmd.Run()
}