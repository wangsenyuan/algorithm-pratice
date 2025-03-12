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
	nums := readNNums(reader, 4)
	return solve(nums[0], nums[1], nums[2], nums[3])
}

func solve(l int, r int, i int, k int) int {

	f := func(x int) int {
		n := uint64(x)
		rem := n % 4
		if rem == 0 {
			return x
		}
		if rem == 1 {
			return 1
		}
		if rem == 2 {
			return x + 1
		}
		return 0
	}

	g := func(l int, r int) int {
		return f(r) ^ f(l-1)
	}

	res := g((l-k+(1<<i)-1)>>i, (r-k)>>i) << i

	count := func(x int) int {
		n := uint64(x)
		n -= uint64(k)
		n /= (1 << i)

		return int(n)
	}

	cnt := count(r) - count(l-1)

	if cnt&1 == 1 {
		res ^= k
	}

	res ^= g(l, r)

	return res
}
