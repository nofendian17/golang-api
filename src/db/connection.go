package db

import (
	"fmt"
	"github.com/Kamva/mgm/v3"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func init()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectMongoDB() {
	dsn := fmt.Sprintf( "mongodb://%s:%s@%s:%s/%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DATABASE"))
	fmt.Println("connecting... " + dsn)
	err := mgm.SetDefaultConfig(nil, os.Getenv("MONGO_DATABASE"), options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}
	fmt.Println("connected...")
}