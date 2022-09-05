package main

import (
	"context"
	"fmt"

	"github.com/caraml-dev/protopath"
	samplev1 "github.com/caraml-dev/protopath/internal/gen/sample/v1"
)

func main() {
	order := &samplev1.Order{
		OrderId: "MR-123",
		LineItems: []*samplev1.LineItem{
			{
				LineItemId: "1",
				BasePrice:  1000,
				Markup:     0.0,
				Quantity:   1,
				IsPromo:    false,
			},
			{
				LineItemId: "2",
				BasePrice:  2000,
				Markup:     0.0,
				Quantity:   2,
				IsPromo:    false,
			},
			{
				LineItemId: "3",
				BasePrice:  3000,
				Markup:     0.2,
				Quantity:   3,
				IsPromo:    true,
			},
		},
	}

	// Find line items that is promotion item
	compiledProtopath, err := protopath.NewJsonPathCompiler("$.line_items[?(@.is_promo == true)]")
	if err != nil {
		panic(err)
	}
	res, err := compiledProtopath.Lookup(context.TODO(), order)
	if err != nil {
		panic(err)
	}
	fmt.Println(res) // [line_item_id:"3"  base_price:3000  markup:0.2  is_promo:true  quantity:3]
}
