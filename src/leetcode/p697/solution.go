package main

import "fmt"

func findShortestSubArray(nums []int) int {
	cnts := make(map[int]int)
	start := make(map[int]int)
	end := make(map[int]int)

	most := 0

	for i, num := range nums {
		cnts[num]++
		if _, found := start[num]; !found {
			start[num] = i
		}
		end[num] = i
		if cnts[num] > most {
			most = cnts[num]
		}
	}

	res := len(nums)

	for num, cnt := range cnts {
		if most == cnt {
			i, j := start[num], end[num]
			if j-i+1 < res {
				res = j - i + 1
			}
		}
	}

	return res
}

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
