package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFloat(s []byte, from int, num *float64) int {
	i := from
	for i < len(s) && s[i] == ' ' {
		i++
	}
	var d int
	for i < len(s) && s[i] != '.' && s[i] != ' ' {
		d = d*10 + int(s[i]-'0')
		i++
	}
	if i == len(s) || s[i] == ' ' {
		*num = float64(d)
		return i
	}
	var f int
	base := 1
	i++
	for i < len(s) && s[i] != ' ' {
		f = f*10 + int(s[i]-'0')
		base *= 10
		i++
	}
	*num = float64(d) + float64(f)/float64(base)
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n int
		var T float64
		scanner.Scan()
		bs := scanner.Bytes()
		pos := readInt(bs, 0, &n)
		readFloat(bs, pos+1, &T)
		x := make([]float64, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			readFloat(scanner.Bytes(), 0, &x[j])
		}
		fmt.Println(solve(n, x, T))
	}
}

func solve(n int, x []float64, T float64) float64 {
	check := func(D float64) bool {
		y := max(0.0, x[0]-D)
		for i := 1; i < n; i++ {
			if x[i]+D < y+T {
				return false
			}
			y = max(y+T, x[i]-D)
		}
		return true
	}

	left, right, mid := float64(0), 1.0, float64(0.0)

	for !check(right) {
		right *= 2
	}

	eps := 0.00000001
	for left+eps < right {
		mid = left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	return mid
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
