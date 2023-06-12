package models

import "time"

type Trip struct {
	Id             int               `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Title          string            `json:"title" form:"title"  `
	CountryId      int               `json:"country_id" form:"country_id"`
	Country        CountriesResponse `json:"country" form:"country" gorm:"foreignkey:CountryId"`
	Accomodation   string            `json:"accomodation" form:"accomodation"  `
	Transportation string            `json:"transportation" form:"transportation"  `
	Eat            string            `json:"eat" form:"eat"  `
	Day            int               `json:"day" form:"day"  `
	Night          int               `json:"night" form:"night"  `
	DateTrip       string            `json:"dateTrip" form:"dateTrip"  `
	Price          int               `json:"price" form:"price"  `
	QuotaMax       int               `json:"quotaMax" form:"quotaMax"  `
	QuotaFilled    int               `json:"quotaFilled" form:"quotaFilled"  `
	Description    string            `json:"description" form:"description"  `
	Image          string            `json:"image" form:"image" `
	CreatedAt      time.Time         `json:"-"`
	UpdatedAt      time.Time         `json:"-"`
}

type TripResponse struct {
	ID             int               `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Title          string            `json:"title" form:"title"`
	CountryId      int               `json:"country_id" form:"country_id"`
	Country        CountriesResponse `json:"country" form:"country" gorm:"foreignkey:CountryId"`
	Accomodation   string            `json:"accomodation" form:"accomodation"`
	Transportation string            `json:"transportation" form:"transportation"`
	Eat            string            `json:"eat" form:"eat"`
	Day            int               `json:"day" form:"day"`
	Night          int               `json:"night" form:"night"`
	DateTrip       string            `json:"dateTrip" form:"dateTrip"`
	Price          int               `json:"price" form:"price"`
	QuotaMax       int               `json:"quotaMax" form:"quotaMax"`
	QuotaFilled    int               `json:"quotaFilled" form:"quotaFilled"`
	Description    string            `json:"description" form:"description"`
	Image          string            `json:"image" form:"image"`
}

func (TripResponse) TableName() string {
	return "trips"
}
