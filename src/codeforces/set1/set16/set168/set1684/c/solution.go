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
		n, m := readTwoNums(reader)
		board := make([][]int, n)
		for i := 0; i < n; i++ {
			board[i] = readNNums(reader, m)
		}
		res := solve(board)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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
func solve(board [][]int) []int {

	n := len(board)
	//m := len(board[0])

	theRow := -1

	for i := 0; i < n; i++ {
		if !sort.IntsAreSorted(board[i]) {
			theRow = i
			break
		}
	}

	if theRow < 0 {
		return []int{1, 1}
	}

	ans := check(board[theRow])
	// 2, 2, 2, 1 只能是 第一个 2 和1
	// 2，1，1，2， 只能是第一个2和最后一个1
	// 6， 2， 1， 2
	if len(ans) == 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		x, y := ans[0], ans[1]
		board[i][x], board[i][y] = board[i][y], board[i][x]
		ok := sort.IntsAreSorted(board[i])
		board[i][x], board[i][y] = board[i][y], board[i][x]
		if !ok {
			return nil
		}
	}
	return []int{ans[0] + 1, ans[1] + 1}
}

func check(arr []int) []int {
	n := len(arr)
	for i := 0; i+1 < n; i++ {
		if arr[i] > arr[i+1] {
			// need to move arr[i] to back or move arr[i+1] to front
			k := search(i, func(j int) bool {
				return arr[j] > arr[i+1]
			})
			// move i+1 to k, make sure  the prefix is sorted still
			if arr[i] == arr[k] {
				arr[k], arr[i+1] = arr[i+1], arr[k]
				ok := sort.IntsAreSorted(arr)
				arr[k], arr[i+1] = arr[i+1], arr[k]
				if ok {
					return []int{k, i + 1}
				}
			}
			// or move arr[i] to some suffix position
			// [i+1...n]
			j := n
			for j-1 > i && arr[j-1] >= arr[i] {
				j--
			}
			j--
			// 必须找到一个位置它至少要比arr[i]大
			// arr[j] < arr[i]
			arr[i], arr[j] = arr[j], arr[i]
			ok := sort.IntsAreSorted(arr)
			arr[i], arr[j] = arr[j], arr[i]
			if ok {
				return []int{i, j}
			}
			return nil
		}
	}
	return nil
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
