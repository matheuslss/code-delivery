package main

import (
	"fmt"

	simRoute "github.com/matheuslss/code-delivery/simulator/application/route"
)

func main() {
	route := simRoute.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[0])
}
