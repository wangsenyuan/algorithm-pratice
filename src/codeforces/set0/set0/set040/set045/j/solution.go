package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	res := solve(n, m)

	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	for i := range n {
		for j := range m {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, m int) [][]int {
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, m)
	}
	if m >= 4 {
		for i := range n {
			first := (i + 1) * m
			second := first - 1
			for j := 0; j < m; j++ {
				if j < m/2 {
					res[i][j] = second
					second -= 2
				} else {
					res[i][j] = first
					first -= 2
				}
			}
		}
	} else if n >= 4 {
		for j := range m {
			first := (j + 1) * n
			second := first - 1
			for i := 0; i < n; i++ {
				if i < n/2 {
					res[i][j] = second
					second -= 2
				} else {
					res[i][j] = first
					first -= 2
				}
			}
		}
	} else {
		res = bruteForce(n, m, res)
	}
	return res
}

func bruteForce(n int, m int, res [][]int) [][]int {
	// n * m <= 9
	for i := range n {
		for j := range m {
			res[i][j] = -1
		}
	}
	w := n * m
	var dfs func(pos int, flag int) bool

	dfs = func(pos int, flag int) bool {
		if pos == w {
			return true
		}
		r := pos / m
		c := pos % m
		for i := 0; i < w; i++ {
			if (flag>>i)&1 == 0 {
				v := i + 1
				if r > 0 && abs(res[r-1][c]-v) == 1 {
					// invalid
					continue
				}
				if c > 0 && abs(res[r][c-1]-v) == 1 {
					continue
				}
				res[r][c] = v

				if dfs(pos+1, flag|(1<<i)) {
					return true
				}

				res[r][c] = -1
			}
		}
		return false
	}

	if !dfs(0, 0) {
		return nil
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
