package main

import (
	"bufio"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)
	for tc > 0 {
		tc--
		n, Q := readTwoNums(scanner)
		P := readNNums(scanner, n)
		X := make([]int, Q)
		for i := 0; i < Q; i++ {
			X[i] = readNum(scanner)
		}
		res := solve(n, Q, P, X)

		for i := 0; i < Q; i++ {
			fmt.Println(res[i])
		}

	}
}

func solve(N, Q int, P []int, X []int) []int {
	// var INF int64 = math.MaxInt64
	// x 增加 1, 就可以产生很大的增幅
	sort.Ints(P)

	qs := make([]Pair, Q)

	for i := 0; i < Q; i++ {
		qs[i] = Pair{i, X[i]}
	}

	sort.Sort(Pairs(qs))

	const INF = int64(1 << 30)

	var prev, rem int64
	var ans int

	Y := make([]int, Q)

	for i := 0; i < Q; i++ {
		if int64(qs[i].second) > prev {
			if ans > 30 {
				rem = INF
			} else {
				rem += (int64(qs[i].second) - prev) * int64(1<<uint(ans))
			}
			prev = int64(qs[i].second)

			for ans < N && rem > int64(P[ans]) {
				rem = min(INF, 2*(rem-int64(P[ans])))
				ans++
			}
		}
		Y[qs[i].first] = ans
	}

	return Y
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
