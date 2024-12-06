package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readString(reader)
	c := readString(reader)
	a := readString(reader)
	res := solve(c, a)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func solve(c string, a string) []int {
	n := len(c)
	// 0 for none
	// 1 for clown only
	// 2 for acrobat only
	// 3 for both
	roles := make([][]int, 4)

	for i := range n {
		x := int(c[i] - '0')
		y := int(a[i] - '0')
		v := (x << 1) | y
		roles[v] = append(roles[v], i+1)
	}

	for i := 0; i <= len(roles[0]); i++ {
		if i > n/2 {
			break
		}
		v := len(roles[1]) + len(roles[3]) - n/2 + i

		if v < 0 {
			continue
		}
		if v > len(roles[3]) {
			break
		}

		for j := 0; j <= len(roles[1]); j++ {
			// i + j + u + v == n / 2
			// j + u + v = n / 2 - i
			u := n/2 - i - v - j
			if u >= 0 && u <= len(roles[2]) {
				var ans []int
				ans = append(ans, roles[0][:i]...)
				ans = append(ans, roles[1][:j]...)
				ans = append(ans, roles[2][:u]...)
				ans = append(ans, roles[3][:v]...)
				return ans
			}
		}
	}

	return nil
}
