package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n, x int
		scanner.Scan()
		bs := scanner.Bytes()
		readInt(bs, readInt(bs, 0, &n)+1, &x)
		res := solve(n, x)
		fmt.Println(res)
	}
}

func solve(n int, x int) int {
	m := calculate(n, x)
	k := len(m)
	pre := make([]int, k)
	for i := 0; i < k; i++ {
		pre[i] = m[i]
		if i > 0 {
			pre[i] += pre[i-1]
		}
	}
	y := make([]int, 2*n)
	j := 0
	a, b, c := 0, 0, 0
	for {
		y[j] = pre[a] + c
		if b > 0 {
			y[j] -= pre[b-1]
		}
		c = y[j] / 10
		y[j] %= 10
		j++
		if a < k-1 {
			a++
		}
		if a >= n {
			b++
		}

		if a < b {
			break
		}
	}

	var MOD = 1000000007
	var ans int
	base := 1
	for i := j - 1; i >= 0; i-- {
		tmp := (y[i] * base) % MOD
		ans = (ans + tmp) % MOD
		base = (base * 23) % MOD
	}

	return ans
}

func calculate(n int, x int) []int {
	k := n
	if x > 3 {
		k++
	}
	m := make([]int, k)
	carray := 0
	for i := 0; i < n; i++ {
		y := x*x + carray
		m[i] = y % 10
		carray = y / 10
	}
	if carray > 0 {
		m[k-1] = carray
	}
	return m
}
