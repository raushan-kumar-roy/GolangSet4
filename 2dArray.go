package main

import "fmt"

func main() {
	array := make([][]int, 5)
	for i := range array {
		array[i] = make([]int, 2)
	}

	for i := 0; i < 5; i++ {
		array[i][0] = 10 + i*3
		array[i][1] = 2 * array[i][0]
	}

	fmt.Println(array)
}
