package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, b, c := readThreeNums(reader)
	solve(a, b, c, reader)
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

func solve(a, b, c int, reader *bufio.Reader) {
	// choose first to play
	fmt.Println("First")

	ps := make([]*Pair, 3)
	ps[0] = &Pair{int64(a), 1}
	ps[1] = &Pair{int64(b), 2}
	ps[2] = &Pair{int64(c), 3}

	sort.Sort(Pairs(ps))

	arr := make([]int, 3)
	for i := 0; i < 3; i++ {
		arr[ps[i].second-1] = i
	}

	var play func()

	play = func() {
		p, q, r := ps[0].first, ps[1].first, ps[2].first

		y := 2*r - p - q

		fmt.Printf("%d\n", y)

		x := readNum(reader)

		i := arr[x-1]
		if i == 0 {
			fmt.Printf("%d\n", r-q)
		} else if i == 1 {
			fmt.Printf("%d\n", r-p)
		} else {
			ps[2].first += y
			play()
		}
	}
	play()
}

type Pair struct {
	first  int64
	second int
}

type Pairs []*Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
