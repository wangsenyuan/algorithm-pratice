package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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
	n, r, k := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(a, r, k)
}

func solve(a []int, radius int, k int) int {
	n := len(a)

	b := make([]int, n)
	check := func(x int) bool {
		copy(b, a)
		var sum int
		k1 := k
		for l, r, i := 0, 0, 0; i < n; i++ {
			for r < n && r <= i+radius {
				sum += b[r]
				r++
			}
			for l < i-radius {
				sum -= b[l]
				l++
			}
			if sum < x {
				d := min(k1, x-sum)
				sum += d
				b[r-1] += d
				k1 -= d
			}
			if sum < x {
				return false
			}
		}
		return true
	}

	l, r := 0, 1<<60
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}
