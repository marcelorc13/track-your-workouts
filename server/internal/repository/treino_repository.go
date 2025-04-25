package repository

import (
	"context"
	"server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TreinoRepository struct {
	Coll *mongo.Collection
}

func NewTreinoRepository(db *mongo.Database) *TreinoRepository {
	return &TreinoRepository{db.Collection("treino")}
}

func (tr TreinoRepository) CreateTreino(t models.Treino) (models.DBResponse, error) {
	_, err := tr.Coll.InsertOne(context.TODO(), t)
	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}
	return models.DBResponse{Success: true}, nil
}

func (tr TreinoRepository) GetTreinos() (models.DBResponse, error) {
	cursor, err := tr.Coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}

	var res []models.Treino
	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}

	return models.DBResponse{Success: true, Data: res}, nil
}
