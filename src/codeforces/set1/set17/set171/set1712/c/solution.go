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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, n)
		res := solve(P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
func solve(A []int) int {
	n := len(A)
	first := make([]int, n+1)
	for i := 0; i <= n; i++ {
		first[i] = -1
	}
	last := make([]int, n+1)
	for i := 0; i < n; i++ {
		if first[A[i]] < 0 {
			first[A[i]] = i
		}
		last[A[i]] = i
	}

	type Pair struct {
		first  int
		second int
	}

	var events []Pair

	for i := 1; i <= n; i++ {
		if first[i] >= 0 {
			events = append(events, Pair{first[i], i})
			events = append(events, Pair{last[i], -i})
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].first < events[j].first || events[i].first == events[j].first && events[i].second > events[j].second
	})

	sorted := make([]bool, n)
	sorted[n-1] = true
	for i := n - 2; i >= 0; i-- {
		sorted[i] = false
		if i+1 < n && sorted[i+1] && A[i] <= A[i+1] {
			sorted[i] = true
		}
	}

	if sorted[0] {
		return 0
	}

	res := n
	var open int
	nums := make(map[int]bool)
	for _, evt := range events {
		if evt.second > 0 {
			open++
		} else {
			open--
		}
		nums[abs(evt.second)] = true
		if open == 0 {
			x := evt.first
			if x+1 == n || sorted[x+1] {
				res = min(res, len(nums))
			}
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
