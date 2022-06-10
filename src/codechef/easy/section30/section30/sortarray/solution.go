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
		A := readNNums(reader, n)
		res := solve(A)
		m := len(res) / 2
		buf.WriteString(fmt.Sprintf("%d\n", m))
		for i := 0; i < m; i++ {
			a := res[2*i]
			b := res[2*i+1]
			buf.WriteString(fmt.Sprintf("%d %d\n", a[0], a[1]))
			for _, j := range b {
				buf.WriteString(fmt.Sprintf("%d ", j))
			}
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

func solve(A []int) [][]int {
	n := len(A)

	diff := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		// B[i] = 1e5 + i - A[i]
		diff[i] = 100000 + i - A[i]
	}

	var ops [][]int
	for i := 0; i < 20; i++ {
		// x := 1 << i
		X := 1 << uint(i)
		var S []int
		for j := 0; j < n; j++ {
			if (diff[j]>>uint(i))&1 == 1 {
				S = append(S, j+1)
				A[j] += X
			}
		}
		if len(S) > 0 {
			ops = append(ops, []int{len(S), X})
			ops = append(ops, S)
		}
		ok := true
		for j := 1; j < n; j++ {
			if A[j] <= A[j-1] {
				ok = false
			}
		}
		if ok {
			break
		}
	}
	return ops
}
