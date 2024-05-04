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
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}

	res := solve(words)

	fmt.Println(res)
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

func solve(words []string) string {
	//n := len(words)
	var m int
	for _, word := range words {
		m += len(word)
	}

	next := make([]int, 26)
	prev := make([]int, 26)
	vis := make([]bool, 26)
	for i := 0; i < 26; i++ {
		next[i] = -1
		prev[i] = -1
	}

	for _, word := range words {
		ln := len(word)
		if ln > 26 {
			return "NO"
		}
		vis[int(word[0]-'a')] = true
		for i := 0; i+1 < ln; i++ {
			x := int(word[i] - 'a')
			y := int(word[i+1] - 'a')
			if next[x] < 0 && prev[y] < 0 {
				next[x] = y
				prev[y] = x
			} else if next[x] >= 0 && prev[y] < 0 || next[x] < 0 && prev[y] >= 0 || next[x] != y {
				return "NO"
			}
			vis[y] = true
		}
	}
	var sum int
	var buf bytes.Buffer

	for i := 0; i < 26; i++ {
		if vis[i] {
			sum++
			if prev[i] < 0 {
				// 这样子保证了i是一个起点，且每个字符只会在一条路径上
				for j := i; j >= 0; j = next[j] {
					sum--
					buf.WriteByte(byte('a' + j))
				}
			}
		}
	}

	if sum > 0 {
		// 如果sum > 0, 表明存在一个环
		return "NO"
	}

	return buf.String()
}
