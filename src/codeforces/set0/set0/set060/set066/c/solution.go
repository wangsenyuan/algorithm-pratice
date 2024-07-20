package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		line, err := readString(reader)
		if err != nil || len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	res := solve(lines)
	fmt.Printf("%d %d\n", res[0], res[1])
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

func readString(reader *bufio.Reader) (string, error) {
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i], nil
		}
	}
	return s, nil
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

func solve(lines []string) []int {
	sort.Strings(lines)

	res := []int{0, 0}
	for i := 0; i < len(lines); {
		j := i
		p := next(lines[j], 0)
		p = next(lines[j], p)
		var cur []string

		for i < len(lines) {
			q := next(lines[i], 0)
			q = next(lines[i], q)
			if p != q || lines[j][:p] != lines[i][:q] {
				break
			}
			cur = append(cur, lines[i][p+1:])
			i++
		}

		tmp := solveFolder(cur)
		res[0] = max(res[0], tmp[0])
		res[1] = max(res[1], tmp[1])
	}

	return res
}

func solveFolder(lines []string) []int {
	dirs := make(map[string]int)

	for _, line := range lines {

		for i := 0; i < len(line); {
			j := next(line, i)
			if j < len(line) {
				dirs[line[:j]]++
			}
			i = j
		}
	}

	return []int{len(dirs), len(lines)}
}

func next(s string, cur int) int {
	for i := cur + 1; i < len(s); i++ {
		if s[i] == '\\' {
			return i
		}
	}
	return len(s)
}
