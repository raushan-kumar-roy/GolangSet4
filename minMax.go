package main

import "fmt"

func main() {
	numbers := []int{3, 5, 1, 7, 2, 8, 4, 6}

	max := numbers[0]
	nextHighest := numbers[0]
	min := numbers[0]
	nextLowest := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			nextHighest = max
			max = numbers[i]
		} else if numbers[i] > nextHighest {
			nextHighest = numbers[i]
		}
		if numbers[i] < min {
			nextLowest = min
			min = numbers[i]
		} else if numbers[i] < nextLowest {
			nextLowest = numbers[i]
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Next highest value:", nextHighest)
	fmt.Println("Min:", min)
	fmt.Println("Next lowest value:", nextLowest)
}
