package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := make([]int, n)
	C := make([]int, n)
	for i := 0; i < n; i++ {
		A[i], C[i] = readTwoNums(reader)
	}
	res := solve(n, A, C)
	fmt.Println(res)
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, A []int, C []int) int64 {
	P := make([]Pair, n)
	for i := 0; i < n; i++ {
		P[i] = Pair{A[i], C[i]}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	var ans int64
	var ac int64
	var sum int64

	for i := 0; i < n; i++ {
		if i != 0 && ac < int64(P[i].first) {
			ans += int64(P[i].first) - ac
		}
		ac = max(ac, int64(P[i].first+P[i].second))
		sum += int64(P[i].second)
	}

	return ans + sum
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}
