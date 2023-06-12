package transactiondto

import "dumbmerch/models"

type TransactionResponse struct {
	Id             int `json:"id_trans" form:"id_trans" gorm:"primary_key:auto_increment"`
	IdUser         int `json:"idUser" form:"idUser"`
	User           models.UsersProfileResponse
	Title          string `json:"title" form:"title"`
	Day            int    `json:"day" form:"day"`
	Night          int    `json:"night" form:"night"`
	Country        int    `json:"country" form:"country"`
	DateTrip       string `json:"dateTrip" form:"dateTrip"`
	Transportation string `json:"transportation" form:"transportation"`
	Status         string `json:"status" form:"status"`
	Date           string `json:"date" form:"date"`
	CustomerName   string `json:"customerName" form:"customerName"`
	CustomerGender string `json:"customerGender" form:"customerGender"`
	CustomerPhone  string `json:"customerPhone" form:"customerPhone"`
	Amount         int    `json:"amount" form:"amount"`
	Total          int    `json:"total" form:"total"`
	IdTrip         int    `json:"idTrip" form:"idTrip"`
}
