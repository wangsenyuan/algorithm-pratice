package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) []int {
	val1 := []int{0, 0}
	val2 := []int{inf, -inf}
	var sum int
	// x 是不包括奇异值的最大/最小的前缀和
	x := []int{0, 0}
	// y是包括了奇异值的最大/最小的前缀和
	y := []int{inf, -inf}

	for _, num := range a {
		sum += num
		if num != -1 && num != 1 {
			y[0], y[1] = x[0], x[1]
			x[0], x[1] = sum, sum
		}
		// val1[0]不包括奇异值的最小的区间和
		val1[0] = min(val1[0], sum-x[1])
		val1[1] = max(val1[1], sum-x[0])
		val2[0] = min(val2[0], sum-y[1])
		val2[1] = max(val2[1], sum-y[0])
		x[0] = min(x[0], sum)
		x[1] = max(x[1], sum)
	}

	var res []int

	if val1[1] < val2[0] {
		res = append(res, gen(val1[0], val1[1])...)
		res = append(res, gen(val2[0], val2[1])...)
	} else if val2[1] < val1[0] {
		res = append(res, gen(val2[0], val2[1])...)
		res = append(res, gen(val1[0], val1[1])...)
	} else {
		res = gen(min(val1[0], val2[0]), max(val2[0], val2[1]))
	}

	return res
}

func gen(l int, r int) []int {
	if l > r {
		return nil
	}
	res := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		res[i-l] = i
	}
	return res
}
