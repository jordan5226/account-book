package main

import (
	"account-book/control"
	"account-book/lib/pgdb"
	"account-book/lib/pgdb/schema"
	"account-book/middleware"
	"account-book/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func chkNote(fl validator.FieldLevel) bool {
	if note, ok := fl.Field().Interface().(string); ok {
		if strings.Contains(note, ";--") {
			return false
		} else {
			return true
		}
	}
	return false
}

func ValidateJSONDateType(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(schema.LocalTime{}) {
		timeStr := field.Interface().(schema.LocalTime).String()

		if timeStr == "0001-01-01 00:00:00" {
			// 0001-01-01 00:00:00 is nil in time.Time
			// return nil will judge as empty value in validator and cannot pass `binding:"required"` rule
			return nil
		}

		return timeStr
	}

	return nil
}

func init() {
	model.Init(pgdb.GetConnect())
}

func main() {
	// Create Gin Instance
	g := gin.New()
	g.Use(middleware.CustomHttpErrorHandler)

	// Set Logger
	gin.DisableConsoleColor()
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	fileErr, _ := os.Create("gin_err.log")
	gin.DefaultErrorWriter = io.MultiWriter(fileErr, os.Stdout)

	// Register Validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("chkNote", chkNote); err != nil {
			fmt.Println("[ AccountBook ] main.main - RegisterValidation success")
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(ValidateJSONDateType, schema.LocalTime{})
	}

	// Set Base Route
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hi!")
	})

	// Create Route Management
	c := control.NewAccountBook()
	c.SetRout(g.Group("/acctbook"))

	// Run Service
	g.Run(":8080")
}
