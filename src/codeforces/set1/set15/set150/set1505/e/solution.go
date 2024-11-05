package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _ := readTwoNums(reader)
	cake := make([]string, n)
	for i := 0; i < n; i++ {
		cake[i] = readString(reader)
	}

	res := solve(cake)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')

	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 60

func solve(cake []string) int {
	n := len(cake)
	m := len(cake[0])

	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, m)
		for j := 0; j < m; j++ {
			next[i][j] = -1
		}
	}

	get := func(i int, j int) (int, int) {
		if i == n || j == m {
			return -1, -1
		}
		v := next[i][j]
		return v / m, v % m
	}

	last := n*m - 1
	next[n-1][m-1] = last

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				continue
			}

			if cake[i][j] == '*' {
				next[i][j] = i*m + j
				continue
			}

			a, b := get(i, j+1)
			c, d := get(i+1, j)
			if a < 0 && c < 0 {
				// not found
				next[i][j] = last
			} else if a < 0 {
				next[i][j] = c*m + d
			} else if c < 0 {
				next[i][j] = a*m + b
			} else {
				if a+b <= c+d {
					next[i][j] = a*m + b
				} else {
					next[i][j] = c*m + d
				}
			}
		}
	}

	var ans int
	var x, y int

	for {
		if cake[x][y] == '*' {
			ans++
		}

		if x == n-1 && y == m-1 {
			break
		}
		a, b := get(x, y+1)
		c, d := get(x+1, y)
		if c < 0 || a >= 0 && a+b <= c+d {
			x, y = a, b
		} else {
			x, y = c, d
		}
	}

	return ans
}
