package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateProjSectDescStorage(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	collection := s.DB.Collection(projSectDescCollection)

	_, err := collection.InsertOne(ctx, projSectDesc)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetAllProjSectDescStorage(ctx context.Context) ([]model.ProjSectDesc, error) {
	collection := s.DB.Collection(projSectDescCollection)

	result := make([]model.ProjSectDesc, 0)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, handleError("error occurred in Find", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		record := model.ProjSectDesc{}

		if err := cursor.Decode(&record); err != nil {
			return nil, fmt.Errorf("error occurred in cursor.Decode: %w", err)
		}

		result = append(result, record)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error occurred in cursor.Err: %w", err)
	}

	return result, nil
}

func (s Storage) UpdateProjSectDescByIDStorage(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	collection := s.DB.Collection(projSectDescCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: projSectDesc.ID}}, projSectDesc)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}
