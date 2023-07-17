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
	deck := readNNums(reader, n)
	res := solve(deck)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d ", len(cur)))
		for i := 0; i < len(cur); i++ {
			buf.WriteString(fmt.Sprintf("%d ", cur[i]))
		}
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

func solve(deck []int) [][]int {
	var res [][]int
	n := len(deck)

	tmp := make([]int, n)

	perform := func(op []int) {
		var cur int
		for _, d := range op {
			for i := 0; i < d; i++ {
				tmp[n-cur-d+i] = deck[cur+i]
			}
			cur += d
		}
		copy(deck, tmp)
	}

	for {
		var op []int
		for i := 0; i < n; i++ {
			if deck[i] == i+1 {
				continue
			}
			if i > 0 {
				op = append(op, i)
			}
			for j := i + 1; j < n; j++ {
				if deck[j] == deck[j-1]+1 {
					continue
				}
				op = append(op, j-i)
				for k := j; k < n; k++ {
					if deck[k] != deck[i]-1 {
						continue
					}
					op = append(op, k-j+1)
					if k < n-1 {
						op = append(op, n-k-1)
					}
					break
				}
				break
			}
			break
		}
		if len(op) == 0 {
			break
		}
		res = append(res, op)
		perform(op)
	}

	return res
}
