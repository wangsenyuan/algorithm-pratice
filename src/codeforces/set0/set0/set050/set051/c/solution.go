package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)

	X := readNNums(reader, n)

	res, pos := solve(n, X)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%.6f\n", res))

	for i := 0; i < len(pos); i++ {
		buf.WriteString(fmt.Sprintf("%.6f ", pos[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, X []int) (float64, []float64) {
	sort.Ints(X)

	check := func(d int) bool {
		var cover int
		var cnt int
		for i := 0; i < n; i++ {
			if X[i] > cover {
				cover = X[i] + d
				cnt++
			}
		}
		return cnt <= 3
	}

	var l, r int = 0, 2 * 1e9

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	ans := r

	var cover int
	var cnt int

	pos := make([]float64, 3)

	for i := 0; i < n; i++ {
		if X[i] > cover {
			cover = X[i] + ans
			pos[cnt] = float64(X[i]) + float64(ans)/2
			cnt++
		}
	}

	sort.Float64s(pos)

	return float64(ans) / 2, pos
}

func search(l int, r int, f func(int) bool) int {
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
