package main

import (
	"context"
	"os"

	// "fmt"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGOHOST = os.Getenv("MONGO_HOST")
var MONGOPORT = os.Getenv("MONGO_PORT")
var MONGOUSER = os.Getenv("MONGO_USER")
var MONGOPASS = os.Getenv("MONGO_PASS")
var MONGOAUTHDB = os.Getenv("MONGO_AUTHDB")

var RMQHOST = os.Getenv("RMQ_HOST")
var RMQPORT = os.Getenv("RMQ_PORT")
var RMQUSER = os.Getenv("RMQ_USER")
var RMQPASS = os.Getenv("RMQ_PASS")

// var RMQVHOST = os.Getenv("RMQ_VHOST")

var REDISHOST = os.Getenv("REDIS_HOST")
var REDISPORT = os.Getenv("REDIS_PORT")

var RedisConnection = REDISHOST + ":" + REDISPORT
var RabbitMQConnection = "amqp://" + RMQUSER + ":" + RMQPASS + "@" + RMQHOST + ":" + RMQPORT + "/"

func MongoConnect() *mongo.Database {
	// Set client options

	clientOptions := options.Client().ApplyURI("mongodb://" + MONGOHOST + ":" + MONGOPORT)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// Connect to DB
	db := client.Database("simple-editor")
	// fmt.Printf("%T", db)

	return db
}
