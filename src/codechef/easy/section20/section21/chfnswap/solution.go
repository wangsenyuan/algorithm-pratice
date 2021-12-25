package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		buf.WriteString(fmt.Sprintf("%d\n", solve(n)))
	}
	fmt.Print(buf.String())
}

func solve(n int) int64 {
	N := int64(n)
	sum := N * (N + 1) / 2
	if sum%2 != 0 {
		return 0
	}
	// find some m, that sum(m) >= sum - sum(m)

	m := sort.Search(n+1, func(m int) bool {
		M := int64(m)
		x := M * (M + 1) / 2
		return 2*x >= sum
	})

	if m > n {
		return 0
	}

	res := int64(n - m)

	if int64(m)*int64(m+1) == sum {
		// divides equals
		res += int64(m)*int64(m-1)/2 + int64(n-m)*int64(n-m-1)/2
	} else {
		// sum(m) * 2 > sum, sum(m - 1) * 2 < sum
		res++
	}

	return res
}
