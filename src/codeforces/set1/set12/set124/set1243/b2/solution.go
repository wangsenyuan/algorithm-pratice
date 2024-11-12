package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		readNum(reader)
		s := readString(reader)
		t := readString(reader)
		ok, res := solve(s, t)
		if !ok {
			buf.WriteString("No\n")
		} else {
			buf.WriteString(fmt.Sprintf("Yes\n%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
			}
		}
	}

	fmt.Print(buf.String())
}
func solve(s string, t string) (bool, [][]int) {
	n := len(s)

	a := []byte(s)
	b := []byte(t)

	get := func(x byte) int {
		return int(x - 'a')
	}

	cnt := make([]int, 26)
	for i := range n {
		cnt[get(a[i])]++
		cnt[get(b[i])]++
	}

	for _, v := range cnt {
		if v%2 == 1 {
			return false, nil
		}
	}
	var res [][]int

	for i := range n {
		if a[i] == b[i] {
			continue
		}
		j := i + 1
		for j < n && a[i] != a[j] {
			j++
		}
		if j < n {
			res = append(res, []int{j + 1, i + 1})
			a[j], b[i] = b[i], a[j]
			continue
		}
		j = i + 1
		for j < n && a[i] != b[j] {
			j++
		}
		// j < n must hold
		res = append(res, []int{j + 1, j + 1})
		a[j], b[j] = b[j], a[j]
		res = append(res, []int{j + 1, i + 1})
		a[j], b[i] = b[i], a[j]
	}

	return true, res
}
