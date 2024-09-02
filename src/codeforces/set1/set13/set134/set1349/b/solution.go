package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k)
		if res {
			buf.WriteString("yes\n")
		} else {
			buf.WriteString("no\n")
		}
	}

	fmt.Print(buf.String())
}

func solve(a []int, k int) bool {
	// k 必须是a中的一个数字，否则肯定是不行的
	// k可以把所有临近它的，比它大的数，都可以变成k
	// 只有当k的个数 <= 比它小的数时，才不行
	// 还有一种可能性，就是在某个区间内，k是多数派
	n := len(a)
	if n == 1 {
		return a[0] == k
	}

	// 还有一种情况就是  k-1, k+1, k+1, k-1, k-1, k
	// 那么此时可以先把前面的k-1都变成k+1, 然后再和k进行比较
	ok := false
	ok2 := false
	for i := 0; i < n; i++ {
		if a[i] == k {
			ok = true
		}
		if a[i] >= k && i+1 < n && a[i+1] >= k {
			ok2 = true
		}
		if a[i] >= k && i+2 < n && a[i+2] >= k {
			ok2 = true
		}
	}

	return ok && ok2
}
