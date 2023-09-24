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
		queries := make([]string, n)
		for i := 0; i < n; i++ {
			queries[i] = readString(reader)
		}
		res := solve(queries)
		for _, ok := range res {
			if ok {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
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

func solve(queries []string) []bool {
	// 如果一旦t有了非a的字符，那么答案就都是yes
	// 如果s有非a的字符，那么答案都是no
	// 然后就是计算两者的长度
	type Item struct {
		ln        int
		has_non_a bool
	}
	players := []*Item{
		&Item{1, false},
		&Item{1, false},
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		if players[1].has_non_a {
			ans[i] = true
			continue
		}
		buf := []byte(cur)
		id := int(cur[0] - '1')
		var k int
		pos := readInt(buf, 2, &k)
		x := cur[pos+1:]
		var has_non_a bool
		for j := 0; j < len(x); j++ {
			if x[j] != 'a' {
				has_non_a = true
				break
			}
		}
		if has_non_a {
			// 如果有非a的了
			players[id].has_non_a = true
		} else {
			players[id].ln += k * len(x)
		}

		if players[1].has_non_a {
			ans[i] = true
			continue
		}
		if players[0].has_non_a {
			ans[i] = false
			continue
		}
		ans[i] = players[0].ln < players[1].ln
	}

	return ans
}
