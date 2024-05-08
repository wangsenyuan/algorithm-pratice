package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	if len(res) == 0 {
		fmt.Println("-1")
		return
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d", len(cur)))
		for _, i := range cur {
			buf.WriteString(fmt.Sprintf(" %d", i))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(s string) [][]int {
	// 1结尾的和0结尾的排在两个队列里面
	// a for 0, b for 1
	var a [][]int
	var b [][]int

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			if len(b) == 0 {
				// 新开一个队列
				a = append(a, []int{i + 1})
			} else {
				m := len(b)
				last := b[m-1]
				b = b[:m-1]
				last = append(last, i+1)
				a = append(a, last)
			}
		} else {
			// 1
			if len(a) == 0 {
				// no answer
				return nil
			}
			m := len(a)
			last := a[m-1]
			a = a[:m-1]
			last = append(last, i+1)
			b = append(b, last)
		}
	}

	if len(b) > 0 {
		return nil
	}
	return a
}
