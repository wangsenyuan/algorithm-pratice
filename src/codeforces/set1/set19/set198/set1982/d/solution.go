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
		n, m, k := readThreeNums(reader)
		H := make([][]int, n)
		for i := 0; i < n; i++ {
			H[i] = readNNums(reader, m)
		}
		M := make([]string, n)
		for i := 0; i < n; i++ {
			M[i] = readString(reader)
		}
		res := solve(H, M, k)
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

func solve(H [][]int, M []string, k int) bool {
	n := len(H)
	m := len(H[0])

	cnt := make([][]int, n)
	var sum int
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, m)
		for j := 0; j < m; j++ {
			v := H[i][j]
			w := 1
			if M[i][j] == '0' {
				v *= -1
				w = 0
			}
			sum += v
			cnt[i][j] = w
			if i > 0 {
				cnt[i][j] += cnt[i-1][j]
			}
			if j > 0 {
				cnt[i][j] += cnt[i][j-1]
			}
			if i > 0 && j > 0 {
				cnt[i][j] -= cnt[i-1][j-1]
			}
		}
	}

	sum = abs(sum)

	if sum == 0 {
		return true
	}

	var g int
	for i := k - 1; i < n; i++ {
		for j := k - 1; j < m; j++ {
			v := cnt[i][j]
			if i-k >= 0 {
				v -= cnt[i-k][j]
			}
			if j-k >= 0 {
				v -= cnt[i][j-k]
			}
			if i-k >= 0 && j-k >= 0 {
				v += cnt[i-k][j-k]
			}

			v = k*k - 2*v

			g = gcd(g, abs(v))
		}
	}
	if g == 0 {
		return false
	}
	return sum%g == 0
}

func abs(num int) int {
	return max(num, -num)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
