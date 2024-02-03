package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Nas-virat/PFin-personal-finance/constant"
	"github.com/Nas-virat/PFin-personal-finance/db"
	"github.com/Nas-virat/PFin-personal-finance/log"
	"github.com/Nas-virat/PFin-personal-finance/router"
	"github.com/Nas-virat/PFin-personal-finance/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "github.com/Nas-virat/PFin-personal-finance/docs"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {

	logg := log.NewLogger()

	// Initialize Timezone
	initTimeZone()
	logg.Info("Initialized timezone")

	// Initialize config
	initConfig()
	logg.Info("Initialized config")

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

	router.SetupAccountRoutes(app,db)
	router.SetupTransactionRoute(app,db)
	router.SetupBalanceRoutes(app,db)


	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	   })


	app.Listen(":8000")
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initConfig(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") 
	viper.AutomaticEnv() 
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))


	err := viper.ReadInConfig()
	if err != nil {
	  panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

}

func initDB() *gorm.DB {
	// Connect to the database.
	dbConfig := db.Config{
		User:	   viper.GetString("db.user"),
		Password:  viper.GetString("db.password"),
		Host:	   viper.GetString("db.host"),
		Port:	   viper.GetInt("db.port"),
		Name:	   viper.GetString("db.database"),
	}
	db := db.ConnectDB(dbConfig)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	  }), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return gormDB
}
