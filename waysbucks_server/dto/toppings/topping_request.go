package toppingsdto

type CreateToppingRequest struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
}

type UpdateToppingRequest struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
}
