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
		a := readNNums(reader, n)
		buf.WriteString(fmt.Sprintf("%d\n", solve(a)))
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

const inf = 1 << 30

func solve(a []int) int {
	n := len(a)

	suff := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suff[i] = suff[i+1]
		if i%2 == 0 {
			suff[i] += a[i]
		}
	}

	best := []int{0, 0}

	// sum[0] for even, sum[1] for odd
	sum := []int{0, 0}

	res := suff[0]

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum[0] += a[i]
			// 始终可以在某个偶数位，将前面的全部反转
			res = max(res, sum[1])
		} else {
			sum[1] += a[i]
		}

		res = max(res, sum[1]+best[i&1]+suff[i+1])

		if best[i&1] < sum[0]-sum[1] {
			best[i&1] = sum[0] - sum[1]
		}
	}

	return res
}
