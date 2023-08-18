package interfaces

import (
	"rest_api_app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ITransact interface {
	GetAllTransaction(id int)
	CreateTransaction(transact *models.Transact) (*mongo.InsertOneResult, error)
}
