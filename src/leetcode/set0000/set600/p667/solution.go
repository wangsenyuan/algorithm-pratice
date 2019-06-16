package main

import "fmt"

func constructArray(n int, k int) []int {
	arr := make([]int, 0, n)

	i, j := 1, n

	for i <= j {
		if k%2 == 1 {
			arr = append(arr, i)
			i++
		} else {
			arr = append(arr, j)
			j--
		}
		if k > 1 {
			k--
		}
	}

	return arr
}

func isCorrect(input []int, k int) bool {
	set := make(map[int]bool)

	for i := 0; i < len(input)-1; i++ {
		set[abs(input[i+1]-input[1])] = true
	}

	fmt.Println(len(set))
	return len(set) == k
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(constructArray(6, 5))
	fmt.Println(constructArray(6, 4))
	fmt.Println(constructArray(6, 3))
	fmt.Println(constructArray(6, 2))
	fmt.Println(constructArray(6, 1))
	fmt.Println(constructArray(9999, 9998))
	fmt.Println(isCorrect(constructArray(9999, 9998), 9998))
}
