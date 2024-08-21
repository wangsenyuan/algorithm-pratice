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
		nr, ng, nb := readThreeNums(reader)
		r := readNNums(reader, nr)
		g := readNNums(reader, ng)
		b := readNNums(reader, nb)
		res := solve(r, g, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const inf = 1 << 62

func solve(r []int, g []int, b []int) int {
	sort.Ints(r)
	sort.Ints(g)
	sort.Ints(b)

	res := process(r, g, b)
	res = min(res, process(g, r, b))
	res = min(res, process(b, r, g))

	return res
}

func process(r, g, b []int) int {
	best := inf

	for i := 0; i < len(r); i++ {
		x := r[i]
		u, v := find(g, x)
		c, d := find(b, x)

		for j := u; j <= v; j++ {
			for k := c; k <= d; k++ {
				tmp := (x-g[j])*(x-g[j]) + (x-b[k])*(x-b[k]) + (g[j]-b[k])*(g[j]-b[k])
				best = min(best, tmp)
			}
		}
	}
	return best
}

func find(arr []int, x int) (int, int) {
	i := sort.SearchInts(arr, x)
	l := max(0, i-2)
	r := min(len(arr)-1, i+2)
	return l, r
}
