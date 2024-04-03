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
		n, k := readTwoNums(reader)
		buf.WriteString(solve(n, k))
	}

	fmt.Println(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

var f0 = "What are you doing at the end of the world? Are you busy? Will you save us?"

var pref = "What are you doing while sending \""
var mid = "\"? Are you busy? Will you send \""
var suf = "\"?"

const N = 54

var F [N]int

func init() {
	F[0] = len(f0)

	// f[i] 需要知道它的长度
	// f[i] = len(pref) + f[i-1] + len(mid) + f[i-1] + len(suf)
	// 这个增长太快了
	// f[i] = 2 * f[i-1] + constant
	for i := 1; i < N; i++ {
		F[i] = 2*F[i-1] + len(pref) + len(mid) + len(suf)
	}
}

func solve(n int, k int) string {
	return work(n, k-1)
}

func work(n int, k int) string {
	if n == 0 {
		if k >= F[0] {
			return "."
		}
		return f0[k : k+1]
	}

	if k < len(pref) {
		return pref[k : k+1]
	}

	if n >= N {
		return work(n-1, k-len(pref))
	}

	k -= len(pref)
	if k < F[n-1] {
		return work(n-1, k)
	}
	k -= F[n-1]

	if k < len(mid) {
		return mid[k : k+1]
	}
	k -= len(mid)

	if k < F[n-1] {
		return work(n-1, k)
	}
	k -= F[n-1]
	if k < len(suf) {
		return suf[k : k+1]
	}
	return "."
}
