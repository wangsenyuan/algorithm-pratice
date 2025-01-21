package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, b, a := readThreeNums(reader)
	s := readString(reader)
	s = strings.ReplaceAll(s, " ", "")
	return solve(b, a, s[:n])
}

func solve(b int, a int, s string) int {
	n := len(s)

	c := a

	que := make([]int, n)
	var head, tail int
	var res int

	for i := 0; i <= n; i++ {
		res = max(res, i+b+a)

		if a == 0 {
			// 虽然到达了位置i，但是accumlator 没有电了
			if b == 0 {
				// 完全没电了
				break
			}
			for tail < head {
				j := que[tail]
				tail++
				if i-j < c {
					b--
					a = min(a+2, c)
					break
				}
				// 这个地方使用battery没有好处, accumator是满电的，无法充电
				// 比如c = 2, 那么当前位置i = 3, 那么在位置1的时候，它正好是满电的
			}

		}
		if a > 0 {
			a--
			if i < n && s[i] == '1' {
				que[head] = i
				head++
			}
		} else {
			b--
			if i < n && s[i] == '1' {
				// 充电
				a = min(a+1, c)
			}
		}

	}

	return min(res, n)
}
