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
		n, k := readTwoNums(reader)
		res := solve(n, k)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for _, p := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", p[0], p[1]))
			}
		}
	}

	fmt.Print(buf.String())
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

const H = 17

func solve(n int, k int) [][]int {
	if k == 0 {
		var res [][]int
		for i := 0; i < n/2; i++ {
			res = append(res, []int{i, i ^ (n - 1)})
		}
		return res
	}

	if k < n-1 {
		ck := k ^ (n - 1)
		var res [][]int
		res = append(res, []int{0, ck})
		res = append(res, []int{k, n - 1})
		for i := 0; i < n/2; i++ {
			ci := i ^ (n - 1)
			if i == 0 || i == k || i == ck || ci == ck || ci == k || ci == 0 {
				continue
			}
			res = append(res, []int{i, ci})
		}
		return res
	}
	// k = n - 1
	if n == 4 {
		return nil
	}
	var res [][]int
	res = append(res, []int{n - 1, n - 2})
	res = append(res, []int{n - 3, 1})
	res = append(res, []int{0, 2})
	for i := 3; i < n/2; i++ {
		ci := i ^ (n - 1)
		if ci == n-1 || ci == n-2 || ci == n-3 || ci == 0 {
			continue
		}
		res = append(res, []int{i, ci})
	}
	return res
}
