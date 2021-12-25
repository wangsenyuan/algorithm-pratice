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
		n := readNum(reader)
		P := readNNums(reader, n-1)

		res := solve(n, P)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(n int, P []int) int64 {
	degree := make([]int, n)
	cnt := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		degree[P[i]-1]++
		cnt[i+1]++
		cnt[P[i]-1] += cnt[i+1]
	}

	cnt[0]++

	children := make([][]int, n)
	for i := 0; i < n; i++ {
		children[i] = make([]int, 0, degree[i])
	}

	for i := 0; i < n-1; i++ {
		j := P[i] - 1
		children[j] = append(children[j], i+1)
	}

	var dfs func(u int) int64

	dfs = func(u int) int64 {
		var res int64

		for _, v := range children[u] {
			tmp := dfs(v)
			if tmp > res {
				res = tmp
			}
		}

		res += int64(cnt[u])

		return res
	}
	return dfs(0)
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
