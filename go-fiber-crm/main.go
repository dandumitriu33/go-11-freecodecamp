package main

import (
	"go-fiber-crm/database"
	"go-fiber-crm/lead"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		log.Panic("failed to connect to DB %s", err)
	}
	log.Println("Connection to DB open")
	database.DBConn.AutoMigrate(&lead.Lead{})
	log.Println("DB migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(3000)
	
}
