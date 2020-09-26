package db

import (
	"context"
	"fmt"

	"github.com/guowenshuai/download/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGOCLIENT = "mongoclient"

func Connect(config *conf.Config) (*mongo.Client, error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Mongodb.Server).SetAuth(options.Credential{
		Username:    config.Mongodb.Username,
		Password:    config.Mongodb.Password,
		PasswordSet: false,
	})

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}
