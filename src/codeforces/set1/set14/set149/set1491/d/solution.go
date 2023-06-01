package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	Q := make([][]int, n)
	for i := 0; i < n; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(Q)
	var buf bytes.Buffer

	for _, ans := range res {
		if ans {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(Q [][]int) []bool {
	// u => u + v when u & v = v
	// a, b
	// 假设 a & b 的某一位有 a[0] = 0, but b[0] = 1, answer is no
	// 因为没有任何数字可以在建低位的0变成1 （但是可以建1变成0)
	// a[0] = 1, but b[0] = 0

	var check func(a, b int) bool

	check = func(a, b int) bool {
		if a > b {
			return false
		}
		var z int

		for i := 0; i < 30; i++ {
			z += (a >> i) & 1
			z -= (b >> i) & 1
			if z < 0 {
				return false
			}
		}
		return true
	}

	ans := make([]bool, len(Q))

	for i, cur := range Q {
		ans[i] = check(cur[0], cur[1])
	}

	return ans
}
