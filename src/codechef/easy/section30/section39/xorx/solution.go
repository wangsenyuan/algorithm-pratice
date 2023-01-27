package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	update := func(i int, x int) int {
		fmt.Printf("%d %d\n", i, x)
		return readNum(reader)
	}

	solve(n, update)

	fmt.Println("0")
	readNum(reader)
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(N int, update func(x int, y int) int) {
	// N <= 3000, A[i] < 1 << 30, A[i] distinct
	// query(x, y), x (- [1...n]
	// y < 1 << 30
	//  change A[x] = A[x] ^ y, and returns the max value
	// set(1, 0) => 得到最大值，x
	// 然后，1。。。。n set(i, x) 得到y
	//           如果i就是最大值所在的index, 那么 set(i, x), 将会把i设置位0， 从而 y < x
	//              i, 不是， set(i, x)
	//                         1. 仍然返回 x （= y）,  A[i] = A[i] ^ x <= x
	//                         2.         x != y, 只能是 y > x, A[i] = y ^ x
	// 可以得到一个结论是，通过这种方式可以找到最大值所在的i (y < x), 且能够部分的知道某些A[i] = y ^ x (y > x)
	// 这个无法保证在 1e5次能完成
	// set(1, 2**30 -1), A[1] ^= mask
	// 返回的结果有两种可能，一种 x 不是 A[1]， 一种是 x是A[1]
	// 如果 先 set(1, 0) => x, 然后 set(1, x) 会得到什么信息呢？
	//     如果A[1] = x, 那么 A[1] ^ x = 0，得到的y肯定是另外一个值 y < x
	//             != x， 那么 A[1] ^ x, .......有可能重新得到x
	//                                                    y,  y > x
	x := update(1, 0)

	work := func(b int) {
		for i := 1; i <= N; i++ {
			y := update(i, 1<<b)
			if y == 0 {
				return
			}
			if y != x {
				y = update(i, y)
				if y == 0 {
					return
				}
				x = y
			}
		}
	}
	for i := 0; i < 30 && x > 0; i++ {
		work(i)
		if i == 0 {
			work(i)
		}
	}
}
