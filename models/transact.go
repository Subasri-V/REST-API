package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transact struct{
	Id primitive.ObjectID `json:"id" bson:"id"`
	Transaction_date time.Time `json:"transaction_date" bson:"transaction_date"`
	Amount int `json:"amount" bson:"amount"`
	Transaction_type string `json:"transaction_type" bson:"transaction_type"`
	Customer_id int `json:"customer_id" bson:"customer_id"`
}

type TDBResponse struct{
	Id int `json:"id" bson:"id"`
	Transaction_date time.Time `json:"transaction_date" bson:"transaction_date"`
	Amount int `json:"amount" bson:"amount"`
	Transaction_type string `json:"transaction_type" bson:"transaction_type"`
	Customer_id int `json:"customer_id" bson:"customer_id"`
}

