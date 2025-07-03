package model

type Branch struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	OpenTime    string `json:"open_time"`
	CloseTime   string `json:"close_time"`
}
