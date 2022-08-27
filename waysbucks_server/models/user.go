package models

import "time"

type User struct {
	ID        int             `json:"id"`
	Fullname  string          `json:"fullname" gorm:"type: varchar(255)"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"-" gorm:"type: varchar(255)"`
	Status    string          `json:"status" gorm:"default:customer"`
	Address   string          `json:"address" gorm:"type: varchar(255)"`
	PostCode  int             `json:"post"`
	Phone     int             `json:"phone" gorm:"type: varchar(255)"`
	Image     string          `json:"image" gorm:"type: varchar(255)"`
	Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
