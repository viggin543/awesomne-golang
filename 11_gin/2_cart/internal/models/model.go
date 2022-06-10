package models

type Cart struct {
	Id         string `json:"id" bson:"id"`
	TotalCents int64  `json:"totalCents" bson:"totalCents"`
}
