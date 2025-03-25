package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	n, l, r := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(a, l, r)
}

func solve(a []int, l int, r int) int {
	n := len(a)

	if n&1 == 0 {
		var tmp int
		for i := 1; i <= (n+1)/2; i++ {
			tmp ^= a[i-1]
		}
		a = append(a, tmp)
		n++
	}
	p := make([]int, 2*n+1)
	for i := 1; i <= n; i++ {
		p[i] = p[i-1] ^ a[i-1]
	}
	// n is odd
	for i := n + 1; i <= 2*n; i++ {
		a = append(a, p[i/2])
		p[i] = p[i-1] ^ a[i-1]
	}

	if l <= 2*n {
		return a[l-1]
	}
	var res int
	for l > 2*n {
		l /= 2
		res ^= p[n]
		if l%2 == 1 {
			return res
		}
	}
	return res ^ a[l-1]
}
