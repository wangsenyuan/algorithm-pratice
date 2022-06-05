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
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(A []int) int {
	// at most n operations， 只从一头操作
	// 任何一个重复的数字，要需要修改
	// 假设现在处理完i位，要么i移动到了0位，并被删除后放到了n-1位，要么反过来
	// 各位的相对位置不变
	// 最优的策略应该是先往一个方向转动，然后从另外一个方向转动
	// 假设先向右转动x位，再向左转动y位，则总数为2 * x + y, 在y给定的情况下，算出x
	x := process(A)
	reverse(A)
	y := process(A)
	return min(x, y)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func process(arr []int) int {
	n := len(arr)
	// arr[x] 在 arr[y+1:x]中出现过, 且arr[y..x]中间无重复项
	var x, y int
	cnt := make(map[int]int)

	res := n

	for x < n {
		cnt[arr[x]]++
		for len(cnt) < x-y+1 {
			if cnt[arr[y]] == 1 {
				delete(cnt, arr[y])
			} else {
				cnt[arr[y]]--
			}
			y++
		}
		tmp := y
		if y > 0 {
			tmp += 2 * (n - x - 1)
		} else {
			tmp = n - x - 1
		}
		res = min(res, tmp)
		x++
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
