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
		p := readNNums(reader, n)
		res := solve(p)
		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(a []int) []int {
	// expect a
	// 对于结尾的000来说，它们不需要进行翻转，可以在任何时间点进行（倒序插入即可）
	// 第一个1，它只能被翻转（奇数）次，接下来的0（经过了连续个1后）,翻转偶数次
	// 也就是对于a[i]来说，可以从后往前判断它们需要翻转多少次，最多的应该要首先放入进去
	n := len(a)

	if a[n-1] == 1 {
		return nil
	}

	// we do have a solution
	var res []int

	// 如果有一段连续的0,如果顺序插入
	// 0 => 10 => 010 => 1010
	// 所以此时，只能倒着插入
	// 如果是连续的1呢？
	// 似乎也只能倒着插入，然后在最后进行翻转
	// 0 => 00 => 000 => 1
	for i := n - 1; i >= 0; {
		// a[i]= 0
		for i > 0 && a[i-1] == 0 {
			res = append(res, 0)
			i--
		}
		if i == 0 {
			res = append(res, 0)
			break
		}
		// a[i-1] = 1
		j := i
		for i > 0 && a[i-1] == 1 {
			res = append(res, 0)
			i--
		}
		res = append(res, j-i)
		i--
	}

	return res
}
