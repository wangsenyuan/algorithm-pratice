package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scanf("%d %d %d", &a, &b, &n)
	res := solve(a, b, n)
	fmt.Println(res)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}

func solve(a int, b int, n int) string {

	dp := make([]map[int]int, 31)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	var play func(a int, b int) int

	play = func(a int, b int) (res int) {
		if pow(a, b) >= n {
			// already win
			return 1
		}
		if b == 1 && a*a >= n {
			// 这里不能增加b
			if (n-a)%2 == 1 {
				return -1
			}
			return 1
		}
		if a == 1 {
			if pow(2, b) >= n {
				return 0
			}
		}
		// b > 1
		if b < 31 {
			if v, ok := dp[b][a]; ok {
				return v
			}
			defer func() {
				dp[b][a] = res
			}()
		}
		res = -play(a+1, b)
		if res > 0 {
			return
		}
		res = max(res, -play(a, b+1))
		return
	}

	ans := play(a, b)

	if ans == 1 {
		return "Masha"
	}
	if ans == -1 {
		return "Stas"
	}
	return "Missing"
}
