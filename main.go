// main.go
package main

import (
	"fmt"
	"log"

	"example-go-grpc-gateway/protogen/golang/orders"
	"example-go-grpc-gateway/protogen/golang/product"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	orderItem := orders.Order{
		OrderID:    10,
		CustomerID: 11,
		IsActive:   true,
		OrderDate: &date.Date{
			Year:  2023,
			Month: 10,
			Day:   1,
		},
		Products: []*product.Product{
			{
				ProductID:    1,
				ProductName:  "Product 1",
				ProductTyepe: product.ProductType_DRINK,
			},
		},
	}
	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatalf("Failed to marshal order item: %v", err)
	}
	fmt.Printf("Order Item: %v\n", string(bytes))
}
