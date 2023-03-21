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

	ask := func(a, b, c int) int {
		fmt.Printf("? %d %d %d\n", a, b, c)
		return readNum(reader)
	}

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n, ask)
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("! %d", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
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

func solve(n int, ask func(int, int, int) int) []int {
	// a, b, c 中如果 回答是0, more impostors, 1, else
	// a, b, c, d
	// (a, b, c) (b, c, d), (a, c, d)
	//如果都为0， 那么a, b, c, d中最多只能有一个人为 crewmate
	// a, b, c => 0,
	// b, c, d => 0, 那么 b, c 中至少有一个impstor
	// a, b, c => 0
	// b, c, d => 1 ， 那么b, c中只能有一个impostor, 可以确定a是impostor, d是crewmate
	// 如果 进而可以再得到 b, 和 c的角色 (a, b, d), (a, c, d)
	// find a & d first

	get := func(x int) int {
		return x%n + 1
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		// 0, 1, 2, 3
		a := get(i)
		b := get(i + 1)
		c := get(i + 2)
		ans[i] = ask(a, b, c)
	}

	first, second := -1, -1

	for i := 0; i < n; i++ {
		if ans[i] != ans[(i+1)%n] {
			if ans[i] == 0 {
				first = i + 1
				second = get(i + 3)
			} else {
				first = get(i + 3)
				second = i + 1
			}
			break
		}
	}

	var res []int

	for i := 1; i <= n; i++ {
		if i == first {
			res = append(res, i)
			continue
		}
		if i == second {
			continue
		}
		x := ask(first, i, second)
		if x == 0 {
			res = append(res, i)
		}
	}
	return res
}
