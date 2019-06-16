package main

import (
	"math"
	"strconv"
	"fmt"
)

func main() {
	//fmt.Println(smallestGoodBase("13"))
	//fmt.Println(smallestGoodBase("4681"))
	//fmt.Println(smallestGoodBase("1000000000000000000"))
	//fmt.Println(smallestGoodBase("100"))
	fmt.Println(smallestGoodBase("16035713712910627"))
	fmt.Println(smallestGoodBase("470988884881403701"))
}

func smallestGoodBase(s string) string {
	solve := func(n int64, d int) int64 {
		left, right := int64(1), int64(math.Pow(float64(n), 1.0/float64(d))) + 1

		for left <= right {
			mid := left + (right-left)/2
			sum, cur := int64(1), int64(1)
			for i := 1; i <= d; i++ {
				cur *= mid
				sum += cur
			}
			if sum == n {
				return mid
			}
			if sum > n {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return 0
	}

	n, _ := strconv.ParseInt(s, 10, 64)
	x := int64(1)
	for i := 62; i > 0; i-- {
		if x<<uint(i) < n {
			cur := solve(n, i)
			if cur != 0 {
				return strconv.FormatInt(cur, 10)
			}
		}
	}
	return strconv.FormatInt(n-1, 10)

}
