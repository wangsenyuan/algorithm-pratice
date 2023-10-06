package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	res := solve(a)
	fmt.Println(res)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

const X = 300 * 300

func solve(a []int) int {
	fp := make([]int, X+1)
	gp := make([]int, X+1)
	fp[0] = 3
	gp[0] = 3
	var sum int
	p3 := 1
	for _, v := range a {
		sum += v
		for j := sum; j >= 0; j-- {
			fp[j] = add(mul(fp[j], 2), fp[max(j-v, 0)])
			if j >= v {
				gp[j] = add(gp[j], gp[j-v])
			}
		}
		p3 = mul(p3, 3)
	}

	res := sub(p3, fp[(sum+1)/2])
	if sum%2 == 0 {
		res = add(res, gp[sum/2])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
