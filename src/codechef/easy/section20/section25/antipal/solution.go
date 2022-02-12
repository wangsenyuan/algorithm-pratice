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
		n := readNum(reader)
		s := readString(reader)
		res := solve(n, s)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			buf.WriteString(res)
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, s string) string {
	if n&1 == 1 {
		return ""
	}
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
	}
	var x int
	for i := 0; i < 26; i++ {
		if cnt[i] > cnt[x] {
			x = i
		}
	}
	//aabb
	//aabbabbacc
	if cnt[x] > n/2 {
		// aab is good, but aaaab is bad
		return ""
	}

	buf := make([]byte, n)
	var pos int
	for i := 0; i < 26; i++ {
		for cnt[i] > 0 {
			buf[pos] = byte(i + 'a')
			pos++
			cnt[i]--
		}
	}

	// buf := []byte(s[:n])

	// sort.Sort(Bytes(buf))

	reverse(buf[n/2:])

	return string(buf)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Bytes []byte

func (this Bytes) Len() int {
	return len(this)
}

func (this Bytes) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Bytes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
