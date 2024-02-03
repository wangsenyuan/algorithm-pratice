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

const MAX_A = 1000010

var pos [MAX_A][]int
var ind [MAX_A]int
var divs [MAX_A][]int

func init() {
	for i := 0; i < MAX_A; i++ {
		pos[i] = make([]int, 0, 1)
	}

	for i := 1; i < MAX_A; i++ {
		divs[i] = make([]int, 0, 1)
	}

	for i := 1; i < MAX_A; i++ {
		for j := 2 * i; j < MAX_A; j += i {
			divs[j] = append(divs[j], i)
		}
	}
}
func solve(a []int) int {
	for i := range a {
		pos[a[i]] = append(pos[a[i]], i)
	}

	n := len(a)

	leftHi := make([]int, n)  // >= a[i]
	rightHi := make([]int, n) // > a[i]
	leftLo := make([]int, n)  // < a[i]
	rightLo := make([]int, n) // < a[i]
	s := []int{-1}            // 哨兵
	t := []int{-1}
	for i, v := range a {
		for len(s) > 1 && v > a[s[len(s)-1]] {
			rightHi[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		leftHi[i] = s[len(s)-1]
		s = append(s, i)

		for len(t) > 1 && v <= a[t[len(t)-1]] {
			t = t[:len(t)-1]
		}
		leftLo[i] = t[len(t)-1]
		t = append(t, i)
	}
	for _, i := range s[1:] {
		rightHi[i] = n
	}

	t = append(t[:0], n)
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for len(t) > 1 && v <= a[t[len(t)-1]] {
			t = t[:len(t)-1]
		}
		rightLo[i] = t[len(t)-1]
		t = append(t, i)
	}

	ans := 0
	for i, v := range a { // 最大值为 v
		r := rightHi[i]
		ans += min(rightLo[i], r) - i // 最小值为 v
		for _, d := range divs[v] {   // 最小值为 d
			ps := pos[d]
			l := leftHi[i]
			if len(ps) > 0 && ps[0] < i {
				j := ps[0] // v 左侧最近 d 的下标
				ps = ps[1:]
				if j > l && rightLo[j] > i {
					ans += (j - max(leftLo[j], l)) * (min(rightLo[j], r) - i)
				}
				l = max(l, j) // 避免重复统计
			}
			if len(ps) > 0 {
				j := ps[0] // v 右侧最近 d 的下标
				if j < r && leftLo[j] < i {
					ans += (i - max(leftLo[j], l)) * (min(rightLo[j], r) - j)
				}
			}
		}
		// v 左侧每个数只保留其最右的下标
		if len(pos[v]) > 1 && pos[v][1] == i {
			pos[v] = pos[v][1:]
		}
	}

	for _, x := range a {
		if len(pos[x]) > 0 {
			pos[x] = pos[x][:0]
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
