package main

import (
	"bufio"
	"fmt"
	"math"
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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func readNNums2(bs []byte, n int) []int {
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		dishs := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			bs := scanner.Bytes()
			var m int
			pos := readInt(bs, 0, &m)
			dishs[i] = readNNums2(bs[pos+1:], m)
		}
		fmt.Println(solve(n, dishs))
		t--
	}
}

func solve(n int, dishs [][]int) int64 {
	S := make([][]int64, n)
	for i := 0; i < n; i++ {
		S[i] = make([]int64, len(dishs[i]))
	}

	for i := 1; i < n; i++ {
		var a int64 = math.MinInt64
		var b int64 = math.MinInt64

		m := len(dishs[i-1])
		for j := 0; j < m; j++ {
			k := (j + 1) % m
			a = max2(a, S[i-1][k]-int64(i)*int64(dishs[i-1][j]))
			b = max2(b, S[i-1][k]+int64(i)*int64(dishs[i-1][j]))
		}
		m = len(dishs[i])
		for j := 0; j < m; j++ {
			x := int64(i) * int64(dishs[i][j])
			S[i][j] = max2(S[i][j], max2(a+x, b-x))
		}
	}

	var ans int64

	for i := 0; i < len(dishs[n-1]); i++ {
		ans = max2(ans, S[n-1][i])
	}
	return ans
}
func solve1(n int, dishs [][]int) int64 {
	g := make([]int64, n)
	h := make([]int64, n)

	for i := 0; i < n; i++ {
		g[i] = math.MaxInt64
		h[i] = math.MinInt64
	}

	for i := 0; i < len(dishs[0]); i++ {
		h[0] = max2(h[0], int64(dishs[0][i]))
		g[0] = min2(g[0], int64(dishs[0][i]))
	}
	var ans int64
	for d := 1; d < n; d++ {
		m := len(dishs[d])
		for i := 0; i < m; i++ {
			//i is the end, and j is the front
			j := (i + 1) % m
			s := int64(d) * int64(dishs[d][j])
			x := s - g[d-1]
			y := h[d-1] - s

			z := max2(x, y)

			ans = max2(ans, z)

			s2 := int64(d+1) * int64(dishs[d][i])
			g[d] = min2(g[d], s2-z)
			h[d] = max2(h[d], s2+z)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
