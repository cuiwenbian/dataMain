package iso

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guowenshuai/download/conf"
	"github.com/guowenshuai/download/db"
	"github.com/guowenshuai/download/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func List(c *gin.Context) {
	vl, ok := c.Get(conf.CONFIG)
	if !ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", "err config"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	config := vl.(*conf.Config)
	vl, ok = c.Get(db.MONGOCLIENT)
	if !ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", "no db connect"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	db := vl.(*mongo.Client)
	collection := db.Database(config.Mongodb.Database).Collection("images")
	// 查询多个
	// 将选项传递给Find()
	findOptions := options.Find()

	// 定义一个切片用来存储查询结果
	var results []*types.Image
	// https://kb.objectrocket.com/mongo-db/how-to-find-a-mongodb-document-by-its-bson-objectid-using-golang-452
	// 把bson.D{{}}作为一个filter来匹配所有文档
	// did, _ := primitive.ObjectIDFromHex("5f6f2d1eb983b3bfbe66e429")
	// cur, err := collection.Find(context.TODO(), bson.D{{"_id", did}}, findOptions)
	cur, err := collection.Find(context.TODO(), bson.D{
		{"cid", bson.D{{"$ne", ""}}}, // cid为空的去掉
		{"size", bson.D{{"$gt", 0}}}, // size为空的去掉
	}, findOptions)

	if err != nil {
		log.Fatal(err)
	}
	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem types.Image
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())
	// fmt.Printf("Found multiple documents (array of pointers): %#v\n", results)

	images := make(map[string][]*types.Image)
	for i, v := range results {
		images[v.OS] =  append(images[v.OS], results[i])
	}

	c.JSON(http.StatusOK, images)
}
