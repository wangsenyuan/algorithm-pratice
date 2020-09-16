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

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		cnt, res := solve(n, A)
		fmt.Println(cnt)
		if cnt < 0 {
			continue
		}
		var buf bytes.Buffer
		for i := 0; i < cnt; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		fmt.Println(buf.String())
	}
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

func cycle(a, b, c *int) {
	*a, *b, *c = *c, *a, *b
}

func solve(n int, A []int) (int, []int) {
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], i}
	}

	sort.Sort(Pairs(ps))

	for i := 0; i < n; i++ {
		A[ps[i].second] = i
	}

	var cnt int

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if A[j] < A[i] {
				cnt++
			}
		}
	}

	if cnt&1 == 1 {
		var ok = false
		for i := 0; i < n-1 && !ok; i++ {
			if ps[i].first == ps[i+1].first {
				A[ps[i].second], A[ps[i+1].second] = A[ps[i+1].second], A[ps[i].second]
				ok = true
			}
		}
		if !ok {
			return -1, nil
		}
	}
	res := make([]int, 0, n)
	for i := 0; i < n-2; i++ {
		var pos = i
		for j := i + 1; j < n; j++ {
			if A[j] < A[pos] {
				pos = j
			}
		}

		for pos > i+1 {
			cycle(&A[pos-2], &A[pos-1], &A[pos])
			res = append(res, pos-2+1)
			pos -= 2
		}

		if pos != i {
			// pos = i + 1
			// 3, 1, 2, the only way
			j := pos + 1
			cycle(&A[j-2], &A[j-1], &A[j])
			res = append(res, j-2+1)
			cycle(&A[j-2], &A[j-1], &A[j])
			res = append(res, j-2+1)
			pos = i
		}
	}

	if !sort.IntsAreSorted(A) {
		// 2, 3, 1 the only way
		cycle(&A[n-3], &A[n-2], &A[n-1])
		res = append(res, n-3+1)
	}

	if !sort.IntsAreSorted(A) {
		return -1, nil
	}

	return len(res), res
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	if this.first < that.first {
		return true
	}
	if this.first == that.first && (this.second < that.second) {
		return true
	}
	return false
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
