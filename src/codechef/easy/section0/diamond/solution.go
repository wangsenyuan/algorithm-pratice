package main

import (
	"bufio"
	"fmt"
	"os"
)

var c [][]float64

func init() {
	c = make([][]float64, 2001)
	for i := 0; i < 2001; i++ {
		c[i] = make([]float64, 2002)
	}

	for l := 1; l < 2001; l++ {
		for k := 1; k <= l; k++ {
			c[l][k] = 1.0 - 0.5*(c[l-1][k-1]+c[l-1][k])
		}
	}
}

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
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		v := make([]int, n)
		scanner.Scan()
		bs := scanner.Bytes()
		for x, i := -1, 0; i < n; i++ {
			x = readInt(bs, x+1, &v[i])
		}
		res := solve(n, v)
		fmt.Printf("%.3f\n", res)
	}
}

func solve(n int, v []int) float64 {
	var ans float64

	for i := 1; i <= n; i++ {
		ans += c[n][i] * float64(v[i-1])
	}

	return ans
}
