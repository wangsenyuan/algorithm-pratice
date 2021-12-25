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
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		N, D := readTwoNums(reader)
		P := readNNums(reader, N)
		S := readNNums(reader, N)
		res := solve(N, D, P, S)
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

func solve(N, D int, P []int, S []int) int64 {
	// if pick some candy with p, and s,
	// then the best way is to choose some candy thant price <= D - p, with max s
	// so sort by price in order, and keep the best s so for,
	candies := make([]Candy, N)

	for i := 0; i < N; i++ {
		candies[i] = Candy{P[i], S[i]}
	}

	sort.Sort(Candies(candies))

	if candies[0].price > D {
		return 0
	}

	A := make([]int64, N)

	A[0] = int64(candies[0].sweet)

	var best int64 = int64(candies[0].sweet)

	for i := 1; i < N; i++ {
		if candies[i].price <= D {
			best = max(best, int64(candies[i].sweet))
		}
		j := sort.Search(i, func(j int) bool {
			return candies[j].price+candies[i].price > D
		})
		j--
		if j >= 0 {
			tmp := int64(candies[i].sweet) + A[j]
			best = max(best, tmp)
		}
		A[i] = max(A[i-1], int64(candies[i].sweet))
	}
	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Candy struct {
	price int
	sweet int
}

type Candies []Candy

func (this Candies) Len() int {
	return len(this)
}

func (this Candies) Less(i, j int) bool {
	return this[i].price < this[j].price
}

func (this Candies) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
