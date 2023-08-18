package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
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

func solve(s string, k int) int {
	// 我们考虑f(s) = sum_of(di)
	// 任何一个1（除了最后一个1）， 它的共享都是 11
	// 所以只需要移动最优一个1到最后一位即可
	// 如果1在 1...n-2这些位置，它们的贡献是11
	// 如果在0，它的贡献是10
	// 如果在位置n-1, 它的贡献是1
	n := len(s)

	first := -1
	last := -1
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			if first < 0 {
				first = i
			}
			last = i
		}
	}

	if first < 0 {
		return 0
	}

	buf := []byte(s)

	if first < last {
		if last < n-1 && n-last-1 <= k {
			buf[last], buf[n-1] = buf[n-1], buf[last]
			k -= n - last - 1
		}

		if k > 0 && first > 0 && first <= k {
			buf[0], buf[first] = buf[first], buf[0]
		}
	} else if last < n-1 {
		if n-last-1 <= k {
			buf[last], buf[n-1] = buf[n-1], buf[last]
		} else if first > 0 && first <= k {
			buf[0], buf[first] = buf[first], buf[0]
		}
	}

	var ans int
	for i := 1; i < n-1; i++ {
		if buf[i] == '1' {
			ans += 11
		}
	}
	if buf[0] == '1' {
		ans += 10
	}
	if buf[n-1] == '1' {
		ans++
	}
	return ans
}
