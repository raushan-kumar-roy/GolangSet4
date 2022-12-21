package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ComputeAverage(numbers []float64) float64 {
	sum := 0.0

	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

func ComputeStandardDeviation(numbers []float64) float64 {
	sumOfSquares := 0.0

	average := ComputeAverage(numbers)

	for _, num := range numbers {
		sumOfSquares += (num - average) * (num - average)
	}

	return math.Sqrt(sumOfSquares / float64(len(numbers)))
}

func main() {
	var numbers []float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a list of floating point numbers (enter 'q' or 'Q' to stop):")
	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "q" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid floating point number.")
			continue
		}

		numbers = append(numbers, num)
	}

	average := ComputeAverage(numbers)
	standardDeviation := ComputeStandardDeviation(numbers)
	fmt.Printf("Average: %f\n", average)
	fmt.Printf("Standard deviation: %f\n", standardDeviation)
}
