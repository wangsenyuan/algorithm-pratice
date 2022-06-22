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
		n := readNum(reader)
		ok, A, B := solve(n)
		if !ok {
			buf.WriteString("-1\n")
			continue
		}
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", A[i]))
		}
		buf.WriteByte('\n')
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", B[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int) (bool, []int, []int) {
	if n == 1 {
		return true, []int{1}, []int{1}
	}
	if n&1 == 1 {
		return false, nil, nil
	}
	// 1 & a = 2 & b = 3 & c
	// else 1 & a = 0, only solution would be 1 & a = 0, 2 & b = 0, 3 & c = 0
	// 1， 2， 3， 4
	// 2， 1， 4， 3
	// 1， 2， 3， 4， 5， 6
	// 6， 5， 4， 3， 2 ，1
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	// 2, 1, 4, 3, 10, 9, 8, 7, 6, 5
	A := make([]int, n+1)
	B := make([]int, n+1)

	for i := 0; i <= n; i++ {
		A[i] = i
	}

	var p int64 = 1
	for p <= int64(n) {
		p *= 2
	}

	for i := n; i > 0; i-- {
		if B[i] != 0 {
			continue
		}
		for 2*int64(i) < p {
			p /= 2
		}
		B[i] = int(p-1) - i
		B[int(p-1)-i] = i
	}
	return true, A[1:], B[1:]
}
