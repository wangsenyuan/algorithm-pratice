package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	n, m := readTwoNums(bufio.NewReader(os.Stdin))
	telports, path := solve(n, m)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(telports)))
	for _, t := range telports {
		buf.WriteString(fmt.Sprintf("%d %d %d %d\n", t[0], t[1], t[2], t[3]))
	}
	for _, p := range path {
		buf.WriteString(fmt.Sprintf("%d %d\n", p[0], p[1]))
	}
	fmt.Print(buf.String())
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

func solve(n int, m int) (telports [][]int, path [][]int) {
	if n*m == 2 {
		path = [][]int{{1, 1}, {n, m}, {1, 1}}
		return
	}

	if m > 1 && n%2 == 0 {
		return solve1(n, m)
	}
	if n > 1 && m%2 == 0 {
		return solve2(n, m)
	}
	telports = [][]int{{n, m, 1, 1}}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := range m {
				path = append(path, []int{i + 1, j + 1})
			}
		} else {
			for j := m - 1; j >= 0; j-- {
				path = append(path, []int{i + 1, j + 1})
			}
		}
	}
	path = append(path, []int{1, 1})
	return
}

func solve2(n int, m int) (telports [][]int, path [][]int) {
	// m is even
	var x int
	for j := 0; j < m; j++ {
		if j%2 == 0 {
			for i := x; i < n; i++ {
				path = append(path, []int{i + 1, j + 1})
			}
			x = 1
		} else {
			for i := n - 1; i > 0; i-- {
				path = append(path, []int{i + 1, j + 1})
			}
			if j == m-1 {
				for r := m - 1; r >= 0; r-- {
					path = append(path, []int{1, r + 1})
				}
			}
		}
	}

	return
}

func solve1(n int, m int) (telports [][]int, path [][]int) {
	// n is even
	var x int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := x; j < m; j++ {
				path = append(path, []int{i + 1, j + 1})
			}
			x = 1
		} else {
			for j := m - 1; j > 0; j-- {
				path = append(path, []int{i + 1, j + 1})
			}
			if i == n-1 {
				for r := n - 1; r >= 0; r-- {
					path = append(path, []int{r + 1, 1})
				}
			}
		}
	}
	return
}
