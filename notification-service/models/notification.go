package models

type Notification struct {
	UserId  int `json:"user_id"`
	OrderId int `json:"order_id"`
	Message string `json:"message"`
}
