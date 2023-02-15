package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var pos int
		var t int
		pos = readInt(s, 0, &t)
		if t == 1 {
			var j, x int
			pos = readInt(s, pos+1, &j)
			readInt(s, pos+1, &x)
			Q[i] = []int{t, j, x}
		} else {
			var x int
			readInt(s, pos+1, &x)
			Q[i] = []int{t, x}
		}
	}
	res := solve(A, Q)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int, Q [][]int) []int64 {
	n := len(A)
	updated_at := make([]int, n)
	for i := 0; i < n; i++ {
		updated_at[i] = -1
	}

	last_full := []int{-2, 0}

	var sum int64

	for i := 0; i < n; i++ {
		sum += int64(A[i])
	}

	res := make([]int64, len(Q))

	for i, cur := range Q {
		if cur[0] == 1 {
			j, x := cur[1], cur[2]
			j--
			// single update
			if last_full[0] < updated_at[j] {
				// use the value at A[j]
				sum -= int64(A[j])
			} else {
				// last_full[0] > update_at[j]
				sum -= int64(last_full[1])
			}
			A[j] = x
			sum += int64(A[j])
			updated_at[j] = i
		} else {
			// full update
			x := cur[1]
			sum = int64(n) * int64(x)
			last_full = []int{i, x}
		}
		res[i] = sum
	}

	return res
}
