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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, m)

		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 3)
		}

		res := solve(n, E)

		for _, ans := range res {
			if ans == INF {
				fmt.Println("INF")
			} else {
				fmt.Println(ans)
			}
		}
	}
}

const INF = (int64(1) << 60)
const NINF = -INF

func solve(n int, E [][]int) []int64 {
	st := make([]int64, n)
	en := make([]int64, n)

	for i := 0; i < 2*n; i++ {
		for _, e := range E {
			u, v, w := e[0]-1, e[1]-1, int64(e[2])
			st[u] = min(st[u], w+st[v])
			en[v] = min(en[v], w+en[u])
		}
	}

	ans := make([]int64, n)

	for i := 0; i < n; i++ {
		ans[i] = st[i] + en[i]
	}

	for i := 0; i < 2*n; i++ {
		for _, e := range E {
			u, v, w := e[0]-1, e[1]-1, int64(e[2])
			if st[u] > w+st[v] {
				st[u] = NINF
			}
			if en[v] > w+en[u] {
				en[v] = NINF
			}
		}
	}

	for i := 0; i < n; i++ {
		if st[i] == NINF || en[i] == NINF {
			ans[i] = INF
		}
	}

	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
