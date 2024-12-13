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
		a := readNNums(reader, n)
		ok, res := solve(a)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d ", x))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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
func solve(a []int) (bool, []int) {
	n := len(a)

	p := make([]int, n+1)
	copy(p[1:], a)

	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if p[i]&1 != i&1 {
			return false, nil
		}
		pos[p[i]] = i
	}
	var res []int
	reverse := func(l int) {
		res = append(res, l)
		for i, j := 1, l; i < j; i, j = i+1, j-1 {
			p[i], p[j] = p[j], p[i]
		}
		for i := 1; i <= l; i++ {
			pos[p[i]] = i
		}
	}

	for i := n; i > 1; i -= 2 {
		if pos[i] == i && pos[i-1] == i-1 {
			continue
		}
		// else must change
		// 先把i换到位置1
		reverse(pos[i])
		// 然后把i换到i-1的前面
		reverse(pos[i-1] - 1)
		// 调换他们的位置
		reverse(pos[i-1] + 1)
		// 把i放在位置1
		reverse(pos[i])
		reverse(i)
	}

	return true, res
}
