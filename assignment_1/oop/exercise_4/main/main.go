package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string
	Price    int
	Quantity int
}

func FromInstanceToJson(p Product) string {
	bytes, err := json.Marshal(p) // []byte

	if err != nil {
		fmt.Println(err)
	}

	return string(bytes)
}

func FromJsonToInstance(jsonText string) Product {
	var bytes []byte = []byte(jsonText)

	var product Product

	err := json.Unmarshal(bytes, &product)

	if err != nil {
		fmt.Println(err)
	}

	return product
}

func main() {
	product := Product{"Oreo", 500, 20}

	var dataInJson = FromInstanceToJson(product)
	fmt.Println(dataInJson)

	product1 := FromJsonToInstance(dataInJson)
	fmt.Printf("Product: \nName: %s;\nPrice: %d;\nQuantity: %d.\n", product1.Name, product1.Price, product1.Quantity)

}
