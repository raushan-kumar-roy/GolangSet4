package main

import "fmt"

func powerOf2(n uint8) uint8 {
	return 1 << n
}

func main() {
	inpSlice := make([]uint8, 3, 5)

	inpSlice[0] = 2
	inpSlice[1] = 4
	inpSlice[2] = 8

	powerOf2Slice := make([]uint8, len(inpSlice))

	for i, val := range inpSlice {
		powerOf2Slice[i] = powerOf2(val)
	}

	fmt.Println(powerOf2Slice)

	inpSlice = append(inpSlice, 10, 16)

	powerOf2Slice = make([]uint8, len(inpSlice))

	for i, val := range inpSlice {
		powerOf2Slice[i] = powerOf2(val)
	}

	fmt.Println(powerOf2Slice)
}
