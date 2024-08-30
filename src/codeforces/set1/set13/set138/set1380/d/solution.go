package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	x, k, y := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)

	res := solve(n, a, b, k, x, y)

	fmt.Println(res)
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

func solve(n int, a []int, b []int, k int, cost1 int, cost2 int) int {
	for i := 0; i < len(a); i++ {
		a[i]--
	}
	for i := 0; i < len(b); i++ {
		b[i]--
	}

	pos := make([]int, n)

	for i, x := range a {
		pos[x] = i
	}

	get := func(i int) int {
		if i < 0 || i == n {
			return -1
		}
		return a[i]
	}

	process := func(l int, r int) int {
		if l+1 == r {
			return 0
		}
		x := max(get(l), get(r))
		y := -1
		for i := l + 1; i < r; i++ {
			y = max(y, a[i])
		}
		dist := r - l - 1
		if y < x {
			// 可以安全的使用操作2
			if cost1 >= k*cost2 {
				// 使用操作2更优
				return dist * cost2
			}
			// 优先使用操作1
			cnt, rem := dist/k, dist%k
			return cnt*cost1 + rem*cost2
		}
		// y > x, 至少要一次操作1
		if dist < k {
			// 没法使用操作1
			return -1
		}
		res := cost1
		dist -= k
		if cost1 >= k*cost2 {
			return res + dist*cost2
		}
		cnt, rem := dist/k, dist%k
		return res + cnt*cost1 + rem*cost2
	}

	prev := -1
	var res int
	for _, x := range b {
		cur := pos[x]
		if cur <= prev {
			return -1
		}
		tmp := process(prev, cur)
		if tmp < 0 {
			return -1
		}
		res += tmp
		prev = cur
	}
	tmp := process(prev, n)
	if tmp < 0 {
		return -1
	}
	res += tmp
	return res
}
