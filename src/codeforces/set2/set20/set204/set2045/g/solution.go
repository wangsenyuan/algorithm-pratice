package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []string {
	r, c, x := readThreeNums(reader)
	h := make([]string, r)
	for i := range r {
		h[i] = readString(reader)[:c]
	}
	m := readNum(reader)
	q := make([][]int, m)
	for i := range m {
		q[i] = readNNums(reader, 4)
	}
	return solve(h, x, q)
}

const inf = 1 << 60

func solve(h []string, x int, q [][]int) []string {
	n := len(h)
	m := len(h[0])

	get := func(x byte) int {
		return int(x - '0')
	}

	calc := func(a, b, c, d int) int {
		diff := get(h[a][b]) - get(h[c][d])
		return pow(diff, x)
	}

	for i := 0; i+2 <= n; i++ {
		for j := 0; j+2 <= m; j++ {
			tmp := calc(i, j, i, j+1)
			tmp += calc(i, j+1, i+1, j+1)
			tmp += calc(i+1, j+1, i+1, j)
			tmp += calc(i+1, j, i, j)
			if tmp != 0 {
				return invalid(q)
			}
		}
	}
	// valid
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if j > 0 {
				dist[i][j] = dist[i][j-1] + calc(i, j-1, i, j)
			} else {
				dist[i][j] = dist[i-1][j] + calc(i-1, j, i, j)
			}
		}
	}

	ans := make([]string, len(q))

	for i, cur := range q {
		a, b, c, d := cur[0]-1, cur[1]-1, cur[2]-1, cur[3]-1
		ans[i] = fmt.Sprintf("%d", dist[c][d]-dist[a][b])
	}

	return ans
}

var dd []int = []int{-1, 0, 1, 0, -1}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	res := pow(a, b/2)
	res = res * res
	if b&1 == 1 {
		res *= a
	}
	return res
}

func invalid(q [][]int) []string {
	ans := make([]string, len(q))
	for i := range len(q) {
		ans[i] = "INVALID"
	}
	return ans
}
