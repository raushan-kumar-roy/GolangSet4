package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sequence of comma-separated numbers: ")
	input, _ := reader.ReadString('\n')

	numbers := strings.Split(input, ",")

	for i, num := range numbers {
		numbers[i] = strings.TrimSpace(num)
	}

	fmt.Println(numbers)
}
