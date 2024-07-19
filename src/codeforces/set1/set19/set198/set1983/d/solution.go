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
		b := readNNums(reader, n)
		res := solve(a, b)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(a []int, b []int) bool {
	x := getSorted(a)
	y := getSorted(b)

	for i, num := range x {
		if y[i] != num {
			return false
		}
	}
	a = compact(a)
	b = compact(b)
	u := getInversions(a)
	v := getInversions(b)
	return u&1 == v&1
}

func getInversions(arr []int) int {
	n := len(arr)
	tr := make(fenwick, n+3)
	var res int
	for i := n - 1; i >= 0; i-- {
		num := arr[i]
		res += tr.pre(num)
		tr.update(num, 1)
	}
	return res
}

type fenwick []int

func (f fenwick) update(i int, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func getSorted(a []int) []int {
	n := len(a)
	res := make([]int, n)
	copy(res, a)
	sort.Ints(res)
	return res
}

func compact(a []int) []int {
	r := getSorted(a)
	res := make([]int, len(a))
	for i, num := range a {
		res[i] = sort.SearchInts(r, num)
	}
	return res
}

func getPosition(arr []int) []int {
	n := len(arr)
	pos := make([]int, n)
	for i, num := range arr {
		pos[num] = i
	}
	return pos
}
