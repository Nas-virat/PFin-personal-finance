package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Nas-virat/PFin-personal-finance/constant"
	"github.com/Nas-virat/PFin-personal-finance/db"
	"github.com/Nas-virat/PFin-personal-finance/handlers"
	"github.com/Nas-virat/PFin-personal-finance/log"
	"github.com/Nas-virat/PFin-personal-finance/repository"
	"github.com/Nas-virat/PFin-personal-finance/service"
	"github.com/Nas-virat/PFin-personal-finance/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {

	logger := log.NewLogger()

	// Initialize Timezone
	initTimeZone()
	logger.Info("Initialized timezone")

	// Initialize config
	initConfig()
	logger.Info("Initialized config")

	db := initDB()
	logger.Info("Initialized database")
	
	// Migrate the schema
	utils.Migration(db)
	logger.Info("Migrated the schema")



	if constant.IsDevelopment {
		logger.Warn("Running in development mode")
	}
	
	app := fiber.New()
	app.Use(cors.New())
	logger.Info("Initialized fiber")


	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handlers.NewAccountHandler(accountService)


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	app.Get("/createaccount",accountHandler.CreateAccountHandler)

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
