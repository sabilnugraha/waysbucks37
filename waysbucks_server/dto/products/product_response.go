package productsdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: varchar(255)"`
	Image string `json:"image" form:"image" validate:"required"`
}
