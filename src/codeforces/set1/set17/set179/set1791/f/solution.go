package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		ops := make([][]int, m)
		for i := 0; i < m; i++ {
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				var l, r int
				pos := readInt(s, 2, &l)
				readInt(s, pos+1, &r)
				ops[i] = []int{1, l, r}
			} else {
				var j int
				readInt(s, 2, &j)
				ops[i] = []int{2, j}
			}
		}
		res := solve(a, ops)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int, ops [][]int) []int {
	n := len(a)

	// 要将区间[l..r]内的数都变成它们的digits sum
	// 假设a[i] = 99999999 (8个9)
	// 一次操作后 a[i] = 72
	// 两次操作后 a[i] = 9
	// 也就是操作3次后，它就变成了个位数，就不再变化了，所以只需要记录每一位的变化的次数就可以了
	// use bit
	arr := make([]int, n+1)

	update := func(i int, v int) {
		i++
		for i <= n {
			arr[i] += v
			i += i & -i
		}
	}

	get := func(i int) int {
		i++
		var res int
		for i > 0 {
			res += arr[i]
			i -= i & -i
		}
		return res
	}

	find := func(i int, cnt int) int {
		v := a[i]
		for cnt > 0 && v >= 10 {
			v = digitsSum(v)
			cnt--
		}

		return v
	}

	var ans []int

	for _, op := range ops {
		if op[0] == 1 {
			l, r := op[1], op[2]
			update(l-1, 1)
			update(r, -1)
		} else {
			i := op[1] - 1
			cnt := get(i)
			ans = append(ans, find(i, cnt))
		}
	}

	return ans
}

func digitsSum(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
