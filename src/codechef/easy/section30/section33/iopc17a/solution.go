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
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
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

const INF = 1 << 30

func solve(s string) int {
	// if all s have the same character, well good
	// otherwise, if its length odd, -1
	//            else split it
	n := len(s)
	// 需要一个快速知道[l...r]中的字母是否是同一个的方式
	cnt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cnt[i] = make([]int, 26)
		if i > 0 {
			copy(cnt[i], cnt[i-1])
			cnt[i][int(s[i-1]-'a')]++
		}
	}

	var dfs func(s int, e int) int

	dfs = func(s int, e int) int {
		for x := 0; x < 26; x++ {
			if cnt[e][x]-cnt[s][x] == e-s {
				return 0
			}
		}
		if (e-s)&1 == 1 {
			// odd length
			return -1
		}
		mid := (e + s) / 2
		a := dfs(s, mid)
		b := dfs(mid, e)
		if a < 0 && b < 0 {
			return -1
		}
		if a < 0 {
			return b + 1
		}
		if b < 0 {
			return a + 1
		}
		return min(a, b) + 1
	}

	return dfs(0, n)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
