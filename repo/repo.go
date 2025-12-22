package repo

import (
	"context"
	"errors"
	"streaming-service/config"
	"streaming-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func DB_FindallVideo(id string) (*[]models.Video, error) {
	var objects []models.Video
	filter := bson.M{}
	if id != "" {
		filter["id"] = id
	}
	results, err := config.Database().Collection(VideoCollection).Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}
	for results.Next(context.Background()) {
		var object models.Video
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding order")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
