package models

type Payment struct {
	ID      int    `json:"id"`
	OrderID int    `json:"order_id"`
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
}
