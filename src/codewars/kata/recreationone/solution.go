package main

import (
	"math"
	"fmt"
)

func ListSquared(m, n int) [][]int {

	process := func(x int) (bool, int) {
		sx := int(math.Sqrt(float64(x)))

		sum := 0

		for a := 1; a <= sx; a++ {
			if x%a == 0 {
				sum += a * a
				b := x / a
				if a != b {
					sum += b * b
				}
			}
		}

		sy := int(math.Sqrt(float64(sum)))

		if sy*sy == sum {
			return true, sum
		}
		return false, -1
	}

	res := make([][]int, 0, n-m)
	for x := m; x <= n; x++ {
		ok, y := process(x)
		if ok {
			res = append(res, []int{x, y})
		}
	}
	return res
}

func main() {
	fmt.Println(ListSquared(42, 42))
	fmt.Println(ListSquared(1, 250))
	fmt.Println(ListSquared(250, 500))
	fmt.Println(ListSquared(300, 600))

}
