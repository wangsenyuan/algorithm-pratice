package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func solve(x []string) int {
	n := len(x)
	words := make([]ShWord, n)
	for i, w := range x {
		cnt := make([]int, 2)
		for j := 0; j < len(w); j++ {
			if w[j] == 's' {
				cnt[0]++
			} else {
				cnt[1]++
			}
		}
		words[i] = ShWord{cnt, i}
	}

	slices.SortFunc(words, func(a, b ShWord) int {
		u := a.cnt[0] * b.cnt[1]
		v := a.cnt[1] * b.cnt[0]
		return v - u
	})
	var res int
	// the result
	var cnt int
	for _, cur := range words {
		i := cur.id
		for j := 0; j < len(x[i]); j++ {
			if x[i][j] == 'h' {
				res += cnt
			} else {
				cnt++
			}
		}
	}

	return res
}

type ShWord struct {
	cnt []int
	id  int
}
