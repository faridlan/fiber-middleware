package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Middlware(ctx *fiber.Ctx) error {

	fmt.Println("before middleware")
	err := ctx.Next()
	fmt.Println("after middleware")
	return err

}

func main() {

	app := fiber.New()

	app.Use(Middlware)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Helllo")
	})

	app.Listen("localhost:8080")

}
