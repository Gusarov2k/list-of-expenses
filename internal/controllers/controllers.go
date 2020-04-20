package controllers

import (
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
		Id      int 	`query:"id" validate:"required"`
		Name 	string 	`json:"name" form:"name" validate:"required"`
		Amount  float32	`json:"amount" form:"amount" validate:"required"`
	}

	SpentDataParams struct {
		Start_date 	time.Time	`query:"date_from" validate:"required"`
		End_date 	time.Time	`query:"date_to" validate:"required"`
	}
	)
