package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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
	firstLine := readNNums(reader, 4)
	n, h, m, k := firstLine[0], firstLine[1], firstLine[2], firstLine[3]
	trams := make([][]int, n)
	for i := 0; i < n; i++ {
		trams[i] = readNNums(reader, 2)
	}
	cnt, t, res := solve(n, h, m, k, trams)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %d\n", cnt, t))

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
}

func solve(n int, h int, m int, k int, trams [][]int) (int, int, []int) {
	pp := make([]Pair, 0, 2*n)

	test := make([]int64, 0, 4*n)

	H := int64(h)
	M := int64(m)
	X := H * M
	K := int64(k)
	MOD := M / 2
	for i := 0; i < n; i++ {
		t := (X + int64(trams[i][1])) % MOD
		pp = append(pp, Pair{t, i})
		pp = append(pp, Pair{MOD + t, i})

		test = append(test, (MOD+t-1)%MOD)
		test = append(test, (MOD+t-K)%MOD)
		test = append(test, (t+1)%MOD)
		test = append(test, (t+K)%MOD)
	}

	sort.Sort(Pairs(pp))

	f := func(x, y int64) int {
		i := sort.Search(len(pp), func(i int) bool {
			return pp[i].first > y
		})

		j := sort.Search(len(pp), func(j int) bool {
			return pp[j].first >= x
		})

		return i - j
	}
	var res = math.MaxInt32
	var pos int
	for _, t := range test {
		tmp := f(MOD+t-K+1, MOD+t-1)
		if tmp < res {
			res = tmp
			pos = int(t)
		}
	}

	cancel := make([]int, 0, n)

	for i := 0; i < len(pp); i++ {
		if MOD+int64(pos)-K+1 <= pp[i].first && pp[i].first <= MOD+int64(pos)-1 {
			cancel = append(cancel, pp[i].second+1)
		}
	}

	return res, pos, cancel
}

type Pair struct {
	first  int64
	second int
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
