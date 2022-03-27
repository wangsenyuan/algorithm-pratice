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
		n, m := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 4)
		}
		res := solve(n, E, Q)
		for _, num := range res {
			buf.WriteString(fmt.Sprintf("%d\n", num))
		}
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

func solve(n int, E [][]int, Q [][]int) []int64 {
	deg := make([]int, n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		deg[u]++
		deg[v]++
	}

	var sum int64

	cost := func(d int) int64 {
		return int64(d) * int64(d+1) / 2
	}

	for i := 0; i < n; i++ {
		sum += cost(deg[i])
	}

	sum -= int64(n - 1)

	ans := make([]int64, 1+len(Q))
	ans[0] = sum
	for i, q := range Q {
		a, b, c, d := q[0], q[1], q[2], q[3]
		a--
		b--
		c--
		d--
		tmp := sum
		// deg[a]--, deg[b]--
		tmp -= int64(deg[a] + deg[b])
		deg[a]--
		deg[b]--
		tmp += int64(deg[c] + deg[d] + 2)
		deg[a]++
		deg[b]++
		ans[i+1] = tmp
	}

	return ans
}
