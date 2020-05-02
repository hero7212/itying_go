package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/tidwall/gjson"
)

func main() {
	num1, num2 := 3.1, 4.2
	fmt.Println(num1 + num2)
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(d1)

	const json = `{"name":{"first":"Janet","last":"四"},"age":47}`

	value := gjson.Get(json, "name.last")
	println(value.String())
}
