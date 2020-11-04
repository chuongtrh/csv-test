package main

import (
	"encoding/csv"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/", func(c *fiber.Ctx) error {

		log.Println(c.Get("Content-Type"))
		log.Println(c.Get("Content-Disposition"))

		// Get first file from form field "document":
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		// log.Println("File:", file.Filename)
		// log.Println("Size:", file.Size)
		// log.Println("Header:", file.Header)

		temp, err := file.Open()
		if err != nil {
			return err
		}
		reader := csv.NewReader(temp)

		if _, err := reader.Read(); err != nil {
			panic(err)
		}

		records, err := reader.ReadAll()
		if err != nil {
			return err
		}
		return c.SendString(fmt.Sprintf("len: %d", len(records)))
	})

	log.Fatal(app.Listen("localhost:3000"))

}
