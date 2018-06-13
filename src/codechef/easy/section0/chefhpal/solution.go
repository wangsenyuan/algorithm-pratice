package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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
		scanner.Scan()
		bs := scanner.Bytes()
		var n, a int
		readInt(bs, readInt(bs, 0, &n)+1, &a)
		l, s := solve(n, a)
		fmt.Printf("%d %s\n", l, s)
	}
}

func solve(n int, a int) (int, string) {
	if n == 1 {
		return 1, "a"
	}
	if a == 1 {
		return n, repeat("a", n)
	}

	if a > 2 {
		x, y := n/3, n%3
		s := repeat("abc", x)
		t := "abc"
		return 1, s + t[:y]
	}
	if n == 2 {
		return 1, "ab"
	}

	if n == 3 {
		return 2, "abb"
	}

	if n >= 9 {
		x, y := n/6, n%6
		s := repeat("aabbab", x)
		t := "aabbab"
		return 4, s + t[:y]
	}

	return brute(n)
}

func brute(n int) (int, string) {
	res := make([]byte, n)

	var gen func(i int)

	var best string
	var min int
	gen = func(i int) {
		if i == n {
			tmp := countPalindrome(res)
			if len(best) == 0 || tmp < min {
				best = string(res)
				min = tmp
			}
			return
		}
		res[i] = 'a'
		gen(i + 1)
		res[i] = 'b'
		gen(i + 1)
	}

	gen(0)

	return min, best
}

func repeat(seed string, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += seed
	}
	return res
}

func countPalindrome(bs []byte) int {
	n := len(bs)
	var cnt int

	for i := 0; i < n; i++ {
		j, k := i-1, i+1
		for j >= 0 && k < n && bs[j] == bs[k] {
			j, k = j-1, k+1
		}
		if k-j-1 > cnt {
			cnt = k - j - 1
		}
		j, k = i, i+1
		for j >= 0 && k < n && bs[j] == bs[k] {
			j, k = j-1, k+1
		}
		if k-j-1 > cnt {
			cnt = k - j - 1
		}
	}
	return cnt
}
