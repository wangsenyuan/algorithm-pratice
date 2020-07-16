package main

import (
	"bufio"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		B := make([][]byte, n)
		for i := 0; i < n; i++ {
			B[i], _ = reader.ReadBytes('\n')
		}
		fmt.Println(string(solve(n, m, B)))
	}
}

func solve(n, m int, B [][]byte) []byte {
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}

	vis[n-1][m-1] = true

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				continue
			}
			if B[i][j] == '#' {
				continue
			}
			if j < m-1 {
				vis[i][j] = vis[i][j+1]
			}
			if i < n-1 && !vis[i][j] {
				vis[i][j] = vis[i+1][j]
			}
		}
	}

	// vis[0][0] = true

	que := make([]int, n*m)
	var front, end int
	que[end] = 0
	end++
	buf := make([]byte, n+m-1)
	buf[0] = B[0][0]
	next := make([]bool, 26)

	for p := 1; p < n+m-1; p++ {

		for j := 0; j < 26; j++ {
			next[j] = false
		}

		for i := front; i < end; i++ {
			cur := que[i]
			x, y := cur/m, cur%m

			if x < n-1 && vis[x+1][y] {
				next[int(B[x+1][y]-'a')] = true
			}
			if y < m-1 && vis[x][y+1] {
				next[int(B[x][y+1]-'a')] = true
			}
		}

		var pivot int

		for pivot < 26 && !next[pivot] {
			pivot++
		}
		buf[p] = byte(pivot + 'a')
		// pivot < 26
		tmp := end
		for i := front; i < tmp; i++ {
			cur := que[i]
			x, y := cur/m, cur%m
			if x < n-1 && vis[x+1][y] && int(B[x+1][y]-'a') == pivot {
				que[end] = (x+1)*m + y
				end++
			}

			if y < m-1 && vis[x][y+1] && int(B[x][y+1]-'a') == pivot {
				que[end] = x*m + y + 1
				end++
			}
		}

		front = tmp
	}

	return buf
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
