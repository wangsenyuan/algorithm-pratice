package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) []int {
	params := readNNums(reader, 5)
	return solve(params)
}

func solve(params []int) []int {
	t1, t2 := params[0], params[1]
	x1, x2 := params[2], params[3]
	t0 := params[4]

	if t1 == t2 {
		return []int{x1, x2}
	}

	check := func(y1 int, y2 int) bool {
		// t = (t1 * y1 + t2 * y2) / (y1 + y2) >= t0
		if y1+y2 == 0 {
			return true
		}
		return t1*y1+t2*y2 >= t0*(y1+y2)
	}

	find := func(y1 int) int {
		if y1 == 0 {
			// 当只有热水的时候，温度始终是t2, 所以用最大值
			return x2
		}
		if !check(y1, x2) {
			return -1
		}
		l, r := 0, x2
		for l < r {
			mid := (l + r) / 2
			if check(y1, mid) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r
	}

	comp := func(a []int, b []int) bool {
		if a[0]+a[1] == 0 {
			return true
		}
		if (t1*a[0]+t2*a[1])*(b[0]+b[1]) > (t1*b[0]+t2*b[1])*(a[0]+a[1]) {
			return true
		}

		if (t1*a[0]+t2*a[1])*(b[0]+b[1]) == (t1*b[0]+t2*b[1])*(a[0]+a[1]) && a[0]+a[1] < b[0]+b[1] {
			return true
		}

		return false
	}

	// 0/0 肯定满足条件
	res := []int{0, 0}
	for y1 := 0; y1 <= x1; y1++ {
		y2 := find(y1)
		if y2 < 0 {
			continue
		}

		if comp(res, []int{y1, y2}) {
			res = []int{y1, y2}
		}
	}

	return res
}
