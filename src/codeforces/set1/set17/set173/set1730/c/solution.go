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
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(s string) string {
	// 首先一个d只能变大，所以最好是把它替换到后面去
	// 在第i位的时候，如果存在j i < j 且s[j] < s[i], 那么最优的方案是把j换过来
	// 所以需要快速知道后面的最小的s[j]
	// 不是简单的交换，而是要[i...j)全部加1, 把j换上来
	n := len(s)

	last := make([]int, 11)
	for i := 0; i <= 10; i++ {
		last[i] = -1
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		last[x] = i
	}

	cnt := make([]int, 10)

	buf := make([]byte, n)
	var i int
	var expect int
	for j := 0; j < n; {
		for expect < 10 && last[expect] < i {
			for cnt[expect] > 0 {
				buf[j] = byte('0' + expect)
				j++
				cnt[expect]--
			}
			expect++
		}
		if j == n {
			break
		}

		for i < last[expect] {
			x := int(s[i] - '0')
			if x == expect {
				buf[j] = s[i]
				j++
			} else {
				cnt[min(x+1, 9)]++
			}
			i++
		}

		buf[j] = byte('0' + expect)
		j++
		i = last[expect] + 1
		for cnt[expect] > 0 {
			buf[j] = byte(expect + '0')
			j++
			cnt[expect]--
		}
		expect++
	}

	return string(buf)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
