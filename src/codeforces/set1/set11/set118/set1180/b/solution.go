package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	//tc := readNum(scanner)
	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Println(buf.String())
}

func solve(a []int) []int {
	//n := len(a)
	// a[i] >= 0 => -a[i] - 1 绝对值更大，最好是变
	// a[i] < 0 => -a[i] - 1 绝对值更小，最好是不变 (-1)
	var first []Pair
	var second []Pair
	for i, num := range a {
		if num >= 0 {
			first = append(first, Pair{num, i})
		} else {
			second = append(second, Pair{num, i})
		}
	}
	sort.Slice(first, func(i, j int) bool {
		return first[i].first > first[j].first
	})

	n := len(a)

	// 要保证len(second)的size是偶数
	for _, x := range first {
		second = append(second, x.change())
	}

	sort.Slice(second, func(i, j int) bool {
		return second[i].first > second[j].first
	})

	if n%2 == 1 {
		second[n-1] = second[n-1].change()
	}

	res := make([]int, n)

	for _, x := range second {
		res[x.second] = x.first
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func (p Pair) change() Pair {
	return Pair{-p.first - 1, p.second}
}
