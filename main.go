package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
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
	fmt.Println("Conexión exitosa")
	db.SingularTable(true)
	var acc Account
	db.AutoMigrate(Account{})
	db.Debug().First(&acc)
	fmt.Println(acc)
	db.Close()
}

type Account struct {
	ID         string
	EMAIL      string
	FIRST_NAME string
	LAST_NAME  string
	LOGIN      string
	PASSWORD   string
	ROLE_NAME  string
	IS_ENABLED bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}
