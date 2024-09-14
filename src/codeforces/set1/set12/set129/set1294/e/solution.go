package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(reader, m)
	}

	res := solve(grid)

	fmt.Println(res)
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

func solve(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	arr := make([]int, 2*n)
	cnt := make([]int, 2*n+1)
	process := func(a []int, c int) int {
		for i := 0; i < 2*n; i++ {
			arr[i] = a[i%n] - 1 - c
		}
		// 需要的结果是 i * m + c, 把c都减去后 那么就是希望 i * m
		// 假设shift了i次，那么就是在区间 i...i+n中间
		// 有多少个 数 v = (j - i) * m
		// 可以求diff
		clear(cnt)
		for i := 0; i < 2*n; i++ {
			if arr[i]%m == 0 {
				x := arr[i] / m
				// 比如 arr[i] = 6, m = 3
				// 那么 arr[i] 可以是第2个元素
				if i-x >= 0 && x < n {
					// 这个的起点是定的就是 i-x
					cnt[i-x]++
				}
			}
		}
		res := n
		for i := 0; i < n; i++ {
			res = min(res, i+n-cnt[i])
		}
		return res
	}

	col := make([]int, n)

	var res int
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			col[i] = grid[i][j]
		}
		res += process(col, j)
	}

	return res
}
