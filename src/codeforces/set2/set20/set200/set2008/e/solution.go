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
		readNum(reader)
		s := readString(reader)
		res := solve(s)
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

func solve(s string) int {
	n := len(s)
	if n%2 == 0 {
		return solveEven(s)
	}
	return solveOdd(s)
}

func solveEven(s string) int {
	// just find the most frequent in odd/even position
	freq := make([][]int, 2)
	for i := 0; i < 2; i++ {
		freq[i] = make([]int, 26)
	}

	for i, x := range []byte(s) {
		freq[i%2][int(x-'a')]++
	}
	best := make([]int, 2)

	n := len(s)

	for i := 0; i < 26; i++ {
		best[0] = max(best[0], freq[0][i])
		best[1] = max(best[1], freq[1][i])
	}

	return n/2 - best[0] + n/2 - best[1]
}

func solveOdd(s string) int {
	n := len(s)
	pref := make([][]int, n)
	for i := 0; i < n; i++ {
		pref[i] = make([]int, 26)
	}

	for i, x := range []byte(s) {
		if i > 1 {
			copy(pref[i], pref[i-2])
		}
		pref[i][int(x-'a')]++
	}
	suf := make([][]int, 2)
	for i := 0; i < 2; i++ {
		suf[i] = make([]int, 26)
	}

	get := func(i int, suf [][]int) int {
		// 如果删除i
		best := []int{0, 0}
		for x := 0; x < 26; x++ {
			if i > 0 {
				best[(i-1)%2] = max(best[(i-1)%2], pref[i-1][x]+suf[(i+2)%2][x])
			} else {
				best[(i%2)^1] = max(best[(i%2)^1], suf[(i+2)%2][x])
			}
			if i > 1 {
				best[i%2] = max(best[i%2], pref[i-2][x]+suf[(i+1)%2][x])
			} else {
				best[i%2] = max(best[i%2], suf[(i+1)%2][x])
			}
		}
		return n/2*2 - (best[0] + best[1])
	}

	res := n

	for i := n - 1; i >= 0; i-- {
		res = min(res, get(i, suf))
		x := int(s[i] - 'a')
		suf[i%2][x]++
	}

	return res + 1
}
