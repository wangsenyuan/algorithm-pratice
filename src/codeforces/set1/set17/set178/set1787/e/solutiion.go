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
		n, k, x := readThreeNums(reader)
		res := solve(n, k, x)

		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d", len(cur)))
			for _, num := range cur {
				buf.WriteString(fmt.Sprintf(" %d", num))
			}
			buf.WriteByte('\n')
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

func solve(n int, k int, x int) [][]int {
	var y int

	vis := make([]int, n+1)
	gid := 1
	for i := 1; i <= n; i++ {
		y ^= i
		if vis[i] > 0 {
			continue
		}
		j := x ^ i
		if i == j || j == 0 || j > n || vis[j] > 0 {
			continue
		}
		vis[i] = gid
		vis[j] = gid
		gid++
	}

	if k&1 == 0 && y != 0 {
		return nil
	}
	if k&1 == 1 && y != x {
		return nil
	}

	if x <= n {
		vis[x] = gid
		gid++
	}

	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			vis[i] = gid
		}
	}

	gid--
	if gid < k {
		return nil
	}
	// 每次要减少两个
	if (gid-k)%2 != 0 {
		return nil
	}
	res := make([][]int, k)
	for i := 1; i <= n; i++ {
		if vis[i] >= k {
			vis[i] = k
		}
		j := vis[i] - 1
		if res[j] == nil {
			res[j] = make([]int, 0, 1)
		}
		res[j] = append(res[j], i)
	}

	return res
}
