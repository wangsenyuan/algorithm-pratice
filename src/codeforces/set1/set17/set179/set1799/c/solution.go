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
	n := len(s)

	cnt := make([]int, 26)
	var distinct int
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
		if cnt[int(s[i]-'a')] == 1 {
			distinct++
		}
	}

	// 考虑在位置i (在i之前的，前缀和后缀相同)
	// 如果ans[i] > ans[j], 那么剩余的其他的，可以按照递增排列就好
	// 所以，这里所有 < ans[i] 能够配对的，都应该放在i的前面
	//  也就是说和ans[i]对应的ans[j] < ans[i] 无法完成配对
	buf := make([]byte, n)
	i, j := 0, n-1

	for x := 0; x < 26; x++ {
		if cnt[x] > 0 {
			distinct--
		}
		for cnt[x] >= 2 {
			buf[i] = byte('a' + x)
			buf[j] = byte('a' + x)
			i++
			j--
			cnt[x] -= 2
		}
		if cnt[x] == 1 {
			// check if any y
			if distinct == 1 {
				// another y
				for y := x + 1; y < 26; y++ {
					if cnt[y] > 0 {
						for i < j {
							buf[i] = byte('a' + y)
							i++
							buf[j] = byte('a' + y)
							j--
							cnt[y] -= 2
						}
						buf[i] = byte('a' + x)
						cnt[x]--
						break
					}
				}

				break
			}

			// x is the last one
			buf[j] = byte('a' + x)
			cnt[x]--

			for y := 0; y < 26; y++ {
				for cnt[y] > 0 {
					buf[i] = byte('a' + y)
					i++
					cnt[y]--
				}
			}
		}
	}

	return string(buf)
}
