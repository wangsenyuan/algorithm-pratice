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
		boxes := make([][]int, 3)
		for i := 0; i < 3; i++ {
			boxes[i] = readNNums(reader, 3)
		}
		res := solve(n, boxes)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

func solve(n int, boxes [][]int) int {
	// boxes[i] R, G, B
	var res int

	swap := func(i int) int {
		// try to make sure box i, has all color i (0 for R, 1 for G, 2 for B)
		var res int

		j := (i + 1) % 3
		k := (i + 2) % 3

		a := boxes[i][j]
		b := boxes[j][i]

		x := min(a, b)
		res += x
		boxes[i][j] -= x
		boxes[j][i] -= x

		if boxes[i][j] > 0 {
			// then boxes[j][k], boxes[k][i]
			res += 2 * boxes[i][j]
			boxes[j][k] -= boxes[i][j]
			boxes[k][i] -= boxes[i][j]
			boxes[i][j] = 0
		}

		a = boxes[i][k]
		b = boxes[k][i]
		x = min(a, b)
		res += x
		boxes[i][k] -= x
		boxes[k][i] -= x
		if boxes[i][k] > 0 {
			res += 2 * boxes[i][k]
			boxes[k][j] -= boxes[i][k]
			boxes[j][i] -= boxes[i][k]
			boxes[i][k] = 0
		}

		return res
	}

	res += swap(0)
	res += swap(1)
	res += swap(2)

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
