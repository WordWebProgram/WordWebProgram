package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Progress struct {
	user  string
	lists []List
}
type List struct {
	location  int
	startedAt string
	updatedAt string
}

func FindList(db *mongo.Client, username string) (*Progress, error) {
	filter := bson.D{{"username", username}}
	// 指定获取要操作的数据集
	collection := db.Database("neet-words").Collection("lists")
	var result Progress
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return &result, err
}
