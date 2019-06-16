package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(20))
	/*for i := 1; i < 20; i++ {
		fmt.Println(nthUglyNumber(i))
	}*/
}

func nthUglyNumber(n int) int {
	nums := make([]int, n)
	i, j, k := 0, 0, 0
	a, b, c := 2, 3, 5
	nums[0] = 1

	for x := 1; x < n; x++ {
		d := min(a, b, c)
		nums[x] = d
		if d == a {
			i++
			a = 2 * nums[i]
		}
		if d == b {
			j++
			b = 3 * nums[j]
		}
		if d == c {
			k++
			c = 5 * nums[k]
		}
	}

	return nums[n-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
