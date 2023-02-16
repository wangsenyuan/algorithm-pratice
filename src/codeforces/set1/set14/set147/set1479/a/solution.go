package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(i int) int {
		fmt.Printf("? %d\n", i)
		return readNum(reader)
	}

	n := readNum(reader)

	res := solve(n, ask)

	fmt.Printf("! %d\n", res)
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

// 不变性是在区间[l...r] 中间肯定存在局部最小值
func solve(n int, ask func(int) int) int {
	if n < 2 {
		return 1
	}
	// ask(i) = a[i]
	// ask(1) = a[1]
	// ask(n) = a[n]
	A := make([]int, n+1)

	cache := func(i int) int {
		if A[i] == 0 {
			A[i] = ask(i)
		}
		return A[i]
	}

	l, r := 1, n

	for l < r {
		mid := (l + r) / 2
		x := cache(mid)
		y := cache(mid + 1)
		if x > y {
			l = mid + 1
		} else {
			// x < y
			r = mid
		}
	}

	return r
}
