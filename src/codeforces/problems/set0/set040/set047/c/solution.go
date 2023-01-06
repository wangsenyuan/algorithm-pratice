package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	words := make([]string, 6)

	for i := 0; i < 6; i++ {
		words[i] = readString(reader)
	}
	res := solve(words)

	if len(res) == 0 {
		fmt.Println("Impossible")
		return
	}
	for _, row := range res {
		fmt.Println(row)
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(words []string) []string {
	// the longest words must be the one put vertically or horizongtally
	// try all combinations

	var dfs func(i int, flag int) [][]int

	dfs = func(i int, flag int) [][]int {
		if i == 6 {
			return [][]int{{}}
		}
		var res [][]int
		for x := 0; x < 6; x++ {
			if (flag>>x)&1 == 0 {
				tmp := dfs(i+1, flag|(1<<x))
				for _, row := range tmp {
					var cur []int
					cur = append(cur, x)
					cur = append(cur, row...)
					res = append(res, cur)
				}
			}
		}
		return res
	}

	put := func(buf [][]byte, word string, x, y int, dx int, dy int) bool {
		var k int
		for i, j := x, y; i < len(buf) && j < len(buf[i]) && k < len(word); i, j, k = i+dx, j+dy, k+1 {
			if buf[i][j] == '.' || buf[i][j] == word[k] {
				buf[i][j] = word[k]
				continue
			}
			return false
		}

		return k == len(word)
	}

	place := func(pos []int) []string {
		w := len(words[pos[0]])
		h := len(words[pos[1]])
		buf := make([][]byte, h)
		for i := 0; i < h; i++ {
			buf[i] = make([]byte, w)
			for j := 0; j < w; j++ {
				buf[i][j] = '.'
			}
		}
		//
		if !put(buf, words[pos[2]], 0, 0, 0, 1) {
			return nil
		}
		if !put(buf, words[pos[3]], 0, 0, 1, 0) {
			return nil
		}

		if !put(buf, words[pos[0]], len(words[pos[3]])-1, 0, 0, 1) {
			return nil
		}

		if !put(buf, words[pos[1]], 0, len(words[pos[2]])-1, 1, 0) {
			return nil
		}

		if !put(buf, words[pos[4]], len(words[pos[3]])-1, w-1, 1, 0) {
			return nil
		}

		if !put(buf, words[pos[5]], h-1, len(words[pos[2]])-1, 0, 1) {
			return nil
		}
		if buf[1][1] != '.' || buf[0][w-1] != '.' || buf[h-1][0] != '.' || buf[h-2][w-2] != '.' {
			return nil
		}

		return make_strings(buf)
	}

	arr := dfs(0, 0)

	var ans []string

	for _, row := range arr {
		s := place(row)
		if len(s) == 0 {
			continue
		}
		if len(ans) == 0 || less_strings(s, ans) {
			ans = s
		}
	}

	return ans
}

func less_strings(a, b []string) bool {
	var i int
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	if i == len(a) {
		return true
	}
	if i == len(b) {
		return false
	}
	return a[i] < b[i]
}

func make_strings(arr [][]byte) []string {
	res := make([]string, len(arr))
	for i, row := range arr {
		res[i] = string(row)
	}
	return res
}
