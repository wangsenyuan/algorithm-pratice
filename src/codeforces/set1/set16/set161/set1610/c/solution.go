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
		n := readNum(reader)
		A := make([]int, n)
		B := make([]int, n)
		for i := 0; i < n; i++ {
			A[i], B[i] = readTwoNums(reader)
		}
		res := solve(A, B)
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

func solve(A []int, B []int) int {
	// 对于i来说，他有i dollars，他happy <= 至多有a[i]个人richer，至多b[i]个人 poorer
	// 换句话说，就是如果i 参加了，那么在i后面至多有a[i]，在i前面至多有b[i]
	// 假设第一个是（最小的index)u, 那么能参加的人数 <= a[u]
	//  第二个是 v, 如果 a[v] < a[u] - 1, 则不能满足条件
	//    如果  a[v] >= a[u]-1, 且 b[v] >= 1, 则ok
	//    三个是 w,    a[w] >= a[u] - 2, 且 b[w] >= 2
	// 在给定人数的时候， 是否可以从最小的下标开始进行检查呢？
	// 假设u， v,都满足a[u] >= x, a[v] >= x, 且 u < v
	// 如果 b[v] >= 1, 则从 u 开始，同时也可以包括v
	// 如果 b[v] = 0, 则将v替换成u，不会更差
	n := len(A)
	check := func(x int) bool {
		// x表示还剩余能增加几个人
		// l表示已经增加了几个人
		var l int
		for i := 1; i <= n && x > 0; i++ {
			// 因为增加自己，还需要增加x-1个人
			if A[i-1] >= x-1 && B[i-1] >= l {
				l++
				x--
			}
		}
		return x == 0
	}

	l, r := 1, n+1

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
