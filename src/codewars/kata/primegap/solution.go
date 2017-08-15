package main

import "fmt"

func Gap(g, m, n int) []int {
	nums := make([]bool, n+1)

	x := 2

	for x <= n {
		if !nums[x] {
			y := x - g
			if y >= m && !nums[y] {
				i := y + 1
				for i < x && nums[i] {
					i++
				}
				if i == x {
					return []int{y, x}
				}
			}
			for i := 2 * x; i <= n; i += x {
				nums[i] = true
			}
		}

		x++
	}

	return nil
}

func main() {
	fmt.Println(Gap(2, 100, 110))
	fmt.Println(Gap(4, 100, 110))
	fmt.Println(Gap(6, 100, 110))
	fmt.Println(Gap(8, 300, 400))
	fmt.Println(Gap(10, 300, 400))
}
