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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == 'r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readString(reader)[:n]
		b := readString(reader)[:n]
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(A, B string) int {
	// 当前状态是x, x ^ (1 << i) ^ mask_1
	// 最后要达到y的状态, mask_1如果是偶数次，就可以忽略; 奇数次相当于进行了一次
	// 这里还有一个限制，就是只能操作lit的，所以，如果一个点是未点亮的，必须先点亮后，才能继续处理
	if A == B {
		return 0
	}
	dif := make([]int, 2)
	sam := make([]int, 2)

	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			dif[int(A[i]-'0')]++
		} else {
			sam[int(A[i]-'0')]++
		}
	}
	res := len(A) + 1
	if dif[1] == dif[0] {
		res = min(res, dif[1]+dif[0])
	}
	if sam[1]-sam[0] == 1 {
		res = min(res, sam[1]+sam[0])
	}
	if res <= len(A) {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
