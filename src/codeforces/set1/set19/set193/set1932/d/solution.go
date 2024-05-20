package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		readNum(reader)
		trump := readString(reader)
		s := readString(reader)
		cards := strings.Split(s, " ")
		res := solve(trump, cards)
		if len(res) == 0 {
			buf.WriteString("IMPOSSIBLE\n")
		} else {
			for i := 0; i < len(res); i += 2 {
				buf.WriteString(fmt.Sprintf("%s %s\n", res[i], res[i+1]))
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(trump string, cards []string) []string {
	// n := len(cards)

	ranks := make([][]int, 4)

	suite := "CDHS"

	id := func(x byte) int {
		for i := 0; i < 4; i++ {
			if suite[i] == x {
				return i
			}
		}
		return -1
	}

	for _, cur := range cards {
		x := id(cur[1])
		y := int(cur[0] - '2')
		ranks[x] = append(ranks[x], y)
	}

	x := id(trump[0])
	for i := 0; i < 4; i++ {
		sort.Ints(ranks[i])
	}

	rep := func(x int, num int) string {
		return fmt.Sprintf("%d%c", num+2, suite[x])
	}

	var res []string
	var pos int
	for i := 0; i < 4; i++ {
		if i == x {
			continue
		}
		for j := 0; j+1 < len(ranks[i]); j += 2 {
			res = append(res, rep(i, ranks[i][j]))
			res = append(res, rep(i, ranks[i][j+1]))
		}
		if len(ranks[i])%2 == 1 {
			if pos == len(ranks[x]) {
				return nil
			}
			res = append(res, rep(i, ranks[i][len(ranks[i])-1]))
			res = append(res, rep(x, ranks[x][pos]))
			pos++
		}
	}

	for pos < len(ranks[x]) {
		res = append(res, rep(x, ranks[x][pos]))
		res = append(res, rep(x, ranks[x][pos+1]))
		pos += 2
	}

	return res
}
