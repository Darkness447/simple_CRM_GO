package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Post(NewLeads)
	app.Delete(DeleteLeads)
}

func main() {
	fmt.Println("main server file....")
	app := fiber.New()
	setupRoutes(app)
}
