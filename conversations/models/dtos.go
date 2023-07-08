package models

type ConvCreateDto struct {
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId"`
}
