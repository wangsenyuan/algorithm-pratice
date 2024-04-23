package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	x := readNum(reader)
	res := solve(a, b, x)

	fmt.Println(res)
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

const oo = 1 << 60

func solve(a []int, b []int, x int) int {
	n := len(a)
	m := len(b)

	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + a[i]
	}

	// 相同长度时，选择最小的sum
	arr := make([]int, n+1)
	for l := 1; l <= n; l++ {
		val := oo
		for i := 0; i+l <= n; i++ {
			val = min(val, pref[i+l]-pref[i])
		}
		arr[l] = val
	}

	// arr[l] 是一个递增序列 (很容易证明, 越长的序列越大）
	pb := make([]int, m+1)
	for i := 0; i < m; i++ {
		pb[i+1] = pb[i] + b[i]
	}

	var best int
	for l, w := 1, n; l <= m; l++ {
		val := oo
		for i := 0; i+l <= m; i++ {
			val = min(val, pb[i+l]-pb[i])
		}
		for w > 0 && val*arr[w] > x {
			w--
		}
		if arr[w]*val <= x {
			best = max(best, l*w)
		}
	}

	return best
}

type Pair struct {
	first  int
	second int
}
