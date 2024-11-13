package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	res := solve(x, n)

	fmt.Println(res)
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return (a % mod) * (b % mod) % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(x int, n int) int {
	pfs := getPrimeFactors(x)

	get := func(p int) int {
		pw := p
		for pw <= n/p {
			pw *= p
		}

		res := 1
		var cnt int

		for pw > 1 {
			// 能整除9，27，81等的，也可以整除3
			tmp := n/pw - cnt
			res = mul(res, pow(pw, tmp))
			cnt = n / pw
			pw /= p
		}

		return res
	}

	res := 1

	for _, p := range pfs {
		res = mul(res, get(p))
	}

	return res
}

func getPrimeFactors(num int) []int {
	var res []int

	for x := 2; x <= num/x; x++ {
		if num%x == 0 {
			res = append(res, x)
			for num%x == 0 {
				num /= x
			}
		}
	}
	if num > 1 {
		res = append(res, num)
	}
	return res
}
