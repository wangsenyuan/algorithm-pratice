package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)[:m]
	}

	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(grid, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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
func solve(grid []string, queries [][]int) []int {
	n := len(grid)
	m := len(grid[0])
	buf := make([]byte, n*m)
	var sum int
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			buf[j*n+i] = grid[i][j]
			if grid[i][j] == '*' {
				sum++
			}
		}
	}
	res := count(buf[:sum], '.')
	ans := make([]int, len(queries))

	for i, cur := range queries {
		x, y := cur[0], cur[1]
		p := (y-1)*n + x - 1
		if buf[p] == '*' {
			if p < sum {
				res++
			}
			buf[p] = '.'
		} else {
			if p < sum {
				res--
			}
			buf[p] = '*'
		}
		if buf[p] == '*' {
			if buf[sum] == '.' {
				res++
			}
			sum++
		} else {
			sum--
			if buf[sum] == '.' {
				res--
			}
		}
		ans[i] = res
	}
	return ans
}

func count(s []byte, expect byte) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == expect {
			res++
		}
	}
	return res
}
