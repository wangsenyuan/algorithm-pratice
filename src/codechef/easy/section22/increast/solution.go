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
		S, _ := reader.ReadString('\n')
		res := solve(len(S)-1, S)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

func solve(n int, S string) string {
	add := make([]bool, n)
	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && S[stack[p-1]] > S[i] {
			add[stack[p-1]] = false
			p--
		}
		add[i] = true
		stack[p] = i
		p++
	}

	var buf1, buf2 bytes.Buffer

	for i := 0; i < n; i++ {
		if add[i] {
			buf1.WriteByte(S[i])
		} else {
			buf2.WriteByte(S[i])
		}
	}
	st := buf2.Bytes()
	bef := true
	for i := 1; i < len(st); i++ {
		if st[i] != st[i-1] {
			bef = st[i] < st[i-1]
			break
		}
	}

	ans := buf1.Bytes()
	ind := len(ans)
	if len(st) > 0 && st[0] <= ans[len(ans)-1] {
		for i := 0; i < len(ans); i++ {
			if bef && ans[i] >= st[0] {
				ind = i
				break
			}
			if !bef && ans[i] > st[0] {
				ind = i
				break
			}
		}
	}

	var res bytes.Buffer
	var cur int
	for i := 0; i < n; i++ {
		if add[i] {
			if cur < ind {
				res.WriteByte(S[i])
				cur++
			} else {
				add[i] = false
			}
		}
	}

	for i := 0; i < n; i++ {
		if !add[i] {
			res.WriteByte(S[i])
		}
	}

	return res.String()
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
