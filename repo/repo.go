package repo

import (
	"context"
	"streaming-service/config"
	"streaming-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var VideoCollection = "Videos"

func DB_CreateVideo(object *models.Video) error {
	filter := bson.M{"id": object.Id}
	update := bson.M{"$set": object}

	_, err := config.Database().Collection(VideoCollection).UpdateOne(
		context.Background(),
		filter,
		update,
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}
	return nil
}