package main

import "fmt"

func main() {

	places := []string{"New York", "Paris", "London", "Sydney", "Tokyo", "Moscow", "Toronto"}

	lengths := processArray(places)
	fmt.Println(lengths)
}

func processArray(places []string) []int {
	lengths := []int{}

	for _, place := range places {
		lengths = append(lengths, len(place))
	}

	return lengths
}
