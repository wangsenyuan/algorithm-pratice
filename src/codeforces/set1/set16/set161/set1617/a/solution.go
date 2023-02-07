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
		S, _ := reader.ReadString('\n')
		T, _ := reader.ReadString('\n')
		res := solve(normalize(S), normalize(T))
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func solve(S, T string) string {
	cnt := make([]int, 26)
	for i := 0; i < len(S); i++ {
		cnt[int(S[i]-'a')]++
	}
	var buf bytes.Buffer
	// only a, b, c matters
	var x int
	if cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0 {
		//abc => acb
		if T == "abc" {
			for cnt[0] > 0 {
				buf.WriteByte('a')
				cnt[0]--
			}
			for cnt[2] > 0 {
				buf.WriteByte('c')
				cnt[2]--
			}
			for cnt[1] > 0 {
				buf.WriteByte('b')
				cnt[1]--
			}
			x = 3
		}
	}

	for x < 26 {
		for cnt[x] > 0 {
			buf.WriteByte(byte(x + 'a'))
			cnt[x]--
		}
		x++
	}

	return buf.String()
}
