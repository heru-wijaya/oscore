package model

import (
	"context"
	"fmt"
	"ovo-score/db"
	"ovo-score/dto"
)

// InsertEntity for inserting entity into db
func InsertEntity(ctx context.Context, entity *dto.Entity) error {
	client := db.GetClient()

	result, err := client.Database("Test").Collection("Entity").InsertOne(ctx, entity)
	defer client.Disconnect(ctx)
	fmt.Println("Inserted a Single Document: ", result.InsertedID)

	return err
}
