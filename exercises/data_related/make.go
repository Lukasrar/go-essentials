package main

import "fmt"

func main() {
	userNames := make([]string, 2, 3)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Lukas")

	fmt.Println(userNames)
	fmt.Println(cap(userNames))

	courseRatings := map[string]float64{}

	courseRatings["Interestellar"] = 4.7
	courseRatings["Smile"] = 5.0

	fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
