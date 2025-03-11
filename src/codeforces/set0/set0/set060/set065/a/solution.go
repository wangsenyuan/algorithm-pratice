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

func process(reader *bufio.Reader) string {
	nums := readNNums(reader, 6)
	return solve(nums)
}

func solve(nums []int) string {
	a, b, c, d, e, f := nums[0], nums[1], nums[2], nums[3], nums[4], nums[5]

	if c == 0 && d > 0 || a == 0 && b > 0 && d > 0 {
		return "Ron"
	}

	if b == 0 || d == 0 || f == 0 {
		return "Hermione"
	}
	if a == 0 || c == 0 || e == 0 {
		// 无中生有
		return "Ron"
	}
	x := lca(d, e)
	f *= x / e
	e = x
	c *= x / d
	d = x
	y := lca(b, c)
	a *= y / b
	b = y
	e *= y / c
	f *= y / c
	d *= y / c
	c = y
	if f > a {
		return "Ron"
	}
	return "Hermione"
}

func lca(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
