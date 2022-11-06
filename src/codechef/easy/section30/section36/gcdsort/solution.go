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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const H = 20

func solve(A []int) bool {

	n := len(A)

	for i := n - 2; i >= 0; i-- {
		if A[i] <= A[i+1] {
			continue
		}
		var g int
		for j := i; j >= 0; j-- {
			g = gcd(A[j], g)
			if g <= A[i+1] {
				for k := j + 1; k <= i; k++ {
					A[k] = gcd(A[k], A[k-1])
				}
				break
			}
		}
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func solve1(A []int) bool {
	// let A[i] = gcd(j, i) <= A[i+1]
	// find the larget j, how?
	// use binary lift
	n := len(A)
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, H)
		g[i][0] = A[i]
	}

	for h := 1; h < H; h++ {
		for i := 0; i < n; i++ {
			if i+(1<<(h-1)) < n {
				g[i][h] = gcd(g[i][h-1], g[i+(1<<(h-1))][h-1])
			}
		}
	}

	// gcd(arr[l....r])
	fn := func(l int, r int) int {
		m := (r - l + 1)
		var res int
		for i := H - 1; i >= 0; i-- {
			if (m>>i)&1 == 1 {
				res = gcd(res, g[l][i])
				l += (1 << i)
				m -= (1 << i)
			}
		}
		return res
	}

	for i := n - 2; i >= 0; i-- {
		if A[i] <= A[i+1] {
			continue
		}
		// A[i] > A[i+1]
		j := search(i, func(j int) bool {
			return fn(j, i) > A[i+1]
		})
		if j == 0 {
			return false
		}
		A[i] = fn(j-1, i)
		// gcd(j-1, i) <= A[i+1]
	}
	return true
}
func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
