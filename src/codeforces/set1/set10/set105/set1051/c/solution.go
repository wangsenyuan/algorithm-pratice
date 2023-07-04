package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		nums := readNNums(reader, n)
		res := solve(nums)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES\n%s\n", res))
		}
	}
	fmt.Print(buf.String())
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

func solve(nums []int) string {
	type Pair struct {
		first  int
		second int
	}

	n := len(nums)
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{nums[i], i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	var singles []int
	mulps := make(map[int]int)

	for i := 0; i < n; {
		j := i
		for i < n && ps[i].first == ps[j].first {
			i++
		}
		if i == j+1 {
			singles = append(singles, ps[j].second)
		} else {
			mulps[ps[j].second] = i - j
		}
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = 'A'
	}
	for i := 0; i < len(singles); i += 2 {
		if i+1 < len(singles) {
			res[singles[i+1]] = 'B'
		}
	}
	if len(singles)%2 == 0 {
		// divides it equally
		return string(res)
	}
	// one has more
	// 如果有任何一个多余2个
	for i, v := range mulps {
		if v > 2 {
			res[i] = 'B'
			return string(res)
		}
	}

	return ""
}
