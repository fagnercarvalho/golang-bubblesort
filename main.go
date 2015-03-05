package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := convertToNumbers(os.Args[1:])
	fmt.Println(bubbleSort(numbers))
}

func bubbleSort(numbers []int) []int {
	n := len(numbers)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i <= n-1; i++ {
			if numbers[i-1] > numbers[i] {
				numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
				swapped = true
			}
		}
	}

	return numbers
}

func convertToNumbers(parameters []string) []int {
	var numbers []int
	for _, parameter := range parameters {
		number, _ := strconv.Atoi(parameter)
		numbers = append(numbers, number)
	}

	return numbers
}
