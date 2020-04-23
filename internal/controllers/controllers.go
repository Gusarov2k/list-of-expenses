package controllers

import (
	"github.com/Gusarov2k/list-of-expenses/internal/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type (
	SpentIdParams struct {
		Id int `query:"id" validate:"required"`
	}

	SpentParams struct {
		Name 	string 	`json:"name" form:"name" query:"name" validate:"required"`
		Amount  float32	`json:"amount" form:"amount" query:"amount" validate:"required"`
	}
	)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type Router struct {

}

func (r *Router) Handler(db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/api/v1/spent/:id", func(c echo.Context) (err error) {
		var spent postgres.Spents

		u := new(SpentIdParams)
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

		return c.JSON(http.StatusOK, spent)
	})

	e.POST("/api/v1/spent", func(c echo.Context) (err error) {
		u := new(SpentParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		spent := db.Create(&postgres.Spents{Name: u.Name, Amount: u.Amount})

		return c.JSON(http.StatusCreated, spent.Value)
	})

	e.Logger.Fatal(e.Start(":3002"))
	return e
}

