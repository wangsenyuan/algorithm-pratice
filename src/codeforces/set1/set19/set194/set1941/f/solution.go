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
		n, m, k := readThreeNums(reader)
		a := readNNums(reader, n)
		d := readNNums(reader, m)
		f := readNNums(reader, k)
		res := solve(a, d, f)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(a []int, d []int, f []int) int {

	first, second := Pair{0, -1}, Pair{0, -1}
	for i := 0; i+1 < len(a); i++ {
		diff := a[i+1] - a[i]
		if diff >= first.first {
			second = first
			first = Pair{diff, i}
		} else if diff >= second.first {
			second = Pair{diff, i}
		}
	}
	// try to reduce first, and check the new result
	if first.first == second.first {
		// no way to reduce the result
		return first.first
	}

	sort.Ints(d)
	sort.Ints(f)

	best := first.first
	// d[i] + f[j] > a[first.second] and < a[first.second + 1]
	// 且越在中间越好

	get := func(cur int) int {
		return max(cur-a[first.second], a[first.second+1]-cur)
	}

	expect := (a[first.second] + a[first.second+1]) / 2

	for _, num := range d {
		j := search(len(f), func(j int) bool {
			return num+f[j] > expect
		})
		if j < len(f) {
			cur := num + f[j]
			best = min(best, max(second.first, get(cur)))
		}
		if j > 0 {
			j--
			cur := num + f[j]
			best = min(best, max(second.first, get(cur)))
		}
	}

	return best
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Pair struct {
	first  int
	second int
}
