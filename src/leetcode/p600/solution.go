package main

import "fmt"

func main() {
	fmt.Println(findIntegers(5))
}

func findIntegers(num int) int {
	f := make([]int, 32)

	f[0] = 1
	f[1] = 2
	f[2] = 3

	for i := 3; i < 32; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	ans := 0

	for i, prev := 30, 1; i >= 0; i-- {
		if num&(1<<uint(i)) != 0 {
			ans += f[i]

			if prev == 1 {
				ans--
				break
			}
			prev = 1
		} else {
			prev = 0
		}
	}

	return ans + 1
}
