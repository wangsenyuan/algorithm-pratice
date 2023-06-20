package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNNums(reader, 3)
	A := make([][]int, 3)
	for i := 0; i < 3; i++ {
		A[i] = readNNums(reader, n[i])
	}
	res := solve(A)
	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(A [][]int) int64 {
	// a & b => a - b
	// 有没有可能是把所有的数都加起来呢？不行
	// aij都是大于0的，应该尽可能的让它们最后的符号是+
	// 要怎么做到这一点呢？
	// 这里一共能进行的操作数是  len(A) + len(B) + len(C) - 1
	// 如果一个数，被操作了偶数次，(0, 2, 4...)次，那么它就是 +， 否则就是 -
	for i := 0; i < 3; i++ {
		sort.Ints(A[i])
	}

	calc := func() int64 {
		ans := sum(A[0])
		m1 := int64(A[1][0])
		m2 := int64(A[2][0])
		s1 := sum(A[1][1:])
		s2 := sum(A[2][1:])
		ans += max(s2-m1+s1-m2, m2+s2-m1-s1, m1+s1-m2-s2)
		return ans
	}

	a := calc()
	A[0], A[1] = A[1], A[0]
	b := calc()
	A[0], A[2] = A[2], A[0]
	c := calc()
	return max(a, b, c)
}

func max(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func sum(arr []int) int64 {
	var res int64
	for _, num := range arr {
		res += int64(num)
	}
	return res
}
