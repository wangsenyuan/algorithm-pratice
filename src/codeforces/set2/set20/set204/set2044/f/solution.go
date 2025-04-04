package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer

	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []bool {
	n, m, q := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	x := make([]int, q)
	for i := range q {
		x[i] = readNum(reader)
	}
	return solve(a, b, x)
}

func solve(a []int, b []int, x []int) []bool {

	get := func(arr []int) map[int]bool {
		var sum int
		for _, num := range arr {
			sum += num
		}
		res := make(map[int]bool)
		for _, num := range arr {
			res[sum-num] = true
		}
		return res
	}

	u := get(a)
	v := get(b)

	ok := func(d int, x int, y int) bool {
		if d > 0 {
			if u[x] && v[y] || u[-x] && v[-y] {
				return true
			}
		} else {
			if u[-x] && v[y] || u[x] && v[-y] {
				return true
			}
		}
		return false
	}

	check := func(num int) bool {
		d := sign(num)
		num = abs(num)
		for x := 1; x*x <= num; x++ {
			if num%x == 0 {
				y := num / x
				if ok(d, x, y) {
					return true
				}
				if ok(d, y, x) {
					return true
				}
			}
		}
		return false
	}

	ans := make([]bool, len(x))

	for i, num := range x {
		ans[i] = check(num)
	}

	return ans
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}

func abs(num int) int {
	return max(num, -num)
}
