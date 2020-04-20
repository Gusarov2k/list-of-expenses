package main

import (
	"fmt"
	"github.com/Gusarov2k/list-of-expenses/internal/controllers"
	"github.com/Gusarov2k/list-of-expenses/internal/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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

	db.Exec(postgres.Schema)

	// Drop table `users`
	//db.DropTable("spent")
	// init DB

	e := echo.New()

	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group("/api/v1")

	g.GET("/spent/:id", func(c echo.Context) (err error) {
		var spent postgres.Spent

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

		return c.JSON(http.StatusOK, spent)
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
