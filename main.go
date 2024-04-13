package main

import (
	"fmt"
	"simple_crm_go/database"
	"simple_crm_go/lead"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLeads)
}

func main() {
	fmt.Println("main server file....")
	database.Initialize()
	app := fiber.New()
	setupRoutes(app)
	// start my server
	app.Listen("3000")
}
