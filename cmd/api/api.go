package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"github.com/go-playground/validator/v10"
	"time"
)

type (
	SpentIdParams struct {
		Id int `query:"id" validate:"required"`
	}

	SpentParams struct {
		Name 	string 	`json:"name" form:"name" query:"name" validate:"required"`
		Amount float32	`json:"amount" form:"amount" query:"amount" validate:"required"`
	}

	SpentDataParams struct {
		Start_date 	time.Time	`query:"date_from" validate:"required"`
		End_date 	time.Time	`query:"date_to" validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main()  {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group("/api/v1")

	g.GET("/spent/:id", func(c echo.Context) (err error) {
		u := new(SpentIdParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	g.POST("/spent", func(c echo.Context) (err error) {
		u := new(SpentParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusCreated, u)
	})

	g.PUT("/spent/:id", func(c echo.Context) (err error) {
		u := new(SpentIdParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	g.DELETE("/spent/:id", func(c echo.Context) (err error) {
		u := new(SpentIdParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	g.GET("/spents", func(c echo.Context) (err error) {
		u := new(SpentDataParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusCreated, u)
	})
	
	e.Logger.Fatal(e.Start(":3002"))
}
