package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-api-server/handler"
	"go-api-server/repository"
)

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)
	// Routes
	e.GET("/", handler.ArticleIndex)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	return e
}

func connectDB() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}
