package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, x := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(x, a)
}

func solve(x int, a []int) int {
	n := len(a)
	if n == 1 {
		return 0
	}
	sort.Ints(a)

	get := func(l int, r int) int {
		// a[l] < a[r]
		// r - l + 1 == n - 1
		var res1 int
		if a[l] <= x {
			res1 = x - a[l]
		} else {
			res1 = inf
		}
		if a[r] > x {
			// 必须返回
			res1 += a[r] - a[l]
		}
		var res2 int
		if x <= a[r] {
			res2 = a[r] - x
		} else {
			res2 = inf
		}
		if a[l] < x {
			res2 += a[r] - a[l]
		}
		return min(res1, res2)
	}

	return min(get(0, n-2), get(1, n-1))
}

const inf = 1 << 60
