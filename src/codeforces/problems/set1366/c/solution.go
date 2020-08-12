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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)

		G := make([][]int, n)

		for i := 0; i < n; i++ {
			G[i] = readNNums(reader, m)
		}

		fmt.Println(solve(n, m, G))
	}
}

func solve(n int, m int, G [][]int) int {
	N := n + m - 1

	cnt := make([]int, N)
	cnt1 := make([]int, N)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			k := i + j
			if k <= N-1-k {
				cnt[k]++
				cnt1[k] += G[i][j]
			} else {
				cnt[N-1-k]++
				cnt1[N-1-k] += G[i][j]
			}
		}
	}

	var res int

	h := N / 2

	for i := 0; i < h; i++ {
		if cnt1[i] == cnt[i] {
			continue
		}
		res += min(cnt1[i], cnt[i]-cnt1[i])
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
