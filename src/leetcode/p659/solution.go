package main

import "fmt"

func isPossible(nums []int) bool {
	n := len(nums)

	if n < 3 {
		return false
	}

	set := make(map[int][]int)

	for _, num := range nums {

		if arr, found := set[num-1]; found {
			j := 0
			for j < len(arr) && arr[j] >= 3 {
				j++
			}
			cnt := 0
			if j == len(arr) {
				cnt = arr[0]
				arr = arr[1:]
			} else {
				cnt = arr[j]
				copy(arr[j:], arr[j+1:])
				arr = arr[:len(arr)-1]
			}

			if len(arr) > 0 {
				set[num-1] = arr
			} else {
				delete(set, num-1)
			}

			if arr, found := set[num]; found {
				set[num] = append(arr, cnt+1)
			} else {
				set[num] = []int{cnt + 1}
			}
		} else {
			if arr, found := set[num]; found {
				set[num] = append(arr, 1)
			} else {
				set[num] = []int{1}
			}
		}
	}

	//fmt.Println(set)

	for _, arr := range set {
		for _, cnt := range arr {
			if cnt < 3 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 4, 4, 5}))
}
