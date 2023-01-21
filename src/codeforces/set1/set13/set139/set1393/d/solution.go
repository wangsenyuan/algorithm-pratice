package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	m, n := readTwoNums(reader)

	G := make([]string, m)

	for i := 0; i < m; i++ {
		G[i], _ = reader.ReadString('\n')
	}

	fmt.Println(solve(m, n, G))
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

func solve(m, n int, G []string) int {
	up := make([][]int, m)
	down := make([][]int, m)

	for i := 0; i < m; i++ {
		up[i] = make([]int, n)
		down[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && G[i-1][j] == G[i][j] {
				up[i][j] = up[i-1][j] + 1
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i < m-1 && G[i+1][j] == G[i][j] {
				down[i][j] = down[i+1][j] + 1
			}
		}
	}

	L := make([]int, n)
	R := make([]int, n)

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			jj := j
			for j < n && G[i][j] == G[i][jj] {
				j++
			}
			j--
			L[jj] = jj
			for k := jj + 1; k <= j; k++ {
				L[k] = max(L[k-1], k-min(up[i][k], down[i][k]))
			}
			R[j] = j

			for k := j - 1; k >= jj; k-- {
				R[k] = min(R[k+1], k+min(up[i][k], down[i][k]))
			}

			for k := jj; k <= j; k++ {
				ans += min(k-L[k]+1, R[k]-k+1)
			}
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
