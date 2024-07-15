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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, m)
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(a []int, b []int) int {
	i := sort.SearchInts(a, 0)
	j := sort.SearchInts(b, 0)
	// a[i] >= 0, b[i] >= 0
	res := play(a[i:], b[j:])
	c := reverse(negative(a[:i]))
	d := reverse(negative(b[:j]))
	res += play(c, d)

	return res
}

func negative(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	for i := 0; i < len(res); i++ {
		res[i] *= -1
	}
	return res
}

func reverse(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func play(a []int, b []int) int {
	// a[0] >= 0, b[0] >= 0
	// 还必须知道那些没有移动的box的是否在正确的位置

	m := len(a)
	n := len(b)

	suf := make([]int, n+1)
	for i, j := n-1, m-1; i >= 0; i-- {
		for j >= 0 && a[j] > b[i] {
			j--
		}
		suf[i] = suf[i+1]
		if j >= 0 && a[j] == b[i] {
			suf[i]++
			j--
		}
	}

	var best int
	for r, i := 0, 0; r < n; r++ {
		for i < m && a[i] <= b[r] {
			// 先遇到一个box， push it
			i++
		}
		// b[r] < a[i]
		// b[r] - cnt 个中间有多少个specical position
		j := sort.Search(r+1, func(j int) bool {
			return b[j] >= b[r]-(i-1)
		})
		// b[j] < b[r] - i
		// b[j] >= b[r] - i
		best = max(best, min(r-j+1, i)+suf[r+1])
	}

	return best
}
