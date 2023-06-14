package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	ask := func(tp string, i int, j int) int {
		fmt.Printf("%s %d %d\n", tp, i, j)
		return readNum(reader)
	}

	res := solve(n, ask)

	var buf bytes.Buffer

	buf.WriteString("!")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	buf.WriteByte('\n')
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

func solve(n int, ask func(tp string, i int, j int) int) []int {
	// i & j, i | j, i xor j
	// a and b
	// a xor b
	// a or b

	xor := make([]int, n+1)

	for i := 2; i <= n; i++ {
		xor[i] = ask("XOR", 1, i)
	}

	// 分两种情况，有相同数字的，没有相同数字的
	pos := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if pos[xor[i]] > 0 {
			return solve1(n, xor, pos[xor[i]], i, ask)
		}
		if xor[i] == 0 {
			return solve1(n, xor, 1, i, ask)
		}
		pos[xor[i]] = i
	}
	// 没有相同数字的情况
	// 数字是[1..n]的permutation
	// 肯定存在a ^ b = (n - 1)
	first := 2
	second := pos[(n-1)^xor[2]]
	for second == 0 {
		first++
		second = pos[(n-1)^xor[first]]
	}
	// and(first, second) = 0 or(first, second) = n - 1
	ans := make([]int, n+1)
	ans[1] = ask("AND", 1, first) | ask("AND", 1, second)
	for i := 2; i <= n; i++ {
		ans[i] = xor[i] ^ ans[1]
	}
	return ans[1:]
}

func solve1(n int, xor []int, first int, second int, ask func(tp string, i int, j int) int) []int {
	// ans[first] == ans[second]
	ans := make([]int, n+1)
	ans[first] = ask("AND", first, second)
	//ans[second] = ans[first]
	if first != 1 {
		ans[1] = ans[first] ^ xor[first]
	}
	for i := 2; i <= n; i++ {
		ans[i] = xor[i] ^ ans[1]
	}
	return ans[1:]
}
