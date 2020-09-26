package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/guowenshuai/download/route/iso"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/guowenshuai/download/conf"
	"github.com/guowenshuai/download/db"
)

var config *conf.Config

func init() {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.Printf("yamlFile.Get err   #%v ", err)
	}
	config = &conf.Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		logrus.Panic("read config err: %s", err.Error())
	}
}

func main() {
	mongoClient, err := db.Connect(config)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	router := gin.Default()
	router.Use(func(context *gin.Context) {
		context.Set(db.MONGOCLIENT, mongoClient)
		context.Set(conf.CONFIG, config)
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/iso", iso.List)
	}
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, Geektutu")
	// })
	router.Run() // listen and serve on 0.0.0.0:8080
}
