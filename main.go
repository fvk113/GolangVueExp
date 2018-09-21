package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"vuegootest/handlers"
)

func main() {
	db, err := gorm.Open("postgres", "user=convenios password=convenios dbname=conveniosweb sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()
	e.Static("public/resources", "resources")
	e.File("/home", "public/index.html")
	e.GET("/cuentas", handlers.GetAccounts(db))
	e.PUT("/cuentas", handlers.DeleteTask(db))
	e.DELETE("/cuentas/:id", handlers.PutTask(db))
	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
	fmt.Println("Servidor Iniciado")
}
