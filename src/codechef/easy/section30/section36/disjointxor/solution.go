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
		s := readString(reader)
		res := solve(s[:n])
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 1000000007

func solve(s string) int {
	n := len(s)

	var k int

	for l1 := 0; l1+1 < n; l1++ {
		for l2 := l1 + 1; l2 < n; l2++ {
			if s[l1] != s[l2] {
				k = max(k, min(l2-l1, n-l2))
			}
		}
	}

	if k == 0 {
		return 0
	}

	var active []Pair

	for i := 0; i+2*k <= n; i++ {
		if s[i] != s[i+k] {
			active = append(active, Pair{i, i + k})
		}
		if s[i] != s[n-k] {
			active = append(active, Pair{i, n - k})
		}
	}

	buf := make([]byte, k)
	buf[0] = '1'

	for pos := 1; pos < k; pos++ {
		var next []Pair
		for _, cur := range active {
			if s[pos+cur.first] != s[pos+cur.second] {
				next = append(next, cur)
			}
		}
		if len(next) == 0 {
			buf[pos] = '0'
		} else {
			buf[pos] = '1'
			active = next
		}
	}

	var res int64

	for i := 0; i < k; i++ {
		res = res*2 + int64(buf[i]-'0')
		res %= MOD
	}

	return int(res)
}

type Pair struct {
	first  int
	second int
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
