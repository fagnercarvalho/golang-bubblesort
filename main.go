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
				auxNumber := numbers[i-1]
				numbers[i-1] = numbers[i]
				numbers[i] = auxNumber
				swapped = true
			}
		}
	}

	return numbers
}

func convertToNumbers(parameters []string) []int {
	var numbers []int
	for _, v := range parameters {
		number, _ := strconv.ParseInt(v, 0, 0)
		numbers = append(numbers, int(number))
	}

	return numbers
}
