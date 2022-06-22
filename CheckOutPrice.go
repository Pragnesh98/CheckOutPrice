package main

import (
	"fmt"
	"strconv"
	"strings"
)

var productPrices string = `A $50 3 for $130
B $30 2 for $45
C $20
D $15`

type UnitDetail struct {
	Qty   []int
	Price []int
}

var ProductCatlog map[string]UnitDetail = make(map[string]UnitDetail)

type CheckOut struct {
	ProductCatlog map[string]UnitDetail
	InputUnits    map[string]int
}

func (co *CheckOut) scan(item string) {
	co.InputUnits[item] += 1
}

func (co *CheckOut) total() int {
	totalSum := 0
	for productName, cnt := range co.InputUnits {
		maxQtyLen := len(co.ProductCatlog[productName].Qty) - 1
		for cnt > 0 {
			dscQty := cnt / co.ProductCatlog[productName].Qty[maxQtyLen]
			cnt %= co.ProductCatlog[productName].Qty[maxQtyLen]

			totalSum += co.ProductCatlog[productName].Price[maxQtyLen] * dscQty
			maxQtyLen--
		}
	}
	co.InputUnits = make(map[string]int)
	return totalSum
}

func InitCheckout(inp string) *CheckOut {
	products := strings.Split(inp, "\n")

	for _, product := range products {
		productValues := strings.Split(product, " ")

		arrLen := ((len(productValues) - 2) / 3) + 1
		ud := UnitDetail{Qty: make([]int, arrLen), Price: make([]int, arrLen)}
		ud.Qty[0] = 1
		tempVal, _ := strconv.Atoi(productValues[1][1:])
		ud.Price[0] = tempVal
		for i := 1; i < arrLen; i++ {
			start := (3 * i) - 1
			ud.Qty[i], _ = strconv.Atoi(productValues[start])
			ud.Price[i], _ = strconv.Atoi(productValues[start+2][1:])
		}
		ProductCatlog[productValues[0]] = ud
	}

	cO := CheckOut{
		ProductCatlog: ProductCatlog,
		InputUnits:    make(map[string]int),
	}

	return &cO
}

func main() {
	var product string
	co := InitCheckout(productPrices)
	for {
		fmt.Println("Enter product name or Enter 'TOTAL' if you want total")
		fmt.Scanf("%s", &product)
		if product == "" {
			fmt.Println("You not entered anythhing please enter product name")
		} else if product != "TOTAL" {
			co.scan(product)
		} else {
			price := co.total()
			fmt.Println("Total sum = $", price)
		}
	}
}
