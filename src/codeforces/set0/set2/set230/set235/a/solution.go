package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
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

func solve(n int) int64 {
	// x % a = 0
	// x % b = 0
	// x % c = 0
	// 难道不是 n * (n - 1) * (n - 2)
	if n <= 2 {
		return int64(n)
	}
	if n == 3 {
		return 6
	}
	// 6,5,4 => 30 * 2 = 60
	// 5, 4, 3 => 60
	// 会不会存在某个 n, n - 1, x 能获得最大值?
	// 12, 11, 10 => 60 * 11 = 660
	// 12, 11, 7 = 84 * 11 = 更大
	// 那有没有n, x, y的情况呢？
	// lcm(n, x, y) = lcm(lcm(n, x), y), 在y确定的情况下， lcm(n, x) 越大越好吗？
	// 也不成立的，否则固定n, 那么就是需要 x = n - 1, y = n - 2, 上面已经有反例了

	if n&1 == 1 {
		return int64(n) * int64(n-1) * int64(n-2)
	}
	res := int64(n-1) * int64(n-2) * int64(n-3)

	for i := max(1, n-50); i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			for k := j + 1; k <= n; k++ {
				res = max(res, lcm(i, j, k))
			}
		}
	}

	return res
}

func lcm(a, b, c int) int64 {
	res := int64(a) * int64(b) / int64(gcd(a, b))
	res = res * int64(c) / gcd(res, int64(c))
	return res
}

type Num interface {
	int | int64
}

func gcd[T Num](a, b T) T {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max[T Num](a, b T) T {
	if a >= b {
		return a
	}
	return b
}
