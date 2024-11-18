package main

import "fmt"

func main() {
	var n, p int

	fmt.Scan(&n, &p)

	res := solve(n, p)

	fmt.Println(res)
}

func solve(n int, p int) int {
	check := func(x int, m int) bool {
		// 能否用m个2进制数表示m
		var num int
		for tmp := x; tmp > 0; tmp = tmp & (tmp - 1) {
			num++
		}
		// num是最少需要这么多个数来表示x
		// 同时减少一个高位的数，可以增加两个低一位的数，直到全部变成1
		return num <= m && x >= m
	}

	for m := 1; n-m*p >= 0; m++ {
		if check(n-m*p, m) {
			return m
		}
	}
	return -1
}
