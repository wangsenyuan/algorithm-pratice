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
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int) int {
	// 找出三种数
	// a = 0, b = 0, others
	cnt := make([]int, 2)
	n := len(a)
	freq := make(map[pair]int)
	for i := range n {
		if a[i] == 0 && b[i] == 0 {
			// 这个对d没有要求
			cnt[0]++
			continue
		}
		if a[i] == 0 {
			// 只能 b[i] = 0
			// 所以任何的d都没法让c = 0
			continue
		}
		// a[i] != 0
		if b[i] == 0 {
			// 那么只能 d = 0
			cnt[1]++
		} else {
			g := gcd(abs(a[i]), abs(b[i]))
			x := abs(a[i]) / g
			y := abs(b[i]) / g
			if sign(a[i]) != sign(b[i]) {
				x *= -1
			}
			freq[pair{x, y}]++
		}
	}
	res := cnt[1]
	for _, v := range freq {
		res = max(res, v)
	}
	return res + cnt[0]
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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
