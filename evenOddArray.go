package main

import "fmt"

func main() {
	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	slice := arr[4:14]

	isOdd := func(x int) bool {
		if x%2 == 1 {
			return true
		}
		return false
	}

	isEven := func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	}

	oddNumbersSlice := []int{}
	for _, num := range slice {
		if isOdd(num) {
			oddNumbersSlice = append(oddNumbersSlice, num)
		}
	}

	evenNumbersSlice := []int{}
	for _, num := range slice {
		if isEven(num) {
			evenNumbersSlice = append(evenNumbersSlice, num)
		}
	}

	fmt.Println("Original slice:", slice)
	fmt.Println("Odd numbers slice:", oddNumbersSlice)
	fmt.Println("Even numbers slice:", evenNumbersSlice)
}
