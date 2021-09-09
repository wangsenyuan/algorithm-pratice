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
		n, q := readTwoNums(reader)

		A := make([]uint64, n)
		B := make([]uint64, n)
		s, _ := reader.ReadBytes('\n')
		var pos int
		for i := 0; i < n; i++ {
			pos = readUint64(s, pos, &A[i]) + 1
		}

		pos = 0
		s, _ = reader.ReadBytes('\n')
		for i := 0; i < n; i++ {
			pos = readUint64(s, pos, &B[i]) + 1
		}

		Q := readNNums(reader, q)

		res := solve(n, A, B, Q)

		for i := 0; i < q; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func solve(n int, A []uint64, B []uint64, Q []int) []uint64 {
	sort.Sort(Uint64s(A))
	sort.Sort(Uint64s(B))

	findPos := func(num uint64) int {
		r, c := 0, n-1
		var cnt int
		for r < n && c >= 0 {
			tmp := A[r] + B[c]
			if tmp > num {
				c--
			} else {
				// tmp <= num
				cnt += c + 1
				r++
			}
		}
		return cnt
	}

	ans := make([]uint64, len(Q))

	for i, cur := range Q {
		var left, right uint64 = 0, A[n-1] + B[n-1]
		for left < right {
			mid := (left + right) / 2
			if findPos(mid) >= cur {
				right = mid
			} else {
				left = mid + 1
			}
		}
		ans[i] = right
	}

	return ans
}

type Uint64s []uint64

func (this Uint64s) Len() int {
	return len(this)
}

func (this Uint64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Uint64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
