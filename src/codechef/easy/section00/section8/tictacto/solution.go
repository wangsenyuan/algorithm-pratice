package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		m, n, k := readThreeNums(reader)
		cmds := make([][]int, m*n)
		for i := 0; i < len(cmds); i++ {
			cmds[i] = readNNums(reader, 2)
		}
		res := solve(m, n, k, cmds)

		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

const Alice = "Alice"
const Bob = "Bob"

const Draw = "Draw"

func solve(m, n, k int, cmds [][]int) string {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}

	check := func(d int) bool {
		for i := 0; i < d; i++ {
			x, y := cmds[i][0], cmds[i][1]
			x--
			y--
			if i&1 == 0 {
				grid[x][y]++
			} else {
				grid[x][y]--
			}
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				sum[i][j] = grid[i][j]
				if i > 0 {
					sum[i][j] += sum[i-1][j]
				}
				if j > 0 {
					sum[i][j] += sum[i][j-1]
				}
				if i > 0 && j > 0 {
					sum[i][j] -= sum[i-1][j-1]
				}
			}
		}
		var flag bool
		for i := k - 1; i < m && !flag; i++ {
			for j := k - 1; j < n && !flag; j++ {
				tmp := sum[i][j]
				if i-k >= 0 {
					tmp -= sum[i-k][j]
				}
				if j-k >= 0 {
					tmp -= sum[i][j-k]
				}
				if i-k >= 0 && j-k >= 0 {
					tmp += sum[i-k][j-k]
				}
				if abs(tmp) == k*k {
					flag = true
				}
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				grid[i][j] = 0
				sum[i][j] = 0
			}
		}

		return flag
	}

	left, right := 0, len(cmds)

	if !check(right) {
		return Draw
	}

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if right&1 == 0 {
		return Bob
	}
	return Alice
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n

	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
