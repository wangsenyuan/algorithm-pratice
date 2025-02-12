package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const H = 30

func solve(a []int) int {
	n := len(a)

	arr := make([]int, n)
	check := func(x int) int {
		copy(arr, a)

		var sum int
		for _, num := range arr {
			sum += num
		}
		var ops int
		for sum > n*x {
			cur := sum
			for i := 0; i < n && sum > n*x; i++ {
				if arr[i] > x {
					diff := (arr[i] - x + 1) / 2 * 2
					arr[i] -= diff
					arr[(i+1)%n] += diff / 2
					sum -= diff / 2
					ops += diff / 2
				}
			}
			if sum == cur {
				return -1
			}
		}
		if sum == n*x && slices.Min(arr) == x {
			return ops
		}
		return -1
	}

	if check(1) < 0 {
		return -1
	}

	l, r := 1, slices.Max(a)+1

	for l < r {
		mid := (l + r) / 2
		if check(mid) >= 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	r--
	return check(r)
}
