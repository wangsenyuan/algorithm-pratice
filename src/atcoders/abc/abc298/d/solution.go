package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	cmds := make([]string, n)
	for i := 0; i < n; i++ {
		cmds[i] = readString(reader)
	}
	res := solve(cmds)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func solve(cmds []string) []int {
	n := len(cmds)

	base := make([]int, n+1)
	base[0] = 1
	for i := 1; i <= n; i++ {
		base[i] = mul(base[i-1], 10)
	}
	nums := make([]int, n)
	var front, end int
	nums[end] = 1
	end++
	cur := 1

	addNum := func(s string) {
		x := int(s[0] - '0')
		cur = add(mul(cur, 10), x)
		nums[end] = x
		end++
	}

	removeHead := func() {
		ln := end - front
		x := nums[front]
		front++
		cur = sub(cur, mul(x, base[ln-1]))
	}

	var ans []int

	for _, cmd := range cmds {
		if cmd[0] == '1' {
			addNum(cmd[2:])
		} else if cmd[0] == '2' {
			removeHead()
		} else {
			ans = append(ans, cur)
		}
	}

	return ans
}
