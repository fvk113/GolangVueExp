//models.go
package models

import (
	"awesomeProject/dominio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetCuentas(db *gorm.DB) dominio.AccountCollection {
	db.SingularTable(true)
	var accounts []dominio.Account
	db.Debug().Find(&accounts)
	collection := dominio.AccountCollection{}
	collection.Accounts = accounts
	fmt.Println(accounts)
	db.Close()
	return collection
}
