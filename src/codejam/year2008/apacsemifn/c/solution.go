package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func main() {
	f, err := os.Open("./C-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		scanner.Scan()
		var n int
		pos := readInt(scanner.Bytes(), 0, &n)
		var p float64
		pos = readFloat64(scanner.Bytes(), pos+1, &p)
		var x int
		readInt(scanner.Bytes(), pos+1, &x)
		res := solve(n, p, x)
		fmt.Printf("Case #%d: %.6f\n", i, res)
	}
}

func solve(n int, p float64, x int) float64 {
	dp := make([]float64, 1<<uint(n+1)+1)
	prev := make([]float64, 1<<uint(n+1)+1)
	dp[0] = 0
	dp[1] = 1

	for i := 0; i < n; i++ {
		copy(prev, dp)
		for j := 0; j <= (1 << uint(i+1)); j++ {
			dp[j] = 0
			for k := j / 2; k >= 0 && j-k <= (1<<uint(i)); k-- {
				dp[j] = max(dp[j], p*prev[j-k]+(1.0-p)*prev[k])
			}
		}
	}

	N := 1 << uint(n)
	y := float64(x)/1000000.0*float64(N) + 1e-11
	return dp[int(y)]
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
