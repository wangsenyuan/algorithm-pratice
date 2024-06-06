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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}
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

func solve(p []int) []int {
	// n := len(p)

	q1 := solve1(p, 1)
	q2 := solve1(p, 2)

	s1 := getScore(p, q1)
	s2 := getScore(p, q2)

	if s1 >= s2 {
		return q1
	}
	return q2
}

func solve1(p []int, s int) []int {
	var id []int
	var oth []int
	n := len(p)

	for i := 0; i < n; i++ {
		if i >= s && (i-s)%2 == 0 && i+1 < n {
			id = append(id, i)
		} else {
			oth = append(oth, i)
		}
	}

	sort.Slice(id, func(i, j int) bool {
		return p[id[i]] < p[id[j]]
	})
	q := make([]int, n)
	cur := n
	for _, i := range id {
		q[i] = cur
		cur--
	}

	sort.Slice(oth, func(i, j int) bool {
		return p[oth[i]] < p[oth[j]]
	})

	for _, i := range oth {
		q[i] = cur
		cur--
	}

	return q
}

func getScore(p []int, q []int) int {
	a := make([]int, len(p))
	for i := 0; i < len(p); i++ {
		a[i] = p[i] + q[i]
	}
	var res int
	for i := 1; i+1 < len(p); i++ {
		if a[i] > a[i-1] && a[i] > a[i+1] {
			res++
		}
	}
	return res
}
