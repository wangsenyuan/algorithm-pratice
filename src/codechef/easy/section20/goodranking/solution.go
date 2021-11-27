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
	var buf bytes.Buffer
	tc := readNum(reader)
	// tc := 1
	for tc > 0 {
		tc--
		n := readNum(reader)
		S := make([]string, n)
		for i := 0; i < n; i++ {
			S[i], _ = reader.ReadString('\n')
		}
		res := solve(n, S)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

func solve(n int, S []string) []int {
	// pi < pj if S[P[i]][P[j]] = 1 and j - i < (n + 1) / 2
	win := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if S[i][j] == '1' {
				win[i]++
			}
		}
	}
	P := make([]Pair, n)
	for i := 0; i < n; i++ {
		P[i] = Pair{win[i], i}
	}
	sort.Sort(Pairs(P))

	res := make([]int, n)

	fillOrder := func(start, end int) {
		for i := 0; i < n; i++ {
			win[i] = 0
		}
		pp := make([]Pair, end-start)

		for i := start; i < end; i++ {
			x := P[i].second
			for j := start; j < end; j++ {
				y := P[j].second
				if S[x][y] == '1' {
					win[x]++
				}
			}
			pp[i-start] = Pair{win[x], x}
		}
		sort.Sort(Pairs(pp))
		for i := start; i < end; i++ {
			res[i] = pp[i-start].second
		}
	}

	fillOrder(0, n/2)
	if n&1 == 1 {
		res[n/2] = P[n/2].second
	}
	fillOrder(n/2+(n&1), n)

	for i := 0; i < n; i++ {
		x := res[i]
		for j := i + 1; j < min(i+(n+1)/2+1, n); j++ {
			y := res[j]
			if S[x][y] == '0' {
				return nil
			}
		}
	}

	for i := 0; i < n; i++ {
		res[i]++
	}

	return res
}

func min(a, b int) int {
	if a <= b {
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
	return ps[i].first > ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
