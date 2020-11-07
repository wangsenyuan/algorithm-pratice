package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	A := readNNums(reader, n)

	res := solve(n, A)

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

const MX = 81

const dx, dy, dz = MX + 1, MX + 1, MX*(MX+1)/2 + 1

var dp [dx][dy][dz]int

const INF = 100000000

func init() {
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			for k := 0; k < dz; k++ {
				dp[i][j][k] = INF
			}
		}
	}
}

func solve(n int, A []int) []int {
	gg := make([]int, 0, n/2)

	for i := 0; i <= n; i++ {
		j := i
		for i < n && A[i] == 0 {
			i++
		}
		gg = append(gg, i-j)
	}

	p := make([]int, len(gg))
	p[0] = gg[0]
	for i := 1; i < len(gg); i++ {
		p[i] = gg[i] + p[i-1]
	}
	dp[0][0][0] = 0

	k := len(gg)
	l := p[len(p)-1]

	N := n * (n - 1) / 2

	for i := 0; i < k; i++ {
		for j := 0; j <= l; j++ {
			for s := 0; s <= N; s++ {
				if dp[i][j][s] == INF {
					continue
				}
				for q := j; q <= l; q++ {
					min(&dp[i+1][q][s+abs(q-p[i])], dp[i][j][s]+(q-j)*(q-j))
				}
			}
		}
	}

	res := make([]int, 0, N+1)
	cur := INF
	for s := 0; s <= N; s++ {
		min(&cur, dp[k][l][s])
		val := l*l - cur
		res = append(res, val/2)
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a *int, b int) {
	if *a > b {
		*a = b
	}
}
