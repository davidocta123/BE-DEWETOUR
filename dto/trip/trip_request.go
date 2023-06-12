package tripdto

type CreateTripRequest struct {
	Title          string `json:"title" form:"title" validate:"required"`
	CountryId      int    `json:"country_id" form:"country_id" `
	Accomodation   string `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string `json:"transportation" form:"transportation" validate:"required"`
	Eat            string `json:"eat" form:"eat" validate:"required"`
	Day            int    `json:"day" form:"day" validate:"required"`
	Night          int    `json:"night" form:"night" validate:"required"`
	DateTrip       string `json:"dateTrip" form:"dateTrip" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	QuotaMax       int    `json:"quotaMax" form:"quotaMax" validate:"required"`
	QuotaFilled    int    `json:"quotaFilled" form:"quotaFilled" validate:"required"`
	Description    string `json:"description" form:"description" validate:"required"`
	Image          string `json:"image" form:"image" `
}

type UpdateTripRequest struct {
	Title          string `json:"title" form:"title"`
	CountryId      int    `json:"country_id" form:"country_id"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Eat            string `json:"eat" form:"eat"`
	Day            int    `json:"day" form:"day"`
	Night          int    `json:"night" form:"night"`
	DateTrip       string `json:"dateTrip" form:"dateTrip"`
	Price          int    `json:"price" form:"price"`
	QuotaMax       int    `json:"quotaMax" form:"quotaMax"`
	QuotaFilled    int    `json:"quotaFilled" form:"quotaFilled"`
	Description    string `json:"description" form:"description"`
	Image          string `json:"image" form:"image"`
}
