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
	var buf bytes.Buffer
	for {
		var n uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &n)
		if n == 0 {
			break
		}
		res := solve(int64(n))
		if res < 0 {
			buf.WriteString("L\n")
		} else if res == 0 {
			buf.WriteString("D\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", res))
		}
	}
	fmt.Print(buf.String())
}

const MAX_W = 10000000000

var fib []int64

func init() {
	fib = make([]int64, 0, 102)
	fib = append(fib, 0)
	fib = append(fib, 1)
	i := 2
	for fib[i-1] < MAX_W {
		fib = append(fib, fib[i-2]+fib[i-1])
		i++
	}
}

func solve(N int64) int64 {
	if N == 1 {
		return 0
	}
	save := N

	var zeck int64
	it := len(fib)
	for N > 0 {
		if N >= fib[it-1] {
			N -= fib[it-1]
			zeck = fib[it-1]
		}
		it--
	}
	if save-1 < zeck {
		return -1
	}
	return zeck
}
