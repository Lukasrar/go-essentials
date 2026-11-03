package main

import "fmt"

// func main() {
// 	var productNames [4]string = [4]string{"A book"}

// 	prices := [4]float64{10.99, 9.99, 45, 99}

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[1])

// 	featurePrices := prices[1:3]

// 	fmt.Println(featurePrices)
// }

func array_slices() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)

	fmt.Println(prices)

	discountPrices := []float64{101.99, 99.99}

	prices = append(prices, discountPrices...)
}
