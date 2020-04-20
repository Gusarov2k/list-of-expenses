package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"github.com/Gusarov2k/list-of-expenses/internal/postgres"
	"github.com/Gusarov2k/list-of-expenses/internal/controllers"
)
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main()  {
	// init DB
	db, err := gorm.Open("postgres", postgres.PostgresSys)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	fmt.Printf("%s\n", err)
	db.LogMode(true)

	db.AutoMigrate(&postgres.Spent{})
	// init DB

	e := echo.New()

	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group("/api/v1")

	g.GET("/spent/:id", func(c echo.Context) (err error) {
		var spent postgres.Spent
		var shorSpent postgres.ShortSpent
		u := new(controllers.SpentIdParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		result := db.First(&spent, u.Id)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, "spent not find")
		}

		if err := copier.Copy(&shorSpent, &spent); err != nil {
			return
		}

		return c.JSON(http.StatusOK, shorSpent)
	})

	g.POST("/spent", func(c echo.Context) (err error) {
		u := new(controllers.SpentParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		spent := db.Create(&postgres.Spent{Name: u.Name, Amount: u.Amount})

		return c.JSON(http.StatusCreated, spent.Value)
	})
	e.Logger.Fatal(e.Start(":3002"))
}
