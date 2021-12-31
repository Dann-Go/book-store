package indexes

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateIndex(client *mongo.Client, dbName string, collectionName string, field string, unique bool) {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection(collectionName)

	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Error(err)
	}
}
