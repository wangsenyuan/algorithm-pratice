package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := make([][]int, n)
		for i := 0; i < n; i++ {
			P[i] = readNNums(reader, 2)
		}
		res := solve(n, P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, P [][]int) int {
	C := make([][]Pair, n+1)
	S := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]Pair, 0, n)
		S[i] = make([]int, 0, n)
	}

	sort.Sort(Points(P))

	C[0] = append(C[0], Pair{math.MaxFloat64, 1})
	S[0] = append(S[0], 1)

	ans := 1

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			slop := float64(P[i][1]-P[j][1]) / float64(P[i][0]-P[j][0])
			k := search(len(C[j]), func(k int) bool {
				return C[j][k].first >= slop
			})
			ans = max(ans, S[j][k]+1)
			C[i] = append(C[i], Pair{slop, S[j][k] + 1})
		}

		C[i] = append(C[i], Pair{math.MaxFloat64, 1})
		sort.Sort(Pairs(C[i]))

		for _, c := range C[i] {
			S[i] = append(S[i], c.second)
		}

		for j := len(S[i]) - 2; j >= 0; j-- {
			S[i][j] = max(S[i][j], S[i][j+1])
		}

	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

type Pair struct {
	first  float64
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Points [][]int

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Less(i, j int) bool {
	return ps[i][0] < ps[j][0]
}

func (ps Points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
