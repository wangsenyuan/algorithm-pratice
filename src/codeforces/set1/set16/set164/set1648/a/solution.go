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

	tc := 1

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		table := make([][]int, n)
		for i := 0; i < n; i++ {
			table[i] = readNNums(reader, m)
		}
		res := solve(table)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(table [][]int) int64 {
	//  对于color x的cells, (x0, y0), (x1, y1), (x2, y2)...
	// x 和 y是独立的

	pos := make(map[int][]Pair)

	add := func(c int, i int, j int) {
		if _, ok := pos[c]; !ok {
			pos[c] = make([]Pair, 0, 1)
		}
		pos[c] = append(pos[c], Pair{i, j})
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			c := table[i][j]
			add(c, i, j)
		}
	}

	var res int64

	for _, vs := range pos {
		res += calc(vs)
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func calc(arr []Pair) int64 {
	n := len(arr)
	xs := make([]int, n)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = arr[i].first
		ys[i] = arr[i].second
	}
	return calcSum(xs) + calcSum(ys)
}

func calcSum(arr []int) int64 {
	sort.Ints(arr)
	n := len(arr)

	var res int64
	var pref int64

	for i := 0; i < n; i++ {
		// x[i] distance to other points
		res += int64(i)*int64(arr[i]) - pref
		pref += int64(arr[i])

	}
	// 所有的都被计算了两次
	return res
}
