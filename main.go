package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()

	app.Static("/", "./")
	app.Static("/about", "./about.html")
		app.Listen(":3000")
fmt.Print("testing")
	}