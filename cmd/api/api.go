package main

import (
	"github.com/Gusarov2k/list-of-expenses/internal/controllers"
	"github.com/Gusarov2k/list-of-expenses/internal/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main()  {

	c := postgres.NewClient()

	if err := c.Open(postgres.PostgresSys); err != nil {
		panic("failed to connect database")
	}
	c.Schema()
	defer func() { c.Close() }()

	e := controllers.Router{}
	e.Handler(c.Db)

}
