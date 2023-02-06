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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
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

func solve(commands []string) []int64 {
	var res []int64

	set := make(map[int64]bool)
	set[0] = true

	p := make(map[int64][]int64)
	q := make(map[int64][]int64)

	mex := make(map[int64]int64)

	// 1需要特别处理
	for _, cmd := range commands {
		var x int64
		readInt64([]byte(cmd), 2, &x)
		if cmd[0] == '+' {
			set[x] = true
		} else if cmd[0] == '-' {
			delete(set, x)
			for _, i := range p[x] {
				if len(q[i]) == 0 {
					q[i] = make([]int64, 0, 1)
				}
				q[i] = append(q[i], x)
			}
			delete(p, x)
		} else {
			i := max(mex[x], 1)
			for set[i*x] {
				if len(p[i*x]) == 0 {
					p[i*x] = make([]int64, 0, 1)
				}
				p[i*x] = append(p[i*x], x)
				i++
			}
			mex[x] = i
			now := i * x
			for _, k := range q[x] {
				if !set[k] {
					if now > k {
						now = k
					}
				}
			}
			res = append(res, now)
		}
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
