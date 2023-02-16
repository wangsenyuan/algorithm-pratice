package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)

	query := func(l, r int) int {
		fmt.Printf("? %d %d\n", l+1, r+1)
		return readNum(reader) - 1
	}

	res := solve(n, query) + 1

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

func solve(n int, query func(int, int) int) int {
	smax := query(0, n-1)
	// if max is between [1...smax-1] or [smax + 1...r]
	if smax == 0 || query(0, smax) != smax {
		return solve1(smax, n, query)
	}

	return solve2(smax, n, query)
}

func solve1(smax int, n int, query func(int, int) int) int {
	// find smallest position after smax, having second maximum != smax
	left, right := smax, n-1

	for right-left > 1 {
		mid := (left + right) / 2
		if query(smax, mid) == smax {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func solve2(smax int, n int, query func(int, int) int) int {
	// find largest position before smax, having second max == smax

	left, right := 0, smax

	for right-left > 1 {
		mid := (left + right) / 2
		if query(mid, smax) == smax {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
