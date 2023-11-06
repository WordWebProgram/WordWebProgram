package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	Id        interface{} `json:"id" bson:"_id"`
	Username  string      `json:"name" db:"name"`
	Password  string      `json:"pswd" db:"pswd"`
	CreatedAt string      `json:"createdAt" db:"createdAt"`
}

func FetchUserInfoByName(db *mongo.Client, username string) (*User, error) {
	// 创建一个Student变量用来接收查询的结果
	filter := bson.D{{"username", username}}
	// 指定获取要操作的数据集
	collection := db.Database("neet-words").Collection("user")
	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return &result, err
}
func InsertUser(db *mongo.Client, username string, password string) error {
	collection := db.Database("neet-words").Collection("user")
	result := User{
		Username: username, Password: password, CreatedAt: time.Now().String(),
	}
	insertResult, err := collection.InsertOne(context.TODO(), result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return err
}
