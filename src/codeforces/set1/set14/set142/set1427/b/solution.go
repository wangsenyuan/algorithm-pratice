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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
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

func solve(s string, k int) int {
	n := len(s)
	var wins int
	var win_steaks int
	var gaps []int
	for i := 0; i < n; i++ {
		if s[i] == 'W' {
			wins++
			if i == 0 || s[i-1] == 'L' {
				win_steaks++
			}
		} else {
			if i == 0 || s[i-1] == 'W' {
				gaps = append(gaps, 0)
			}
			gaps[len(gaps)-1]++
		}
	}

	if wins+k >= n {
		return 2*n - 1
	}

	if wins == 0 {
		if k == 0 {
			return 0
		}
		return 2*k - 1
	}

	if s[0] == 'L' {
		gaps = gaps[1:]
	}
	if s[n-1] == 'L' && len(gaps) > 0 {
		gaps = gaps[:len(gaps)-1]
	}

	wins += k
	sort.Ints(gaps)
	it := 0
	for it < len(gaps) && k >= gaps[it] {
		k -= gaps[it]
		it++
		win_steaks--
	}
	return 2*wins - win_steaks
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
