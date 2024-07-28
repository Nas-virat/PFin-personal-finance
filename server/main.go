package main

import (
	"os"
	"strconv"
	"time"

	"github.com/Nas-virat/PFin-personal-finance/constant"
	"github.com/Nas-virat/PFin-personal-finance/db"
	"github.com/Nas-virat/PFin-personal-finance/log"
	"github.com/Nas-virat/PFin-personal-finance/router"
	"github.com/Nas-virat/PFin-personal-finance/utils"
	"github.com/gin-gonic/gin"
	docs "github.com/Nas-virat/PFin-personal-finance/docs"
   	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r := gin.Default()
	

	// app := fiber.New()
	// app.Use(logger.New())
	// app.Use(cors.New(
	// 	cors.Config{
	// 		AllowOrigins: "*",
	// 		AllowHeaders: "Origin, Content-Type, Accept",
	// 	},
	// ))
	// logg.Info("Initialized fiber")

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("root")
	// })

	// app.Get("/swagger/*", swagger.HandlerDefault)
	v1 := r.Group("/api")
	docs.SwaggerInfo.BasePath = "/api"
	
	router.SetupAccountRoutes(v1,db)
	router.SetupTransactionRoute(v1, db)
	router.SetupBalanceRoutes(v1, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}

func initTimeZone() {
	ict, err := time.LoadLocation("Local")
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
