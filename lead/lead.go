package lead

import (
	"simple_crm_go/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) {
	db := database.DbConn
	var lead []Lead
	db.Find(&lead)
	c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DbConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).SendFile("Error Occured Important")
	}

	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLeads(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn

	var lead Lead
	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(500).SendFile("No lead Found")
		return
	}

	db.Delete(&lead)
	c.SendString("Lead successfully Deleted")
}
