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
		s := readString(reader)[:n]
		l, ops := solve(s)
		buf.WriteString(fmt.Sprintf("%d %d\n", l, len(ops)))

		for _, op := range ops {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", op[0], op[1], op[2]))
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

func solve(s string) (int, [][]int) {
	cnt := make([]int, 2)

	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'0')]++
	}
	n := len(s)

	if cnt[0] == n || cnt[0] == 0 {
		// all same, no operations can be done
		return n, nil
	}
	if cnt[0] == cnt[1] {
		return 1, [][]int{{1, n, 1}}
	}
	// 假设cnt[0] > cnt[1], 则一次操作可以增加一次cnt[1]
	// 所以 k = cnt[0] - cnt[1] + 1 (最后一次是将所有的变成1)
	var ops [][]int

	buf := make([]byte, n)

	diff := max(cnt[0], cnt[1]) - min(cnt[0], cnt[1])

	var p int
	for i := 0; i < n; i++ {
		buf[p] = s[i]
		p++
		for p >= 2 && buf[p-2] != buf[p-1] && diff > 0 {
			diff--
			p -= 2
			if cnt[1] > cnt[0] {
				ops = append(ops, []int{p + 1, p + 2, 0})
				buf[p] = '0'
				p++
			} else {
				ops = append(ops, []int{p + 1, p + 2, 1})
				buf[p] = '1'
				p++
			}
		}
	}

	ops = append(ops, []int{1, p, 0})

	return 1, ops
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
