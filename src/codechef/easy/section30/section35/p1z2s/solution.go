package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	t := readNNums(reader, n)
	res := solve(t)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(t []int) int {
	n := len(t)
	return max(n, (sum(t)+1)/2)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(t []int) int {
	n := len(t)
	// 假设buy x tickents, 2 * x vouchers, 3 * x >= sum(t)
	// 假设买了x张票后，那么  x >= n,
	// 同时获得了x张voucher

	// sort.Ints(t)

	check := func(x int) bool {
		y := x
		i := n - 1
		for ; i >= 0 && x > 0; i-- {
			x--
			// want to visit more a times
			a := t[i] - 1
			if a <= y {
				// use voucher first
				y -= a
				continue
			}
			// a > y
			a -= y
			y = 0
			x -= a
		}

		return x >= 0 && i < 0
	}

	lo, hi := n, sum(t)

	for lo < hi {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}

func sum(arr []int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
