package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, k)

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

const inf = 1 << 60

func solve(a []int, k int) int {
	sort.Ints(a)
	n := len(a)
	// 如果x是最大值
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = a[i] + suf[i+1]
	}

	res := suf[0] - k*a[0]

	var pref int

	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		if i-j >= k {
			return 0
		}
		// 前面有j个 a[j], 后面有 (n - i)个 a[i]
		c := i - j
		// c < k, tmp = 变成 a[j] - 1 和 a[j] + 1 的cost
		tmp1 := j*(a[j]-1) - pref
		tmp2 := suf[i] - (n-i)*(a[j]+1)
		if j+c >= k {
			// 把j中k-c个变成a[j]
			res = min(res, tmp1+k-c)
		}
		if n-i+c >= k {
			res = min(res, tmp2+k-c)
		}
		// j + n - i + c = n >= k
		res = min(res, tmp1+tmp2+k-c)

		pref += (i - j) * a[j]
	}

	return res
}
