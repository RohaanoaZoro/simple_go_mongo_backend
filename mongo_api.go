package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mnDB = MongoConnect()

func Mongo_Update_Content(userid string, docid string, content string) error {

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	log.Println("User_ID", userid)
	filter := bson.M{"User_ID": userid, "Doc_ID": docid}
	update := bson.M{"$set": bson.M{
		"Content": content,
	}}

	_, err := mnDB.Collection("user_content").UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	return err
}

func Mongo_Get_Content(userid string, docid string) (string, error) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	var userContent map[string]string

	err := mnDB.Collection("user_content").FindOne(ctx, bson.M{"User_ID": userid, "Doc_ID": docid}, options.FindOne().SetProjection(bson.M{"_id": 0, "Content": 1})).Decode(&userContent)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return userContent["Content"], nil
}
