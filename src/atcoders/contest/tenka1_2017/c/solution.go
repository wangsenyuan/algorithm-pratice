package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
	fmt.Println(res[0], res[1], res[2])
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

const H = 3500

func solve(N int) []int {
	check := func(h int, n int) int {
		a := h + n
		b := n * h
		x := gcd(a, b)
		a /= x
		b /= x
		if N*a >= 4*b {
			return -1
		}
		u := 4*b - N*a
		v := N * b
		g := gcd(u, v)
		u /= g
		v /= g
		if u == 1 {
			return v
		}
		return -1
	}

	// 4 / N = 1/h + 1/n + 1/w
	for h := 1; h <= H; h++ {
		for n := h; n <= H; n++ {
			w := check(h, n)
			if w > 0 {
				return []int{h, n, w}
			}
		}
	}
	return nil
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
