package main

import (
	"strconv"
	"fmt"
)

func Solution(list []int) string {
	if len(list) == 0 {
		return ""
	}

	var res string

	j := 0

	for i := 1; i <= len(list); i++ {
		if i < len(list) && list[i] == 1+list[i-1] {
			continue
		}

		if i-1 == j {
			res += "," + strconv.Itoa(list[j])
		} else if i-1 == j+1 {
			res += "," + strconv.Itoa(list[j]) + "," + strconv.Itoa(list[j+1])
		} else {
			res += "," + fmt.Sprintf("%d-%d", list[j], list[i-1])
		}
		j = i
	}

	return res[1:]
}

func main() {
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}
