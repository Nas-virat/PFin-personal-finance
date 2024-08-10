package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Nas-virat/PFin-personal-finance/constant"
	"github.com/Nas-virat/PFin-personal-finance/db"
	docs "github.com/Nas-virat/PFin-personal-finance/docs"

	"github.com/Nas-virat/PFin-personal-finance/log"
	"github.com/Nas-virat/PFin-personal-finance/router"
	"github.com/Nas-virat/PFin-personal-finance/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func main() {

	logger := log.NewLogger()

	// Initialize Timezone
	initTimeZone()
	logger.Info("Initialized timezone")

	db := initDB()
	logger.Info("Initialized database")

	// Migrate the schema
	utils.Migration(db)
	logger.Info("Migrated the schema")

	if constant.IsDevelopment {
		logger.Warn("Running in development mode")
	}

	r := gin.Default()

	v1 := r.Group("/api")
	docs.SwaggerInfo.BasePath = "/api"

	router.SetupAccountRoutes(v1, db)
	router.SetupTransactionRoute(v1, db)
	router.SetupBalanceRoutes(v1, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Server exiting")
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
		User:       os.Getenv("DATABASE_USER"),
		Password:   os.Getenv("DATABASE_PASSWORD"),
		Host:       os.Getenv("DATABASE_HOST"),
		Port:       port,
		Name:       os.Getenv("DATABASE_NAME"),
		TimeZone:   os.Getenv("DATABASE_TIMEZONE"),
		DisableTLS: true,
	}
	db := db.ConnectDB(dbConfig)

	return db
}
