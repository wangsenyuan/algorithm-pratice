package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s, _ := reader.ReadString('\n')
	t, _ := reader.ReadString('\n')
	fmt.Println(solve(n, s, t))
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

func solve(n int, s, t string) int {
	A := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		if s[i] > t[i] {
			A[i] = 1
		} else if s[i] < t[i] {
			A[i] = -1
		}
		sum += A[i]
	}

	if sum != 0 {
		return -1
	}

	get := func(x int) int {
		var cur int
		var best int

		for i := 0; i < n; i++ {
			cur += x * A[i]
			if best < cur {
				best = cur
			}
			if cur < 0 {
				cur = 0
			}
		}
		return best
	}

	return max(get(1), get(-1))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
