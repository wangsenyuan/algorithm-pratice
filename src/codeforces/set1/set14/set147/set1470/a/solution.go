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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		K := readNNums(reader, n)
		C := readNNums(reader, m)
		res := solve(K, C)
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

func solve(K []int, C []int) int64 {
	n := len(K)
	//m := len(C)

	sort.Ints(K)

	// 前x个礼物，分配给后x个朋友，满足1 <= K[n-x + 1], 2 <= K[n-x+2].... x <= K[n]
	// 这个条件要怎么快速的检查出来呢？
	// 假设对于x，我们已经确定它是满足的
	// 对于 1, 相当于 1 <= K[n] => K[n] - 1 >= 0
	// 对于 2， 相当于 1 <= K[n-1], 2 <= K[n] => K[n] - 2 >= 0, K[n-1] - 1 >= 0
	// 对于x      => K[n] - x >= 0, ... K[n-x + 1] - 1 >= 0
	// 对于1记录 z = K[n] - 1
	// 对于2记录 z = min(z - 1, K[n-1] - 1)
	// 那么当 z < 0的时候，x无法继续

	// 剩下n-x个朋友，直接给钱
	// 全部直接给钱

	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(C[K[i]-1])
	}

	best := sum
	z := 1 << 30
	// 给礼物
	prev := int64(C[0])
	for x := 1; x <= n; x++ {
		// 后x给礼物， 前 n - x个给钱
		sum -= int64(C[K[n-x]-1])
		z = min(z-1, K[n-x]-1)
		if z < 0 {
			break
		}

		tmp := sum + prev
		if tmp < best {
			best = tmp
		}
		if x == len(C) {
			break
		}
		prev += int64(C[x])
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
