package main

import (
	"bufio"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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
	n := readNum(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	res := solve(words)
	fmt.Println(res)
}

const H = 24

func solve(words []string) int64 {
	M := 1 << H
	cnt := make([]int, M)

	for _, word := range words {
		a := int(word[0] - 'a')
		b := int(word[1] - 'a')
		c := int(word[2] - 'a')
		var flag int
		if a >= 0 {
			flag |= 1 << a
		}
		if b >= 0 {
			flag |= 1 << b
		}
		if c >= 0 {
			flag |= 1 << c
		}
		cnt[flag]++
	}

	for i := 0; i < H; i++ {
		for mask := 0; mask < M; mask++ {
			if (mask>>i)&1 == 1 {
				cnt[mask] += cnt[mask^(1<<i)]
			}
		}
	}
	var ans int64
	n := len(words)
	for x := 0; x < M; x++ {
		ans ^= int64(n-cnt[x]) * int64(n-cnt[x])
	}

	return ans
}
