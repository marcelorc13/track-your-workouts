package repository

import (
	"context"
	"server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TreinoRepository struct {
	DB *mongo.Database
}

func NewTreinoRepository(db *mongo.Database) *TreinoRepository {
	return &TreinoRepository{db}
}

func (tr TreinoRepository) CreateTreino(t models.Treino) (models.DBResponse, error) {
	_, err := tr.DB.Collection("treino").InsertOne(context.TODO(), t)
	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}
	return models.DBResponse{Success: true}, nil
}

func (tr TreinoRepository) GetTreinosDoUsuario(usuarioId string) (models.DBResponse, error) {
	cursor, err := tr.DB.Collection("treino").Find(context.TODO(), bson.D{{Key: "criado_por", Value: usuarioId}})
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

func (tr TreinoRepository) GetTreinoById(id string) (models.DBResponse, error) {
	var treino models.Treino

	primtiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.DBResponse{Message: "erro para converter id para primitive"}, err
	}

	err = tr.DB.Collection("treino").FindOne(context.TODO(), bson.M{"_id": primtiveId}).Decode(&treino)

	if err == mongo.ErrNoDocuments {
		return models.DBResponse{Message: "o treino não existe"}, err
	}
	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}

	return models.DBResponse{Success: true, Data: treino}, nil
}

func (tr TreinoRepository) DeleteTreino(id string) (models.DBResponse, error) {
	primtiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.DBResponse{Message: "erro para converter id para primitive"}, err
	}

	_, err = tr.DB.Collection("treino").DeleteOne(context.TODO(), bson.M{"_id": primtiveId})

	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}

	return models.DBResponse{Success: true}, nil
}

func (tr TreinoRepository) CreateSecao(s models.Secao) (models.DBResponse, error) {
	var treino models.Treino
	err := tr.DB.Collection("treino").FindOne(context.TODO(), bson.M{"_id": s.IDTreino}).Decode(&treino)

	if err == mongo.ErrNoDocuments {
		return models.DBResponse{Message: "o treino não existe"}, err
	}

	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}

	_, err = tr.DB.Collection("secao_de_treino").InsertOne(context.TODO(), s)
	if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}
	return models.DBResponse{Success: true}, nil

}
