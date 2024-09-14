package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	cards := make([]string, n)
	for i := 0; i < n; i++ {
		cards[i] = readString(reader)
	}

	res := solve(cards, k)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(cards []string, k int) int {
	freq := make(map[int]int)

	n := len(cards)

	get := func(x byte) int {
		if x == 'S' {
			return 0
		}
		if x == 'E' {
			return 1
		}
		return 2
	}

	// SET
	getValue := func(i int) int {
		var res int
		for j := 0; j < k; j++ {
			res = res*3 + get(cards[i][j])
		}

		return res
	}

	for i := 0; i < n; i++ {
		freq[getValue(i)]++
	}
	var res int

	for i := 0; i < n; i++ {
		// freq现在只有i后面的计数
		freq[getValue(i)]--
		for j := i - 1; j >= 0; j-- {
			// 如果其中两个是j
			var expect int
			for u := 0; u < k; u++ {
				var v int
				if cards[i][u] == cards[j][u] {
					// 需要相同的值
					v = get(cards[i][u])
				} else {
					// 需要不同的值
					v = 3 - get(cards[i][u]) - get(cards[j][u])
				}
				expect = expect*3 + v
			}
			res += freq[expect]
		}
	}

	return res
}
