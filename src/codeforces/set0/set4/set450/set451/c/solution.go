package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		nums := readNNums(reader, 4)
		res := solve(nums[0], nums[1], nums[2], nums[3])
		if res {
			buf.WriteString("yes\n")
		} else {
			buf.WriteString("no\n")
		}
	}

	fmt.Print(buf.String())
}

func solve(n int, k int, d1, d2 int) bool {
	if n%3 != 0 {
		return false
	}

	// x <= y  (2种可能性)
	// z <= y (2种可能性)
	return solve1(n, k, d1, d2) || solve2(n, k, d1, d2) || solve3(n, k, d1, d2) || solve4(n, k, d1, d2)
}

func solve1(n int, k int, d1 int, d2 int) bool {
	// y - x = d1 => y = d1 + x
	// y - z = d2 => y = d2 + z
	// y + x + z = k => y = k - (x + z)
	// 3 * y = k + d1 + d2
	y := k + d1 + d2
	if y%3 != 0 {
		return false
	}
	y /= 3
	return y <= n/3 && y-d1 >= 0 && y-d2 >= 0
}

func solve2(n int, k int, d1 int, d2 int) bool {
	// y - x = d1 => y = d1 + x
	// z - y = d2 => y = z - d2
	y := k - d2 + d1
	if y < 0 || y%3 != 0 {
		return false
	}
	y /= 3
	return y <= n/3 && y >= d1 && y+d2 <= n/3
}

func solve3(n int, k int, d1 int, d2 int) bool {
	// x - y = d1 => y = x - d1
	// y - z = d2 => y = z + d2
	y := k - d1 + d2
	if y < 0 || y%3 != 0 {
		return false
	}
	y /= 3
	return y <= n/3 && y >= d2 && y+d1 <= n/3
}

func solve4(n int, k int, d1 int, d2 int) bool {
	// x - y = d1 => y = x - d1
	// z - y = d2 => y = z - d2
	y := k - d1 - d2
	if y < 0 || y%3 != 0 {
		return false
	}
	y /= 3
	return y <= n/3 && y+d1 <= n/3 && y+d2 <= n/3
}
