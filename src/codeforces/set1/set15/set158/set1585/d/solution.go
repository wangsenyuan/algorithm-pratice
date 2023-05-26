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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(A []int) bool {
	// i, j, k 三个都不能在位置上，
	// 只要有多于3个的下标的位置不一致，就是可行的
	if !all_distinct(A) {
		return true
	}
	// all distinct
	m := countInversions(A)

	return m%2 == 0
}

func all_distinct(A []int) bool {
	n := len(A)
	arr := make([]int, n)
	copy(arr, A)
	sort.Ints(arr)
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			return false
		}
	}
	return true
}

func countInversions(A []int) int64 {
	n := len(A)

	arr := make([]int, 2*n)

	set := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = arr[p] + arr[p^1]
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res += arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	type Pair struct {
		first  int
		second int
	}

	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	var res int64

	for i := 0; i < n; {
		j := i
		// 比 ps[j].first 小的，但是下标在它后面的数有多少个
		for i < n && ps[i].first == ps[j].first {
			res += int64(get(ps[i].second, n))
			i++
		}
		for j < i {
			set(ps[j].second, 1)
			j++
		}
	}

	return res
}
