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

	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(n, A)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
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

func solve(n int, A []int) [][]int {
	ps := make([]Pair, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], i + 1}
	}

	sort.Sort(Pairs(ps))
	for i := 0; i < n; i++ {
		A[ps[i].second-1] = i
	}
	pos := make([]int, n+1)
	for i := 0; i < n; i++ {
		pos[A[i]+1] = i + 1
	}

	X := make([]int, n+1)
	for i := 1; i <= n; i++ {
		X[i] = A[i-1] + 1
	}

	var res [][]int

	for cur := n; cur > 0; cur-- {
		if X[cur] == cur {
			continue
		}
		x := X[cur] + 1

		for x <= cur {
			i, j := pos[x], cur
			res = append(res, []int{i, j})
			X[i], X[j] = X[j], X[i]
			//pos[x], pos[cur] = pos[cur], pos[x]
			pos[X[i]], pos[X[j]] = i, j
			x++
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
