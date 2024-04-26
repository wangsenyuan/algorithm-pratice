package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = readNNums(reader, n)
	}

	res := solve(n, arr)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, neighbors [][]int) int {
	m := len(neighbors)
	pos := make([][]int, m)

	for i := 0; i < m; i++ {
		pos[i] = make([]int, n)
		for j := 0; j < n; j++ {
			neighbors[i][j]--
			pos[i][neighbors[i][j]] = j
		}
	}

	var res int

	ptr := make([]int, m)

	check := func(r int) bool {
		x := neighbors[0][r]

		for i := 0; i < m; i++ {
			if pos[i][x] != ptr[i]+1 {
				return false
			}
		}
		return true
	}

	advance := func() {
		for i := 0; i < m; i++ {
			ptr[i]++
		}
	}

	for l, r := 0, 0; l < n; l = r {

		// 初始化
		for i := 0; i < m; i++ {
			ptr[i] = pos[i][neighbors[0][l]]
		}
		r = l + 1
		for r < n && check(r) {
			advance()
			r++
		}
		res += (r - l) * (r - l + 1) / 2
	}

	return res
}
