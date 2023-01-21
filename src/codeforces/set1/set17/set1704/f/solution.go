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
		readNum(reader)
		s := readString(reader)
		res := solve(s)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

const ALICE = "Alice"
const BOB = "Bob"

const N = 500005

var F [N]int

func init() {
	vis := make([]bool, 1001)
	for i := 1; i <= 1000; i++ {
		for j := 0; j <= i-2; j++ {
			vis[F[j]^F[i-2-j]] = true
		}

		var j int
		for vis[j] {
			j++
		}
		F[i] = j
		for j := 0; j <= i; j++ {
			vis[j] = false
		}
	}

	for i := 1001; i < N; i++ {
		F[i] = F[i-34]
	}
}

func solve(s string) string {
	n := len(s)
	var red int

	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			red++
		}
	}
	blue := n - red

	if red > blue {
		return ALICE
	}
	if red < blue {
		return BOB
	}

	var ans int

	for i := 1; i <= n; {
		j := i + 1
		for j <= n && s[j-1] != s[j-2] {
			j++
		}
		ans ^= F[j-i]
		i = j
	}

	if ans > 0 {
		return ALICE
	}
	return BOB
}
