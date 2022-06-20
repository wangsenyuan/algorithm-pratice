package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		reader.ReadBytes('\n')
		S := readString(reader)
		A := readString(reader)
		ok, R := solve(S, A)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(R)
			buf.WriteByte('\n')
		}
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
		if s[i] == '\n' {
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

const X = "abcde"

func solve(S, A string) (bool, string) {
	n := len(S)
	m := len(A)

	var i, j int

	for i < n && j < m {
		if S[i] == A[j] {
			i++
			j++
		} else {
			i++
		}
	}

	if j == m {
		// A is already a subseq of S
		return false, ""
	}

	buf := []byte(S)

	i, j = 0, 0

	for i < n && j < m {
		if S[i] == A[j] {
			i++
			j++
		} else if S[i] != '?' {
			i++
		} else {
			// S[i] = '?' and A[j] = {a, b, c, d, e}
			// 只要和A[j]不一致就可以吗？
			buf[i] = 'a'
			if A[j] == 'a' {
				buf[i] = 'b'
			}
			i++
		}
	}

	return true, string(buf)
}
