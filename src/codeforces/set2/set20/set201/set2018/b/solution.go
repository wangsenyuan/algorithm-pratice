package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const inf = 1 << 30

func solve(a []int) int {
	type pair struct {
		first  int
		second int
	}
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}
	slices.SortFunc(arr, func(x, y pair) int {
		if x.first != y.first {
			return x.first - y.first
		}
		return x.second - y.second
	})

	l, r := arr[0].second, arr[0].second

	for i := 1; i < n; i++ {
		j := arr[i].second
		v := arr[i].first
		l = min(l, j)
		r = max(r, j)
		if r-l+1 > v {
			return 0
		}
	}

	val1 := inf
	pref := make([]int, n)

	for i := 0; i < n; i++ {
		pref[i] = val1
		val1 = min(val1, a[i]+i)
	}

	val2 := -inf
	var res int
	for i := n - 1; i >= 0; i-- {
		if pref[i] > i && val2 < i {
			res++
		}
		val2 = max(val2, i-a[i])
	}
	return res
}
