package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	buses := make([][]int, k)
	for i := 0; i < k; i++ {
		buses[i] = readNNums(reader, 2)
	}
	res := solve(n, buses)

	fmt.Println(res)
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

const M = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= M {
		a -= M
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, M-b)
}

func solve(n int, buses [][]int) int {
	type Pair struct {
		first  int
		second int
	}

	m := len(buses)

	ends := make([]int, m)

	ps := make([]Pair, m)

	for i := 0; i < m; i++ {
		ps[i] = Pair{buses[i][1], buses[i][0]}
		ends[i] = buses[i][1]
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	sort.Ints(ends)

	dp := make([]int, m+1)
	sum := make([]int, m+1)

	for i := 0; i < m; i++ {
		p := ps[i]
		if p.second == 0 {
			dp[i] = 1
		}
		s := search(m, func(j int) bool {
			return ends[j] >= p.second
		})
		t := search(m, func(j int) bool {
			return ends[j] >= p.first
		})
		dp[i] = modAdd(dp[i], modSub(sum[t], sum[s]))
		sum[i+1] = modAdd(sum[i], dp[i])
	}

	var res int

	for i := 0; i < m; i++ {
		if ps[i].first == n {
			res = modAdd(res, dp[i])
		}
	}
	return res
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
