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

	instructions := make([]string, n)
	for i := 0; i < n; i++ {
		instructions[i] = readString(reader)
	}

	res := solve(instructions)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(cur)
		buf.WriteByte('\n')
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

const H = 10

func solve(instructions []string) []string {
	args := []int{0, 0, 0}

	n := len(instructions)

	params := make([]int, n)

	for i := 0; i < n; i++ {
		readInt([]byte(instructions[i]), 2, &params[i])
	}

	run := func(d int, x int) int {
		for i, cur := range params {
			y := (cur >> d) & 1
			if instructions[i][0] == '|' {
				x |= y
			} else if instructions[i][0] == '&' {
				x &= y
			} else {
				x ^= y
			}
		}

		return x
	}

	for i := 0; i < H; i++ {
		x := run(i, 0)
		y := run(i, 1)
		if x == 0 && y == 0 {
			// 始终为0, 那么就是 and 0
		} else if x == 1 && y == 1 {
			// 始终为1, or 1, 保证始终为1， and 1, 保证第二步的结果也是1
			args[0] |= 1 << i
			args[1] |= 1 << i
		} else if x == 0 && y == 1 {
			// 和原来的值相同, or 0, and 1,
			args[1] |= 1 << i
		} else {
			// 和原来的值相反
			args[1] |= 1 << i
			args[2] |= 1 << i
		}
	}

	res := make([]string, 3)
	res[0] = fmt.Sprintf("| %d", args[0])
	res[1] = fmt.Sprintf("& %d", args[1])
	res[2] = fmt.Sprintf("^ %d", args[2])
	return res
}
