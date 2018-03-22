package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		solve()
		t--
	}
}

func solve() {
	var n int
	fmt.Scanf("%d", &n)
	levels := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &levels[i])
	}

	valid := true

	if levels[0] == -1 {
		levels[0] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if levels[i+1] > 1 {
			if levels[i] == - 1 {
				levels[i] = levels[i+1] - 1
			}
			if levels[i] != levels[i+1]-1 {
				valid = false
				break
			}
		}
	}

	if !valid || levels[0] != 1 {
		println(0)
		return
	}

	mod := int64(1000000007)

	ans := int64(1)

	for i := 0; i < n; i++ {
		if levels[i] == -1 {
			ans = (ans * 2) % mod
		}
	}
	println(ans)
}
