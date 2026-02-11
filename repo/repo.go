package repo

import (
	"context"
	"errors"
	"streaming-service/config"
	"streaming-service/models"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var VideoCollection = "Videos"
var UserPrefetchedVideos = "User_Prefetched_Videos"

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


func DB_FindallUserPrefetchedVideos(id string) (*[]models.Video, error) {
	var objects []models.Video
	filter := bson.M{}
	if id != "" {
		filter["userid"] = id
	}
	results, err := config.Database().Collection(UserPrefetchedVideos).Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}
	for results.Next(context.Background()) {
		var object models.UserPrefetchedVideo
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding order")
		}
		objects = append(objects, object.Video)
	}
	return &objects, nil
}

func DB_CreateUserPrefetchedVideos(object *models.CreatePrefetchedVideosInput) (*[]models.Video, error) {
	// Delete existing prefetched videos for user
	_, err := config.Database().Collection(UserPrefetchedVideos).DeleteMany(
		context.Background(),
		bson.M{"user_id": object.UserId},
	)
	if err != nil {
		return nil,err
	}

	// Fetch full video objects and insert
	var videoDocs []interface{}
	var videos []models.Video
	for _, videoId := range object.Videos {
		var video models.Video
		err := config.Database().Collection(VideoCollection).FindOne(
			context.Background(),
			bson.M{"id": videoId},
		).Decode(&video)
		if err != nil {
			log.Error("Error fetching video with id %s: %v", videoId, err)
		}
		
		prefetchedVideo := models.UserPrefetchedVideo{
			UserId: object.UserId,
			Video:  video,
		}
		videoDocs = append(videoDocs, prefetchedVideo)
		videos = append(videos, video)
	}

	_, err = config.Database().Collection(UserPrefetchedVideos).InsertMany(
		context.Background(),
		videoDocs,
	)
	return &videos, err
}