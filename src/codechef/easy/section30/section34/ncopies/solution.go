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
		_, m := readTwoNums(reader)
		S := readString(reader)
		res := solve(S, m)

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

func solve(A string, M int) int64 {
	n := len(A)
	// 101 3
	// 101101101
	// for i = 5, pref 10110 suf 1101, f(pref) = 3, f(suf) = 3
	// for i = 4, pref 1011, suf 01101
	// if all A is 0, then it is n * m
	// else, if m is even
	// suf_zero(A) + pref_zero(A) + 1 (last 1 in A)
	// 100100
	// else if m is odd
	// 100100100, those positions split A as even cnt
	cnt := count_zero(A)
	if cnt == n {
		return int64(n) * int64(M)
	}
	if M&1 == 0 {
		var pref int
		for pref < n && A[pref] == '0' {
			pref++
		}
		suf := n - 1
		for suf >= 0 && A[suf] == '0' {
			suf--
		}
		return int64(pref + n - suf)
	}
	cnt = n - cnt
	// 1011  1011 1011
	if cnt&1 == 1 {
		// no way to split it even
		return 0
	}
	half := cnt >> 1
	var tmp int
	var res int
	for i := 0; i < n && tmp <= half; i++ {
		tmp += int(A[i] - '0')
		if tmp == half {
			res++
		}
	}
	return int64(res)
}

func count_zero(s string) int {
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cnt++
		}
	}
	return cnt
}
