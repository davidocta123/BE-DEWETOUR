package models

import "time"

type Country struct {
	Id   int    `json:"id_country" form:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"country" form:"name" gorm:"type:varchar(255)" validation:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CountriesResponse struct {
	Id   int    `json:"id_country"`
	Name string `json:"country"`
}

func (CountriesResponse) TableName() string {
	return "Countries"
}
