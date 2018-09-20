//accountHandler.go
package handlers

import (
	"awesomeProject/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// handlers/tasks.go
type H map[string]interface{}

// GetTasks endpoint
func GetAccounts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetCuentas(db))
	}
}

// PutTask endpoint
func PutTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": 123,
		})
	}
}

// DeleteTask endpoint
func DeleteTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
