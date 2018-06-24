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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		scanner.Scan()
		s := scanner.Bytes()
		scanner.Scan()
		t := scanner.Bytes()

		ln, cnt := solve(s, t)

		fmt.Printf("%d %d\n", ln, cnt)

		tc--

		if tc > 0 {
			scanner.Scan()
		}
	}
}

const MOD = 23102009

func solve(s, t []byte) (int, int64) {
	m := len(s)
	n := len(t)
	f := make([][]int, m+1)
	g := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
		g[i] = make([]int64, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				if f[i-1][j] == f[i][j-1] {
					f[i][j] = f[i-1][j]
				} else if f[i-1][j] > f[i][j-1] {
					f[i][j] = f[i-1][j]
				} else {
					f[i][j] = f[i][j-1]
				}
			}
		}
	}

	var count func(m, n int) int64

	count = func(m, n int) int64 {
		if m == 0 || n == 0 {
			return 1
		}
		if g[m][n] > 0 {
			return g[m][n]
		}

		if s[m-1] == t[n-1] {
			g[m][n] = count(m-1, n-1)
			return g[m][n]
		}

		a, b, c := f[m-1][n-1], f[m-1][n], f[m][n-1]

		if b > c {
			g[m][n] = count(m-1, n)
		} else if b < c {
			g[m][n] = count(m, n-1)
		} else if a == b {
			g[m][n] = ((count(m-1, n)+count(m, n-1))%MOD - count(m-1, n-1) + MOD) % MOD
		} else {
			g[m][n] = (count(m-1, n) + count(m, n-1)) % MOD
		}

		return g[m][n]
	}

	return f[m][n], count(m, n)
}
