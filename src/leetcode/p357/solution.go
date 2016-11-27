package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	last := 9
	res := 10

	for i := 2; i <= n; i++ {
		last *= 9 - i + 2
		res += last
	}

	return res
}

func countNumbersWithUniqueDigits1(n int) int {
	if n == 0 {
		return 1
	}

	var count func(steps, used int, first bool)

	var res = 0
	count = func(steps, used int, first bool) {
		res++
		if steps == n {
			return
		}

		for i := 0; i < 10; i++ {
			if i == 0 && first || used&(1<<uint(i)) > 0 {
				continue
			}
			count(steps+1, used|(1<<uint(i)), false)
		}
	}

	count(0, 0, true)
	return res
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits1(2))
}
