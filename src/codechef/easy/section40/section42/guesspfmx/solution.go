package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(Q []int) []int {
		var buf bytes.Buffer
		buf.WriteString("?")
		for i := 0; i < len(Q); i++ {
			buf.WriteString(fmt.Sprintf(" %d", Q[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		s, _ := reader.ReadBytes('\n')
		var m int
		pos := readInt(s, 0, &m)
		R := make([]int, m)
		for i := 0; i < m; i++ {
			pos = readInt(s, pos+1, &R[i])
		}
		return R
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)

		res := solve(n, ask)

		var buf bytes.Buffer

		buf.WriteString("!")
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())

		readNum(reader)
	}
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

func solve(N int, ask func(Q []int) []int) []int {
	// 先找出N在哪里？
	// 如果Q = [1....n]
	// R [1... i, ... ]， 如果出现了i, 那么 P[i] = max(P[0]....P[i])
	// 那么最后一个i就是N所在的位置
	// P = [2,4,3,1], Q = [1, 2, 3, 4]
	// R = [1,2] 所以有 P[2] = N
	// 然后把2放到最后去 Q = [1, 3, 4, 2]
	// 2, 3, 1, 4
	Q := make([]int, N)
	for i := 0; i < N; i++ {
		Q[i] = i + 1
	}
	P := make([]int, N)
	for i := N; i > 1; i-- {
		R := ask(Q)
		l := len(R)
		x := R[l-1-(N-i)]
		j := Q[x-1]
		P[j-1] = i

		reverse(Q[x-1 : i])
		reverse(Q[x-1 : i-1])
	}
	for i := 0; i < N; i++ {
		if P[i] == 0 {
			P[i] = 1
			break
		}
	}
	return P
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
