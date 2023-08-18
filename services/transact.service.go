package services

import (
	"context"
	"rest_api_app/interfaces"
	"rest_api_app/models"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransService struct {
	transactCollection *mongo.Collection
	ctx                context.Context
}

func NewTransServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ITransact {
	return &TransService{collection, ctx}
}

// CreateTransaction implements interfaces.Itransact.
func (t *TransService) CreateTransaction(transact *models.Transact) (*mongo.InsertOneResult, error) {
	transact.Id = primitive.NewObjectID()
	transact.Transaction_date = time.Now()
	res, err := t.transactCollection.InsertOne(t.ctx, &transact)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// GetAllTransaction implements interfaces.Itransact.
func (*TransService) GetAllTransaction(id int) {
	//panic("unimplemented")
}
