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
		s := readString(reader)
		res := solve(s)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(s string) bool {
	// + 入栈
	// - 出栈
	// 1 表示栈是一个连续的非递减区间
	// 0 表示栈不是一个连续的非递减区间
	// 假设到高度i时为1， 但是到高度i+1为0， 那么必然有 stack[i+1] < stack[i]
	// 如果存在0的情况，在未出栈前，不能再有1，必须全是0
	// 所以，只需要记录最高的1的位置
	n := len(s)
	var top int
	// 第一次出现0的位置
	p0 := -1
	// 最后一次出现1的位置
	p1 := -1
	//
	for i := 0; i < n; i++ {
		if s[i] == '+' {
			top++
		} else if s[i] == '-' {
			top--
			if p1 == top {
				// 当前是1的位置，需要递减
				p1--
			}
			if p0 == top {
				// 没有0了
				p0 = -1
			}
		} else if s[i] == '1' {
			if p0 >= 0 {
				// 已经存在0了
				return false
			}
			// 到目前为止都是1，非递减序列
			p1 = top - 1
		} else {
			// p0
			if p0 < 0 {
				p0 = top - 1
			}
			if p0 <= p1 || p0 == 0 {
				// 只有一个元素时，肯定是非递减的
				return false
			}
		}
	}

	return true
}
