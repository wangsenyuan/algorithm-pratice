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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, s := readTwoNums(reader)
		employees := make([][]int, n)
		for i := 0; i < n; i++ {
			employees[i] = readNNums(reader, 2)
		}
		res := solve(s, employees)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(s int, employees [][]int) int {

	n := len(employees)

	h := (n + 1) / 2
	check := func(median int) bool {
		// 按照最高排序
		slices.SortFunc(employees, func(a, b []int) int {
			return a[1] - b[1]
		})

		// 这部分肯定得不到median
		var sum int
		var i int
		for i < n && employees[i][1] < median {
			sum += employees[i][0]
			i++
		}

		if i >= h {
			// median太大了
			return false
		}

		// 前面的肯定无法得到median的工资
		// 让尽量多的人获得刚好median的工资
		slices.SortFunc(employees[i:], func(a, b []int) int {
			return b[0] - a[0]
		})
		// 后面的肯定大于等于median
		// 把那些下限大的部分，尽量取最小值
		j := i
		for j < n && employees[j][0] >= median {
			sum += employees[j][0]
			j++
		}
		// sum <= s 肯定是成立的
		for j-i+1 <= h && j < n {
			sum += median
			j++
		}
		if j-i+1 < h || sum > s {
			return false
		}
		for j < n {
			// 后面的人取最小值
			sum += employees[j][0]
			j++
		}
		return sum <= s
	}
	l, r := 1, 1<<30
	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}
