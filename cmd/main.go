package main

import (
	"FinanceApp/internal/banking"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://db:27017"))
	if err != nil {
		log.Println(err)
	}

	for {
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		id := rand.Intn(10000000)
		res, err := client.Database("finance").Collection("accounts").InsertOne(ctx, banking.Account{fmt.Sprintf("%v", id)})
		if err != nil {
			log.Println(err)
		}
		log.Printf("Created random account with id: %v and stored in the database with id: %v\n", id, res.InsertedID)
		time.Sleep(3 * time.Second)
	}
}
