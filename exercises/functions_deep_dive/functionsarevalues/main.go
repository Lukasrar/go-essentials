package functionsarevalues

import (
	"fmt"
)

type multiplyFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubledNumbers := multiplyNumbers(&numbers, double)
	tripledNumbers := multiplyNumbers(&numbers, triple)

	fmt.Println(doubledNumbers, tripledNumbers)

}

func multiplyNumbers(numbers *[]int, multiplyFunc multiplyFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, multiplyFunc(val))
	}

	return dNumbers
}

func getMultiplierFunction(multiplier int) multiplyFn {

	switch multiplier {
	case 2:
		return double
	case 3:
		return triple

	default:
		return nil
	}

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
