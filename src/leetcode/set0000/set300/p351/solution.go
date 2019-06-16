package main

import "fmt"

func main() {
	fmt.Println(numberOfPatterns(1, 9))
}

//NumberOfPatterns solution
func numberOfPatterns(m int, n int) int {
	if n == 1 {
		return 9
	}

	var count func(last, len, used int) int
	count = func(last, len, used int) int {
		if len == 0 {
			return 1
		}

		sum := 0
		for i := 0; i < 9; i++ {
			if isValid(i, last, used) {
				sum += count(i, len-1, used|(1<<uint(i)))
			}
		}
		return sum
	}
	res := 0
	for l := m; l <= n; l++ {
		a := count(0, l-1, 1)
		res += a * 4
		b := count(1, l-1, 2)
		res += b * 4
		res += count(4, l-1, 1<<4)
	}

	return res
}

func isValid(index, last int, used int) bool {
	if used&(1<<uint(index)) > 0 {
		return false
	}

	if last == -1 {
		return true
	}

	// knight moves or adjacent cells (in a row or in a column)
	if (index+last)%2 == 1 {
		return true
	}

	mid := (index + last) / 2

	if mid == 4 {
		return used&(1<<4) > 0
	}
	if (index%3 != last%3) && (index/3 != last/3) {
		return true
	}

	return used&(1<<uint(mid)) > 0
}
