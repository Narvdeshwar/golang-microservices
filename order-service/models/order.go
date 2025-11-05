package models

type Order struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Item   string  `json:"item"`
	Amount float32 `json:"amount"`
}
