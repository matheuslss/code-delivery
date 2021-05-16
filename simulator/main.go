package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	simRoute "github.com/matheuslss/code-delivery/simulator/application/route"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := simRoute.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[0])
}
