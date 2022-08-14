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
		readString(reader)
		n, m := readTwoNums(reader)
		D := make([][]int, n)
		for i := 0; i < n; i++ {
			D[i] = readNNums(reader, m)
		}
		dist, rank := solve(D)
		buf.WriteString(fmt.Sprintf("%d:", dist))
		for i := 0; i < len(rank); i++ {
			buf.WriteString(fmt.Sprintf(" %d", rank[i]))
		}
		buf.WriteByte('\n')
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

const INF = 1 << 29

func solve(D [][]int) (int, []int) {
	n := len(D)
	m := len(D[0])

	cmp := make([][]int, m)
	for i := 0; i < m; i++ {
		cmp[i] = make([]int, m)
	}

	for r := 0; r < n; r++ {
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				x := D[r][i]
				y := D[r][j]
				x--
				y--
				cmp[y][x]++
			}
		}
	}

	M := 1 << uint(m)

	V := make([]int, M)

	for i := 0; i < M; i++ {
		V[i] = INF
	}

	V[0] = 0

	arr := make([]int, m)

	for state := 1; state < M; state++ {
		var p int
		for i := 0; i < m; i++ {
			if (state>>uint(i))&1 == 1 {
				arr[p] = i
				p++
			}
		}
		if p == 1 {
			V[state] = 0
			continue
		}
		best := INF
		for i := 0; i < p; i++ {
			var tmp int
			x := arr[i]
			// if put x at the last position
			for j := 0; j < p; j++ {
				if i == j {
					continue
				}
				y := arr[j]
				tmp += cmp[y][x]
			}
			tmp += V[state^(1<<uint(x))]
			if tmp < best {
				best = tmp
			}
		}
		V[state] = best
	}

	ans := make([]int, m)
	cur_state := M - 1

	var used int

	for i := m - 1; i >= 0; i-- {
		var p int
		for i := 0; i < m; i++ {
			if (cur_state>>uint(i))&1 == 1 {
				arr[p] = i
				p++
			}
		}
		ans[i] = -1
		for j := 0; j < p; j++ {
			x := arr[j]
			if (used>>uint(x))&1 == 1 {
				continue
			}
			var tmp int
			for k := 0; k < p; k++ {
				if j == k {
					continue
				}
				y := arr[k]
				tmp += cmp[y][x]
			}
			tmp += V[cur_state^(1<<uint(x))]
			if tmp == V[cur_state] {
				if ans[i] < 0 || x < ans[i] {
					ans[i] = x
				}
			}
		}
		cur_state ^= (1 << uint(ans[i]))
		used |= (1 << uint(ans[i]))
		ans[i]++
	}
	return V[M-1], ans
}
