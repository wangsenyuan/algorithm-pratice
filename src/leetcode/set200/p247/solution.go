package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findStrobogrammatic(1))
	fmt.Println(findStrobogrammatic(2))
	fmt.Println(findStrobogrammatic(3))
	fmt.Println(findStrobogrammatic(13))

}

func findStrobogrammatic(n int) []string {

	if n == 1 {
		return []string{"0", "1", "8"}
	}

	var find func(n int, left, right string)
	res := make([]string, 0, int(math.Pow(5.0, float64(n))))

	find = func(n int, left, right string) {
		if n == 0 {
			res = append(res, left+right)
			return
		}
		if n == 1 {
			res = append(res, left+"0"+right)
			res = append(res, left+"1"+right)
			res = append(res, left+"8"+right)
			return
		}

		find(n-2, left+"0", "0"+right)
		find(n-2, left+"1", "1"+right)
		find(n-2, left+"8", "8"+right)
		find(n-2, left+"6", "9"+right)
		find(n-2, left+"9", "6"+right)
	}
	find(n-2, "1", "1")
	find(n-2, "8", "8")
	find(n-2, "6", "9")
	find(n-2, "9", "6")

	return res
}
