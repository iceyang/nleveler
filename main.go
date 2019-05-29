package main

import (
	"fmt"
	"log"

	"github.com/iceyang/nleveler/db"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	action := &Action{
		Name: "english",
		Date: "2019-05-29",
	}
	fmt.Println(action)
	SaveAction(action)
}
