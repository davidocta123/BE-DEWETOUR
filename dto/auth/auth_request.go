package authdto

import "time"

type AuthRequest struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	FullName  string    `json:"Fullname" gorm:"varchar(255)" validation:"required"`
	Email     string    `json:"email" gorm:"varchar(255)" validation:"required"`
	Password  string    `json:"password" gorm:"varchar(255)" validation:"required"`
	Phone     string    `json:"phone" gorm:"varchar(255)" validation:"required"`
	Address   string    `json:"address" gorm:"varchar(255)" validation:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}
