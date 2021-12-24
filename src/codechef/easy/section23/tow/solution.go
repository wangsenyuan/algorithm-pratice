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

		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, m)
		res := solve(A, B)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
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

func solve(A []int, B []int) []int {
	pa := make([]Pair, len(A))
	for i := 0; i < len(A); i++ {
		pa[i] = Pair{A[i], i}
	}

	sort.Sort(Pairs(pa))

	pb := make([]Pair, len(B))
	for i := 0; i < len(B); i++ {
		pb[i] = Pair{B[i], i}
	}

	sort.Sort(Pairs(pb))

	res := make([]int, 0, len(B))

	p1 := len(B) - 1
	res = append(res, pb[p1].first)
	var p2 int
	for p2 < len(A) && p1 >= 0 {
		if pb[p1].first < A[p2] {
			return nil
		}
		if pb[p1].first > A[p2] {
			p2++
			continue
		}
		p1--
		if p1 >= 0 {
			res = append(res, pb[p1].first)
		}
		p2++
	}

	if p1 < 0 {
		return nil
	}

	ans := make([]int, 0, len(B))

	for i := 0; i < p1; i++ {
		ans = append(ans, pb[i].first)
	}

	ans = append(ans, res...)

	return ans
}

func last(arr []Pair) Pair {
	return arr[len(arr)-1]
}

func countELement(arr []int, num int) int {
	var res int
	for _, cur := range arr {
		if cur == num {
			res++
		}
	}
	return res
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].Less(ps[j])
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
