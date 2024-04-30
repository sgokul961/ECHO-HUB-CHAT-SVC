package db

import (
	"context"
	"fmt"

	"github.com/sgokul961/echo-hub-chat-svc/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo(c config.Config) (*mongo.Database, error) {
	ctx := context.TODO()
	mongoConn := options.Client().ApplyURI(c.DB_URL)

	mongoConn.SetAuth(options.Credential{
		Username:      c.DBUsername,
		Password:      c.DBPassword,
		AuthSource:    c.DBUsername,
		AuthMechanism: c.AuthMechanism,
	})
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		return nil, err
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("mongo connection established")

	return mongoClient.Database(c.DBName), nil
}
