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
		n, k := readTwoNums(reader)
		ok, res := solve(n, k)
		if !ok {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))
			for _, idx := range res {
				buf.WriteString(fmt.Sprintf("%d\n", idx))
			}
		}
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
func solve(n int, k int) (bool, []int) {
	// n is inrelavant
	if k == 0 {
		return true, nil
	}
	if k%2 == 0 {
		return false, nil
	}
	var op int
	for (1 << uint(op+1)) <= k {
		op++
	}
	// 1 << op >= k
	res := make([]int, op+1)
	for i := 0; i <= op; i++ {
		res[i] = -1
	}
	idx := 1
	for op >= 0 {
		if k&(1<<uint(op)) == 0 {
			break
		}
		res[op] = idx
		idx += 1 << uint(op)
		k -= 1 << uint(op)
		op--
	}
	if op != -1 {
		v := (1 << uint(op)) - 1 - k
		x := ((1 << uint(op)) + v) / 2
		res[op] = idx - x
		l := idx - x
		r := idx - x + (1 << uint(op))
		op--
		for ; op >= 0; op-- {
			if ((idx - l) & (1 << uint(op))) > 0 {
				res[op] = l
				l += 1 << uint(op)
			} else {
				res[op] = r
				r += 1 << uint(op)
			}
		}
	}

	return true, res
}
