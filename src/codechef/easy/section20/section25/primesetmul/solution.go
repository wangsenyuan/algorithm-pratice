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

		s, _ := reader.ReadBytes('\n')

		var n int
		pos := readInt(s, 0, &n) + 1
		var M uint64
		readUint64(s, pos, &M)

		S := readNNums(reader, n)

		res := solve(S, int64(M))

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(S []int, M int64) int64 {

	var generate func(L []int, pos int, cur int64)

	var res []int64

	generate = func(L []int, pos int, cur int64) {
		if cur > M {
			return
		}
		if pos == len(L) {
			res = append(res, cur)
			return
		}
		generate(L, pos+1, cur)
		for cur*int64(L[pos]) <= M {
			cur *= int64(L[pos])
			generate(L, pos+1, cur)
		}
	}

	sort.Ints(S)
	n := len(S)
	A := make([]int, 0, n-n/2)
	B := make([]int, 0, n/2)
	for i := 0; i < n; i += 2 {
		A = append(A, S[i])
		if i+1 < n {
			B = append(B, S[i+1])
		}
	}

	generate(A, 0, 1)

	L := res

	res = make([]int64, 0, len(L))

	generate(B, 0, 1)

	R := res

	sort.Sort(Int64s(L))

	var ans int64

	for _, cur := range R {
		j := search(len(L), func(i int) bool {
			return cur > M/L[i]
		})
		ans += int64(j)
	}

	return ans
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Int64s []int64

func (this Int64s) Len() int {
	return len(this)
}

func (this Int64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
