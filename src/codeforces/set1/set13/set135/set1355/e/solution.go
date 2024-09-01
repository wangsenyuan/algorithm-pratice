package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	lines := readNNums(reader, 4)
	n := lines[0]
	a, r, m := lines[1], lines[2], lines[3]

	h := readNNums(reader, n)

	res := solve(h, a, r, m)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const inf = 1 << 30

func solve(h []int, add int, rem int, move int) int {
	move = min(move, add+rem)

	sort.Ints(h)
	// 直接修改即可
	var tot int
	for _, num := range h {
		tot += num
	}
	n := len(h)

	v := make([]int, n)

	calc := func(x int) int {
		sum := tot
		var res int
		copy(v, h)
		if sum > n*x {
			// 删除多余的先
			tmp := sum - n*x
			res = tmp * rem
			for i := n - 1; i >= 0 && tmp > 0; i-- {
				y := min(tmp, v[i]-x)
				v[i] -= y
				tmp -= y
			}
		} else if sum < n*x {
			tmp := n*x - sum
			res = tmp * add
			for i := 0; i < n && tmp > 0; i++ {
				t := min(tmp, x-v[i])
				v[i] += t
				tmp -= t
			}
		}
		// sum = n * x
		// 剩余的只需要移动？

		var cnt int
		for i := 0; i < n; i++ {
			if v[i] < x {
				cnt += x - v[i]
			}
		}

		res += cnt * move

		return res
	}

	l, r := slices.Min(h), slices.Max(h)

	for r-l >= 3 {
		m1 := l + (r-l)/3
		m2 := m1 + (r-l)/3
		ans1 := calc(m1)
		ans2 := calc(m2)
		if ans1 < ans2 {
			r = m2
		} else {
			l = m1
		}
	}

	// l, l+1, l+2
	return min(calc(l), calc(l+1), calc(l+2))
}

func abs(num int) int {
	return max(num, -num)
}
