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
	"os"
	"time"
)

type (
	SpentIdParams struct {
		Id int `query:"id" validate:"required"`
	}

	SpentParams struct {
		Name 	string 	`json:"name" form:"name" query:"name" validate:"required"`
		Amount  float32	`json:"amount" form:"amount" query:"amount" validate:"required"`
	}

	SpentUpdateParams struct {
		Id      int `query:"id" validate:"required"`
		Name 	string 	`json:"name" form:"name" validate:"required"`
		Amount  float32	`json:"amount" form:"amount" validate:"required"`
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

// DB


type (
	Spent struct {
		gorm.Model
		Name string `gorm:"type:varchar(256);not null json:"name"`
		Amount float32 `json:"amount"`
	}

	ShortSpent struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Amount float32 `json:"amount"`
		CreatedAt time.Time `json:"crated_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

var (
	PostgresHost           = getEnv("POSTGRES_HOST", "localhost")
	PostgresPort           = getEnv("POSTGRES_PORT", "5432")
	PostgresDB             = getEnv("POSTGRES_DB", "list_expense_development")
	PostgresDBTest         = getEnv("POSTGRES_DB_TEST", "list_expense_test")
	PostgresUser           = getEnv("POSTGRES_USER", "ivan")
	PostgresPassword       = getEnv("POSTGRES_PASSWORD", "1234")
	PostgresConnectTimeout = getEnv("POSTGRES_CONNECT_TIMEOUT", "3")

	PostgresSys = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s connect_timeout=%s sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB, PostgresConnectTimeout)

	PostgresTest = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s connect_timeout=%s sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDBTest, PostgresConnectTimeout)

)

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}

	return value
}
// DB

func main()  {
	// init DB
	db, err := gorm.Open("postgres", PostgresSys)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	fmt.Printf("%s\n", err)
	db.LogMode(true)

	db.AutoMigrate(&Spent{})
	// init DB

	e := echo.New()

	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group("/api/v1")

	g.GET("/spent/:id", func(c echo.Context) (err error) {
		var spent Spent
		var shorSpent ShortSpent
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

		copier.Copy(&shorSpent, &spent)

		return c.JSON(http.StatusOK, shorSpent)
	})

	g.POST("/spent", func(c echo.Context) (err error) {
		u := new(SpentParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		spent := db.Create(&Spent{Name: u.Name, Amount: u.Amount})

		return c.JSON(http.StatusCreated, spent.Value)
	})

	g.PUT("/spent/:id", func(c echo.Context) (err error) {
		var spent Spent
		var shortSpent ShortSpent

		u := new(SpentUpdateParams)
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

		db.Model(&spent).Updates(Spent{Name: u.Name, Amount: u.Amount})
		copier.Copy(&shortSpent, &spent)
		return c.JSON(http.StatusOK, shortSpent)
	})

	g.DELETE("/spent/:id", func(c echo.Context) (err error) {
		var spent Spent
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

		db.Delete(&spent)
		return c.JSON(http.StatusOK, "spent deleted")
	})

	g.GET("/spents", func(c echo.Context) (err error) {
		var spent []Spent
		var shortSpent []ShortSpent

		u := new(SpentDataParams)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}

		result := db.Where("created_at >= ? AND updated_at <= ?", u.Start_date, u.End_date).Find(&spent)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, "spents not find")
		}

		copier.Copy(&shortSpent, &spent)
		return c.JSON(http.StatusCreated, shortSpent)
	})
	
	e.Logger.Fatal(e.Start(":3002"))
}
