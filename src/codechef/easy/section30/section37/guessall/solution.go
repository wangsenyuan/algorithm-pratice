package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

	ask := func(x int) int64 {
		fmt.Printf("? %d\n", x)
		s, _ := reader.ReadBytes('\n')
		var r int64
		readInt64(s, 0, &r)
		return r
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--

		k := readNum(reader)
		n := readNum(reader)
		B := readNNums(reader, n)

		res := solve(k, n, B, ask)

		var buf bytes.Buffer
		buf.WriteByte('!')
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
}

func solve(k int, n int, B []int, F func(int) int64) []int64 {
	X := make([]int64, k+1)

	for _, num := range B {
		X[num%(k+1)] = 1
	}
	var cnt int
	for i := 0; i < k; i++ {
		if X[i] == 1 {
			X[i] = F(i)
			cnt++
		}
	}
	var sum int64
	if X[k] == 1 {
		if cnt == k {
			for i := 0; i < k; i++ {
				sum += int64(X[i])
			}
			sum *= -1
		} else {
			sum = F(k)
		}
		X[k] = sum
	}

	res := make([]int64, n)

	for i, num := range B {
		res[i] = X[num%(k+1)]
	}

	return res
}
