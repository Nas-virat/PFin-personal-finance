package main

import (
	"os"
	"strconv"
	"time"

	"github.com/Nas-virat/PFin-personal-finance/constant"
	"github.com/Nas-virat/PFin-personal-finance/db"
	_ "github.com/Nas-virat/PFin-personal-finance/docs"
	"github.com/Nas-virat/PFin-personal-finance/log"
	"github.com/Nas-virat/PFin-personal-finance/router"
	"github.com/Nas-virat/PFin-personal-finance/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func main() {

	logg := log.NewLogger()

	// Initialize Timezone
	initTimeZone()
	logg.Info("Initialized timezone")

	db := initDB()
	logg.Info("Initialized database")

	// Migrate the schema
	utils.Migration(db)
	logg.Info("Migrated the schema")

	if constant.IsDevelopment {
		logg.Warn("Running in development mode")
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))
	logg.Info("Initialized fiber")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupAccountRoutes(app, db)
	router.SetupTransactionRoute(app, db)
	router.SetupBalanceRoutes(app, db)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":8000")
}

func initTimeZone() {
	ict, err := time.LoadLocation("")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDB() *gorm.DB {
	// Connect to the database.
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	if err != nil {
		panic(err)
	}

	dbConfig := db.Config{
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     port,
		Name:     os.Getenv("DATABASE_NAME"),
		TimeZone: os.Getenv("DATABASE_TIMEZONE"),
		DisableTLS: true,
	}
	db := db.ConnectDB(dbConfig)

	return db
}
