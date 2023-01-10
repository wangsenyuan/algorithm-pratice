package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	field := make([]string, n)
	for i := 0; i < n; i++ {
		field[i] = readString(reader)[:m]
	}
	res := solve(field)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(field []string) int64 {

	n := len(field)
	m := len(field[0])
	row := make([][]int, n)
	col := make([][]int, n)

	for i := 0; i < n; i++ {
		row[i] = make([]int, m)
		col[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == '*' {
				row[i][j]++
				col[i][j]++
			}
			if j > 0 {
				row[i][j] += row[i][j-1]
			}
			if i > 0 {
				col[i][j] += col[i-1][j]
			}
		}
	}
	var res int64
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == '*' {
				// take it as top-left corner
				a := row[i][m-1] - row[i][j]
				b := col[n-1][j] - col[i][j]
				res += int64(a) * int64(b)
				// take it as top-right corner
				if j > 0 {
					a = row[i][j-1]
					b = col[n-1][j] - col[i][j]
					res += int64(a) * int64(b)
				}
				// take it as bot-left corner
				if i > 0 {
					a = row[i][m-1] - row[i][j]
					b = col[i-1][j]
					res += int64(a) * int64(b)
				}
				// take it as bot-right corner
				if i > 0 && j > 0 {
					a = row[i][j-1]
					b = col[i-1][j]
					res += int64(a) * int64(b)
				}
			}
		}
	}
	return res
}
