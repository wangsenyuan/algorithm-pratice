package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	res := solve(l, r)
	fmt.Println(res)
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a int, b ...int) int {
	res := a % mod
	for _, x := range b {
		res = x % mod * res % mod
	}
	return res
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

func inv(a int) int {
	return pow(a, mod-2)
}

func solve(l int, r int) int {
	get := func(pos int) (p1 int, p2 int) {
		var cnt int
		w := 1
		for cnt < pos {
			if cnt+w >= pos {
				p1 += pos - cnt
				break
			}
			p1 += w
			cnt += w
			w *= 2

			if cnt+w >= pos {
				p2 += pos - cnt
				break
			}
			p2 += w
			cnt += w
			w *= 2
		}
		return
	}
	i2 := inv(2)
	l1, l2 := get(l - 1)
	r1, r2 := get(r)
	// res := (2*l1 + 1 + 2*r1 + 1) * (r1 - l1 + 1) / 2
	res1 := mul(2*l1+1+2*r1-1, r1-l1, i2)
	// res += (2*l2 + 2*r2) * (r2 - l2 + 1) / 2
	res2 := mul(2*l2+2+2*r2, r2-l2, i2)
	return add(res1, res2)
}
