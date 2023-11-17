package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(a []int) int {
	sort.Ints(a)
	// 假设i参与，必须至少有a[i] + 1（自己)参与
	// 如果 a[i] = 0,i肯定会参与
	// 如果i参与了，i-1可以不参与吗？
	// 如果 a[i-1] < a[i]， 肯定时不行的
	// 因为i参与了，意味这至少有a[i]个, 那么符合i-1的条件， i-1也必须参与
	// 如果 a[i-1] == a[i] 也时如此
	// 但是如果 a[i] 不参加，那么 a[i+1] 肯定也不参加
	// 6 0 3 3 6 7 2 7
	// 0,2, 3, 3, 6, 6, 7, 7
	// 0
	// 0, 2, 3, 3
	// 0, 2, 3, 3, 6, 6, 7, 7
	var res int
	n := len(a)
	res++

	for i := n - 2; i >= 0; i-- {
		// i个人去参加
		if a[i] >= i+1 {
			// i不能参加
			continue
		}
		// a[i] <= i + 1 i可以参加
		if a[i+1] > i+1 {
			res++
		}
	}
	if a[0] > 0 {
		res++
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
