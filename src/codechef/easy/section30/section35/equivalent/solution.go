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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		A, B := readTwoNums(reader)
		res := solve(A, B)
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

type Pair struct {
	first  int
	second int
}

func factors(num int) []Pair {
	var res []Pair

	for i := 2; i <= num/i; i++ {
		if num%i != 0 {
			continue
		}
		var cnt int
		for num%i == 0 {
			cnt++
			num /= i
		}
		res = append(res, Pair{i, cnt})
	}
	if num > 1 {
		res = append(res, Pair{num, 1})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].first < res[j].first
	})
	return res
}

func solve(A, B int) bool {
	fa := factors(A)
	fb := factors(B)
	// A ** x = a ** (p * x)
	// B ** y = b ** (q * y)
	// p * x == q * y
	//
	if len(fa) != len(fb) {
		return false
	}

	n := len(fa)

	xy := make([]Pair, n)

	for i := 0; i < n; i++ {
		if fa[i].first != fb[i].first {
			return false
		}
		g := gcd(fa[i].second, fb[i].second)
		xy[i] = Pair{fa[i].second / g, fb[i].second / g}
	}

	for i := 0; i < n; i++ {
		if xy[i] != xy[0] {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
