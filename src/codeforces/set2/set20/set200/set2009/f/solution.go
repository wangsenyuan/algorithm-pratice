package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)

		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(a []int, queries [][]int) []int {
	n := len(a)

	pref := make([]int, 2*n+1)
	for i := 0; i < 2*n; i++ {
		pref[i+1] = pref[i] + a[i%n]
	}
	sum := pref[n]

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		// 先算出中间有多少个完整的
		if r/n != l/n {
			cnt := r/n - l/n - 1
			ans[i] = max(0, cnt) * sum
			// l所在的那段是没有计算的
			cl, pl := l/n, l%n
			// 首位是 a[cl],
			// 0, 1, 2, 3, ...., n-1, n, n+1, .... 2 * n
			ans[i] += pref[cl+n] - pref[cl+pl]
			rl, pr := r/n, r%n
			ans[i] += pref[rl+pr+1] - pref[rl]
		} else {
			cl, pl := l/n, l%n
			pr := r % n
			ans[i] = pref[cl+pr+1] - pref[cl+pl]
		}
	}

	return ans
}
