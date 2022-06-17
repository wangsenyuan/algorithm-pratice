package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		r, c, k := readThreeNums(reader)
		res := solve(r, c, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(r, c, k int) int {

	r--
	c--

	var res int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if abs(i-r) <= k && abs(j-c) <= k {
				res++
			}
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func solve1(r, c, k int) int {
	r--
	c--

	dp := make([][]int, 8)
	for i := 0; i < 8; i++ {
		dp[i] = make([]int, 8)
	}
	dp[r][c] = 1

	que := make([]int, 8*8*(k+1))
	var front, end int
	que[end] = r*8*(k+1) + c*(k+1)
	end++

	for front < end {
		cur := que[front]
		front++
		x, y, t := cur/(8*(k+1)), (cur%(8*(k+1)))/(k+1), cur%(k+1)
		if t == k {
			continue
		}

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				u, v := x+i, y+j
				if u >= 0 && u < 8 && v >= 0 && v < 8 && (dp[u][v]>>uint(t+1))&1 == 0 {
					dp[u][v] |= 1 << uint(t+1)
					que[end] = u*8*(k+1) + v*(k+1) + t + 1
					end++
				}
			}
		}
	}

	var res int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if dp[i][j] > 0 {
				res++
			}
		}
	}

	return res
}
