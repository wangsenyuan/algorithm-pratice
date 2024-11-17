package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	prices := readNNums(reader, n)
	first := readNNums(reader, 2)
	second := readNNums(reader, 2)
	k := readNum(reader)
	return solve(prices, first, second, k)
}

func solve(prices []int, first []int, second []int, k int) int {
	if first[0] < second[0] {
		first, second = second, first
	}
	// x >= y
	sort.Ints(prices)
	reverse(prices)

	n := len(prices)
	pref := make([]int, n+1)
	for i, x := range prices {
		pref[i+1] = pref[i] + x/100
	}

	a, b := first[1], second[1]
	x, y := first[0], second[0]
	c := lcm(a, b)

	check := func(m int) bool {
		// 使用m个是否能够得到k的收益
		cnt1 := m / c
		// 前z个的收益是 (x + y)
		sum := (x + y) * pref[cnt1]
		if c != a {
			cnt2 := m/a - m/c
			sum += x * (pref[cnt2+cnt1] - pref[cnt1])
			cnt1 += cnt2
		}
		if c != b {
			cnt2 := m/b - m/c
			sum += y * (pref[cnt2+cnt1] - pref[cnt1])
		}
		return sum >= k
	}

	if !check(n) {
		return -1
	}
	return sort.Search(n, check)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
