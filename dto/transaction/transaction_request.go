package transactiondto

type CreateTransaction struct {
	Id             int    `json:"id_trans" form:"id_trans" gorm:"primary_key:auto_increment"`
	Title          string `json:"title" form:"title" validation:"required"`
	Day            int    `json:"day" form:"day" validation:"required"`
	Night          int    `json:"night" form:"night" validation:"required"`
	Country        int    `json:"country" form:"country" validation:"required"`
	DateTrip       string `json:"dateTrip" form:"dateTrip" validation:"required"`
	Transportation string `json:"transportation" form:"transportation" validation:"required"`
	Status         string `json:"status" form:"status" validation:"required"`
	Date           string `json:"date" form:"date" validation:"required"`
	CustomerName   string `json:"customerName" form:"customerName" validation:"required"`
	CustomerGender string `json:"customerGender" form:"customerGender" validation:"required"`
	CustomerPhone  string `json:"customerPhone" form:"customerPhone" validation:"required"`
	Amount         int    `json:"amount" form:"amount" validation:"required"`
	Total          int    `json:"total" form:"total" validation:"required"`
	IdTrip         int    `json:"idTrip" form:"idTrip" validation:"required"`
}

type UpdateTransaction struct {
	Id             int    `json:"id_trans" form:"id_trans" gorm:"primary_key:auto_increment"`
	Title          string `json:"title" form:"title" `
	Day            int    `json:"day" form:"day" `
	Night          int    `json:"night" form:"night" `
	Country        int    `json:"country" form:"country" `
	DateTrip       string `json:"dateTrip" form:"dateTrip" `
	Transportasi   string `json:"transportasi" form:"transportasi" `
	Status         string `json:"status" form:"status" `
	Date           string `json:"date" form:"date" `
	CustomerName   string `json:"customerName" form:"customerName" `
	CustomerGender string `json:"customerGender" form:"customerGender" `
	CustomerPhone  string `json:"customerPhone" form:"customerPhone" `
	Amount         int    `json:"amount" form:"amount" `
	Total          int    `json:"total" form:"total" `
	IdTrip         int    `json:"idTrip" form:"idTrip" `
}
