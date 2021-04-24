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
		C := readNNums(reader, n)
		res := solve(n, C)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(n int, C []int) int64 {
	/**
	分成两部分偶数和奇数部分；
	偶数部分用于垂直移动，奇数部分用于水平移动；
	那么其中个部分的最小值应该移动最多n - m + 1, 其他的部分移动m-1
	*/
	x, y := C[0], C[1]
	N := int64(n)
	best := N * (int64(x) + int64(y))
	sh := int64(x)
	sv := int64(y)
	for i := 2; i < n; {
		x = min(x, C[i])
		sh += int64(C[i])
		// if take i segs, and i is horizontal
		m := i/2 + 1
		// horiontal part
		tmp := int64(n-m+1)*int64(x) + sh - int64(x)
		m--
		// vertical part
		tmp += int64(n-m+1)*int64(y) + sv - int64(y)
		best = min2(best, tmp)
		i++
		if i < n {
			y = min(y, C[i])
			sv += int64(C[i])
			// if i == 3, then m = 2
			m = i/2 + 1

			tmp = int64(n-m+1)*int64(x) + sh - int64(x)
			tmp += int64(n-m+1)*int64(y) + sv - int64(y)
			best = min2(best, tmp)
			i++
		}
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min2(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
