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
		B := readNNums(reader, n)
		ok, res := solve(A, B)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
			}
		}
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A, B []int) (bool, [][]int) {
	n := len(A)
	pa := make([]Pair, n)

	for i := 0; i < n; i++ {
		pa[i] = Pair{A[i], B[i]}
		for j := i + 1; j < n; j++ {
			if A[i] < A[j] && B[i] > B[j] {
				return false, nil
			}
			if A[i] > A[j] && B[i] < B[j] {
				return false, nil
			}
		}
	}

	sa := NewSorter(pa)
	sort.Sort(sa)

	res := sa.res

	return true, res
}

type Pair struct {
	first  int
	second int
}

type Sorter struct {
	arr []Pair
	res [][]int
}

func NewSorter(arr []Pair) *Sorter {
	s := new(Sorter)
	s.arr = arr

	return s
}

func (s *Sorter) Len() int {
	return len(s.arr)
}

func (s *Sorter) Less(i, j int) bool {
	a := s.arr[i]
	b := s.arr[j]
	return a.first <= b.first && a.second <= b.second
}

func (s *Sorter) Swap(i, j int) {
	s.res = append(s.res, []int{i + 1, j + 1})
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}
