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
		res := solve(n)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for _, row := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", row[0], row[1]))
			}
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

func solve(n int) [][]int {
	// 1 + 2 + .. + 2 * n
	// = (1 + 2 * n) * 2 * n / 2 = (1 + 2 * n) * n
	// n = 3, 3 * 7 = 21
	// s1 + s2 + .. + sn = n * (s1 + sn) / 2
	// sn = s1 + (n - 1)
	// n * (s1 + s1 + n - 1) / 2 = (1 + 2 * n) * n
	//
	//  2 * s1 + n - 1 = 2 * (1 + 2 * n)
	//  2 * s1 = 2 * (1 + 2 * n) + 1 - n
	tmp := 2*(1+2*n) + 1 - n
	if tmp%2 != 0 {
		return nil
	}
	// n = 3 => 2 * 7 + 1 - 3 = 12 => s1 = 6
	// a + b = s1
	// a + 2 + b - 1 = s2
	// a + 4 + b - 2 = s3
	// a + 6 + b - 3 = s4
	// ...
	var ans [][]int
	m := (n - 1) / 2
	// 3 * m + 3 = s1
	s1 := tmp / 2
	a := m + 1
	b := 1
	for b <= m {
		ans = append(ans, []int{a, s1 - a})
		a++
		s1++
		ans = append(ans, []int{b, s1 - b})
		b++
		s1++
	}
	ans = append(ans, []int{2*m + 1, 3*m + 2})
	return ans
}
