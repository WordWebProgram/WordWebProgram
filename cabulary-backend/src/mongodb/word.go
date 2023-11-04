package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Progress struct {
	User  string `json:"user"`
	lists interface{}
}
type List struct {
	location  int
	startedAt string
	updatedAt string
}
type OldResult struct {
	User  string      `json:"user"`
	Lists interface{} `json:"lists"`
	_Id   int         `json:"_id"`
}

func FindList(db *mongo.Client, userId string) (*Progress, error) {
	filter := bson.D{{"userid", userId}}
	// 指定获取要操作的数据集
	collection := db.Database("neet-words").Collection("lists")
	var result Progress
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		result := Progress{
			User: userId, lists: nil,
		}
		_, err := collection.InsertOne(context.TODO(), result)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return &result, err
}
func UpdateList(db *mongo.Client, userId string, lists OldResult) (error, error) {
	filter := bson.D{{"User", userId}}
	var update = bson.M{
		"User":  lists.User,
		"Lists": lists.Lists,
	} // execute the UpdateOne() function to update the first matching document

	// 指定获取要操作的数据集
	collection := db.Database("neet-words").Collection("lists")
	_, err := collection.UpdateOne(context.TODO(), filter, bson.M{
		"$set": update,
	}, options.Update().SetUpsert(true))
	if err != nil {
		fmt.Println(err)
	}
	return err, nil
}
