package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		X := make([][]int, n)

		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var m int
			pos := readInt(s, 0, &m) + 1
			X[i] = make([]int, m)
			for j := 0; j < m; j++ {
				pos = readInt(s, pos, &X[i][j]) + 1
			}
		}
		res := solve(n, X)
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

func solve(n int, X [][]int) int64 {
	coord := make(map[int]int)
	var m int
	for _, x := range X {
		for i := 0; i < len(x); i++ {
			coord[abs(x[i])]++
		}
		if m < len(x) {
			m = len(x)
		}
	}

	ps := make([]Pair, m)
	var res int64
	for i := 0; i < n; i++ {
		var pos, neg int
		for j := 0; j < len(X[i]); j++ {
			ps[j] = Pair{abs(X[i][j]), sign(X[i][j])}
			if ps[j].second == 1 {
				pos++
			} else {
				neg++
			}
		}

		sort.Sort(Pairs(ps[:len(X[i])]))

		for j := 0; j < len(X[i]); j++ {
			p := ps[j]
			if p.second > 0 {
				pos--
			} else {
				neg--
			}

			if coord[p.first] > 1 {
				if p.second < 0 {
					res += int64(neg)
				} else {
					res += int64(pos)
				}
			} else {
				if p.second < 0 {
					res += int64(pos)
				} else {
					res += int64(neg)
				}
			}
		}
	}

	for _, v := range coord {
		if v > 1 {
			res++
		}
	}

	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	return 1
}
