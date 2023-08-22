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
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
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

func solve(arr []int) []int {
	// pick [l..r] a = arr[l]
	// cnt = freq[a][r] - freq[a][l] + 1
	// 非a的数量 = r - l + 1 - cnt
	// cnt - (r - l + 1 - cnt) = 2 * cnt - (r - l + 1)
	// 2 * freq[a][r] - 2 * freq[a][l] + 2 - (r - l + 1)
	// 2 * freq[a][r] - r - (2 * freq[a][l] - l) + 1
	// 对于a来说需要维护前面的那个最大值
	freq := make(map[int]int)
	// 对于a来说，最大的幂
	// 对于r来说，最小的2 * freq[a][l] - l
	prev := make(map[int]Pair)
	ans := []int{arr[0], 1, 1, 1}
	update := func(l int, r int, v int) {
		if ans[3] < v {
			ans = []int{arr[l], l + 1, r + 1, v}
		}
	}
	for r, a := range arr {
		freq[a]++
		cur := 2*freq[a] - r + 1
		if v, ok := prev[a]; ok {
			tmp := cur - v.first
			update(v.second, r, tmp)
			if 2*freq[a]-r < v.first {
				prev[a] = Pair{2*freq[a] - r, r}
			}
		} else {
			prev[a] = Pair{2*freq[a] - r, r}
		}
	}

	return ans[:3]
}

type Pair struct {
	first  int
	second int
}
