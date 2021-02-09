package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, a1, an := readThreeNums(reader)
	query := func(l1, r1, l2, r2 int) int {
		fmt.Printf("Q %d %d %d %d\n", l1, r1, l2, r2)
		return readNum(reader)
	}

	res := solve(n, a1, an, query)

	var buf bytes.Buffer
	buf.WriteByte('!')
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	fmt.Print(buf.String())
}

func solve(n int, A_1, A_n int, query func(l1, r1, l2, r2 int) int) []int {
	res := make([]int, n+1)
	res[1] = A_1
	res[n] = A_n

	for i := 1; i+2 < n; i += 2 {
		x := query(i, i+1, i+1, i+2)
		y := query(i+1, i+2, i, i+1)
		res[i+2] = res[i] - x + y
	}

	for i := n - 2; i > 1; i -= 2 {
		x := query(i, i+1, i+1, i+2)
		y := query(i+1, i+2, i, i+1)
		res[i] = res[i+2] + x - y
	}
	return res[1:]
}
