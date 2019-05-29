package main

import (
	"context"
	"fmt"
	"log"

	"github.com/iceyang/nleveler/db"
)

type Action struct {
	Name string "name"
	Date string "date"
}

func SaveAction(action *Action) {
	collection, err := db.Collection("action")
	if err != nil {
		log.Panic(err)
	}
	result, err := collection.InsertOne(context.Background(), action)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result.InsertedID)
}
