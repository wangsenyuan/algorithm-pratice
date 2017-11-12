package main

import "fmt"

func getModifiedArray(length int, updates [][]int) []int {
	array := make([]int, length, length)

	for _, update := range updates {
		i, j, d := update[0], update[1], update[2]
		array[i] += d
		if j < length-1 {
			array[j+1] -= d
		}
	}

	for i := 1; i < length; i++ {
		array[i] += array[i-1]
	}

	return array
}

func main() {
	updates := [][]int{
		[]int{1, 3, 2},
		[]int{2, 4, 3},
		[]int{0, 2, -2},
	}

	result := getModifiedArray(5, updates)
	fmt.Printf("%v\n", result)
}
