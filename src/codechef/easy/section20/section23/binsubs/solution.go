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
		s, _ := reader.ReadBytes('\n')
		res := solve(n, s)
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

func solve(n int, S []byte) int {
	sum := make([]int, n)

	for i := 0; i < n; i++ {
		sum[i] = int(S[i] - '0')
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	pivots := make([]int, n)

	pivots[n-1] = n

	for i := n - 2; i >= 0; i-- {
		pivots[i] = pivots[i+1]
		if S[i] == '1' && S[i+1] == '0' {
			pivots[i] = i
		}
	}

	best := n

	for i := 0; i < n; i++ {
		// zeros < i && ones >= pivots[i]
		var a int
		if i > 0 {
			a = sum[i-1]
		}
		j := pivots[i]
		var b int
		if j < n {
			b = (n - 1 - j) - (sum[n-1] - sum[j])
		}
		if a+b < best {
			best = a + b
		}
	}
	return best
}
