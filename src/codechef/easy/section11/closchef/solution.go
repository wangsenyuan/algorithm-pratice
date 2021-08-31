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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
		if res {
			buf.WriteString("1\n")
		} else {
			buf.WriteString("0\n")
		}
	}
	fmt.Print(buf.String())
}

func solve(n int, A []int) bool {
	cnt := make([]int, 3)
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			cnt[0]++
		} else if A[i] == 1 {
			cnt[1]++
		} else if A[i] == -1 {
			cnt[2]++
		}
	}

	n -= cnt[0]
	n -= cnt[1]
	if n <= 1 {
		return true
	}
	if cnt[2] == 0 {
		return false
	}
	n -= cnt[2]
	if n == 0 {
		if cnt[2] == 1 {
			// -1
			return true
		}
		// -1, -1, 1
		return cnt[1] > 0
	}

	return false
}
