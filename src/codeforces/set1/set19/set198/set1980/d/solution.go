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
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

const inf = 1 << 30

func solve(a []int) bool {
	// 如果已经是升序的，答案是肯定的
	// else 假设删除a[i]后，可以满足条件
	// 删除前, b[i-1] = gcd(a[i-1], a[i]), b[i] = gcd(a[i], a[i+1]), b[i+1] = gcd(a[i+1], a[i+2])
	// 可以看到, a[i]影响的只有 b[i-1]和 b[i],
	// 删除后, b[i-1] = gcd(a[i-1], a[i+1]), b[i] = b[i+1] = gcd(a[i+1], a[i+2])
	// b[i-1] 有可能变大，也可能变小，b[i+1]保持不变
	// 但是 b[i-2] <= b[i-1] <= b[i+1] <= b[i+2]
	// 所以，在i-2和i+1的关系不能变，且新的b[i-1] 要处在中间
	// 头尾单独处理
	n := len(a)
	b := make([]int, n)
	for i := 0; i+1 < n; i++ {
		b[i] = gcd(a[i], a[i+1])
	}
	if n <= 3 || sort.IntsAreSorted(b) {
		return true
	}
	b[n-1] = inf
	// suf[i] = 从b[i]开始，是否升序
	suf := make([]bool, n)
	suf[n-1] = true
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] && b[i] <= b[i+1]
	}
	if suf[1] {
		// 删掉a[0]后，直接就可以了， 这个好像没啥意义， b[0] = gcd(a[0], a[1])
		return true
	}
	pref := make([]bool, n)
	pref[0] = true
	for i := 1; i < n-1; i++ {
		// 删除a[i]
		cur := gcd(a[i-1], a[i+1])
		if (i == 1 || b[i-2] <= cur && pref[i-2]) && cur <= b[i+1] && suf[i+1] {
			return true
		}
		pref[i] = pref[i-1] && b[i-1] <= b[i]
	}

	return pref[n-3]
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
