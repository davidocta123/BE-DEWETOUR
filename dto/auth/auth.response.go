package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Role     string `json:"role"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
