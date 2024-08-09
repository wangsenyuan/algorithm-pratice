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
		n, m, k := readThreeNums(reader)
		res := solve(n, m, k)
		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		for _, x := range res {
			buf.WriteString(x)
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(n int, m int, k int) []string {
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			buf[i][j] = ' '
		}
	}
	if n%2 == 1 {
		if k < m/2 {
			return nil
		}
		k -= m / 2
		for j := 0; j < m; j += 2 {
			buf[n-1][j] = byte('d' + ((j >> 1) & 1))
			buf[n-1][j+1] = byte('d' + ((j >> 1) & 1))
		}
		n--
	}
	if m%2 == 1 {
		// 垂直的至少要n/2
		if n*m/2-k < n/2 {
			return nil
		}
		for i := 0; i < n; i += 2 {
			buf[i][m-1] = byte('f' + ((i >> 1) & 1))
			buf[i+1][m-1] = byte('f' + ((i >> 1) & 1))
		}
		m--
	}

	if k%2 == 1 {
		return nil
	}

	// n 和m现在都是偶数
	// 然后放置k个水平的， 必须2by2的放
	// a/b 交替放
	for i := 0; i < n && k > 0; i += 2 {
		tmp := "ab"
		for j := 0; j < m && k > 0; j += 2 {
			var x int
			if j > 0 && buf[i][j-1] == tmp[x] {
				x ^= 1
			}
			buf[i][j] = tmp[x]
			buf[i][j+1] = tmp[x]
			buf[i+1][j] = tmp[1^x]
			buf[i+1][j+1] = tmp[1^x]

			k -= 2
		}
	}
	// 其他的都放置垂直的
	for i := n - 2; i >= 0; i -= 2 {
		tmp := "xy"
		for j := 0; j < m; j += 2 {
			if buf[i][j] != ' ' {
				// 这个 2 * 2 放置了水平的
				break
			}
			var x int

			if i+2 < n && buf[i+2][j] == tmp[x] {
				x ^= 1
			}

			buf[i][j] = tmp[x]
			buf[i+1][j] = tmp[x]
			buf[i][j+1] = tmp[1^x]
			buf[i+1][j+1] = tmp[1^x]
		}
	}

	for i := 0; i < n; i += 2 {
		tmp := "uv"
		for j := 0; j < m; j += 2 {
			if buf[i][j] != ' ' {
				continue
			}

			buf[i][j] = tmp[0]
			buf[i+1][j] = tmp[0]
			buf[i][j+1] = tmp[1]
			buf[i+1][j+1] = tmp[1]
		}
	}

	res := make([]string, len(buf))
	for i := 0; i < len(buf); i++ {
		res[i] = string(buf[i])
	}

	return res
}
