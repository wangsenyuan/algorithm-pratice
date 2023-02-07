package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func solve(P []int, s string) []int {
	n := len(P)
	// sum(abs(P[i] - Q[i])) min
	// 需要把大值分配给喜欢的部分
	var A []Pair
	var B []Pair

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			A = append(A, Pair{P[i], i})
		} else {
			B = append(B, Pair{P[i], i})
		}
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i].Less(A[j])
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i].Less(B[j])
	})

	Q := make([]int, n)
	cur := n

	for i := 0; i < len(B); i++ {
		Q[B[i].second] = cur
		cur--
	}

	for i := 0; i < len(A); i++ {
		Q[A[i].second] = cur
		cur--
	}

	return Q
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first > that.first
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, n)
		s, _ := reader.ReadString('\n')
		res := solve(P, s)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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
