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
		arr := readNNums(reader, n)
		res := solve(arr)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

const CHEF = "Chef"
const COOK = "Cook"

func solve(arr []int) string {
	// last one doesn't matter
	// 1, 0, Chef win
	// 2, 0, cook win
	// 1, 2, 0 chef win
	// 2, 2, 0 cook win
	// 假设 从i开始，first win
	// 那么如果 i-1是odd时， 假设为1， 如果 first 把它直接放入最后一位， second 就可以做为后续的先手
	//                   如果first把它放入后面 n - 2 ( >= i) 则还是first win
	//             如果为even，假设为2， 这是因为，如果first迁移了一个到n, 则回到了上面的判断
	//                                         如果迁移到了n-1,
	n := len(arr)
	var res int
	for i := n - 1; i > 0; i-- {
		x := arr[i-1]
		if x&1 == 1 {
			res ^= (n - i)
		}
	}

	if res > 0 {
		return CHEF
	}

	return COOK
}
