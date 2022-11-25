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
		readNum(reader)
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

func solve(s string) int64 {
	n := len(s)
	// x 是s的 substring
	// x is good
	// 1. x is already sorted
	// 2. x can be swap sorted
	//     假设 ... 出现了  1101 => 1011 => 0111
	//                    1100 => 1010 =>  1001  => 0101 => bad
	//     1可以移动到的位置是和它右边的0有关系 right_most[0] 如果没有1
	//                                    left_most[1] 如果有1
	// 如果出现1..1后面有两个0， 就是bad
	// 两个0前面出现两个1的时候，就是不可以的
	// 考虑以i为结尾
	// 1是单独的部分
	//   如果 s[i] = '1' dp[i] += dp[i-1], 前面的可以直接连接i
	//       s[i] = '0', left[prev_0[i-1]]
	l0 := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			l0[i] = l0[i-1]
		} else {
			l0[i] = -1
		}
		if s[i] == '0' {
			l0[i] = i
		}
	}
	// l2[i] = 前面第二个1所在的位置
	l2 := make([]int, n)

	stack := make([]int, n)
	var front int
	for i := 0; i < n; i++ {
		if front >= 2 {
			l2[i] = stack[front-2]
		} else {
			l2[i] = -1
		}
		if s[i] == '1' {
			stack[front] = i
			front++
		}
	}

	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		if s[i] == '1' {
			dp[i] = dp[i-1] + 1
		} else {
			if s[i-1] == '0' {
				// 00, 只能到移动一个1
				// 1100
				x := l2[i] + 1
				dp[i] += (i - x) + 1
			} else {
				// 10只能把 i-1 (不包括i-1)之前最右边的1移动过来
				x := l0[i-1]
				if x < 0 {
					dp[i] += i + 1
				} else {
					// 只能移动0前面的一个1到 x的后面一个位置
					x = l2[x] + 1
					dp[i] += i - x + 1
				}
			}
		}
	}

	var res int64

	for i := 0; i < n; i++ {
		res += int64(dp[i])
	}

	return res
}
