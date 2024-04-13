package lead

import (
	"simple_crm_go/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
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

func NewLead(c *fiber.Ctx) error {
	db := database.DbConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendFile("Error Occured Important")
	}

	db.Create(&lead)
	c.JSON(lead)
}
func DeleteLeads(c *fiber.Ctx) {}
