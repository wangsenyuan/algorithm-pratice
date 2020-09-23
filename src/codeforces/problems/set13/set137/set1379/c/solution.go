package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		m, n := readTwoNums(reader)
		F := make([][]int, n)
		for i := 0; i < n; i++ {
			F[i] = readNNums(reader, 2)
		}
		buf.WriteString(fmt.Sprintf("%d\n", solve(m, n, F)))
		if tc > 0 {
			reader.ReadBytes('\n')
		}
	}
	fmt.Print(buf.String())
}

func solve(m, n int, F [][]int) int64 {
	//let iterate i from 0 ~ n, make sure it (only) has more than 1 flowers
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{F[i][0], i}
	}

	sort.Sort(Pairs(ps))

	X := make([]int64, n+1)
	for i := 0; i < n; i++ {
		X[i+1] = X[i] + int64(ps[i].first)
	}
	var ans int64

	if m <= n {
		ans = X[m]
	}

	for i := 0; i < n; i++ {
		cur := F[ps[i].second][1]

		j := sort.Search(n, func(j int) bool {
			return ps[j].first <= cur
		})
		var left = m - j
		if left > 0 {
			tmp := X[j]
			if j <= i {
				tmp += int64(F[ps[i].second][0]) + int64(left-1)*int64(F[ps[i].second][1])
			} else {
				tmp += int64(left) * int64(F[ps[i].second][1])
			}
			if tmp > ans {
				ans = tmp
			}
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

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first > ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
