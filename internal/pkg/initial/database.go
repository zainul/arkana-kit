package initial

import (
	"context"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var conf Config
var db *gorm.DB
var mongoClient *mongo.Client
var mongoDB *mongo.Database

// var mongoDB *mongo.Database

// Config ...
type Config struct {
	Database database
	Mongo    mongoCred
	MQ       mq
}

func init() {

	if _, err := toml.DecodeFile("../configs/config.toml", &conf); err != nil {
		panic(err)
	}

	dbConn, err := gorm.Open(conf.Database.Provider,
		fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Name,
			conf.Database.Server,
			conf.Database.Port,
		),
	)

	if err != nil {
		panic(err)
	}

	dbConn.DB().SetMaxIdleConns(20)
	dbConn.DB().SetMaxOpenConns(50)

	db = dbConn

	fmt.Println("database is running ....")

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d",
		conf.Mongo.User, conf.Mongo.Password, conf.Mongo.Server, conf.Mongo.Port)))

	if err != nil {
		fmt.Println("failed to run mongo")
		return
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println("failed to ping mongodb")
		return
	}

	fmt.Println("mongodb success to run")

	mongoClient = client

	mongoDB = client.Database("central_log")
}

// GetMongoClient is get client
func GetMongoClient() *mongo.Client {
	return mongoClient
}

func GetMongoDB() *mongo.Database {
	return mongoDB
}

// GetDB ...
func GetDB() *gorm.DB {
	db = db.LogMode(true)
	return db
}
