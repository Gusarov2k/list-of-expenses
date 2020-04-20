package controllers

type (
	SpentIdParams struct {
		Id int `query:"id" validate:"required"`
	}

	SpentParams struct {
		Name 	string 	`json:"name" form:"name" query:"name" validate:"required"`
		Amount  float32	`json:"amount" form:"amount" query:"amount" validate:"required"`
	}
	)
