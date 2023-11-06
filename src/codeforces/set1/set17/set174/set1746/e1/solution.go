package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	ask := func(arr []int) string {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? %d", len(arr)))
		for _, x := range arr {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readString(reader)
	}

	output := func(x int) string {
		fmt.Printf("! %d\n", x)
		return readString(reader)
	}

	solve(n, ask, output)
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

const yes = "YES"
const no = "NO"
const smile = ":)"

func solve(n int, ask func(arr []int) string, output func(int) string) {
	// 怎么判断是joking呢？
	// 能够确定的是x 属于 1...n
	// 先问一次 x是否在l...mid中，它回答是a 然后问一次是否在mid+1...r中, 它的回答是b
	// 如果 a = b, 则其中有一次肯定是错的，也许是a，也许是b
	// 如果 a != b, 则两次都是对的
	// 如果知道第一次的对错，是否就能处理后续的情况了？
	// 如果前一次是错的，那么本次肯定是对的
	// 如果前一次是对的，那就一半的区间检查两次？
	//  如果它回答 NO, NO ? 能确定一定不在吗？
	// 假设就在那一半，那么第一次no是错误的回答，第二次no就犯规了，两个no表示，肯定不在
	//  假设不在那一半，第一次no是正确的回答，但是第二次no 也是正确回答
	//  yes，yes也能确定不在那边
	// 那它回答 yes, no 呢？
	// 假设在那边，第一个yes是正确的，第二个no是骗人的， 或者第一个yes是骗人的，第二个no是正确的

	var dfs func(valid []int)

	dfs = func(valid []int) {
		if len(valid) < 3 {
			tmp := output(valid[0])
			if tmp != smile {
				output(valid[1])
			}
			return
		}
		if len(valid) == 3 {
			ok := make([]bool, 4)
			ok[0] = ask(valid[0:1]) == yes
			ok[1] = ask(valid[1:2]) == yes
			ok[2] = ask(valid[1:2]) == yes
			ok[3] = ask(valid[0:1]) == yes
			if ok[1] && ok[2] {
				dfs(valid[1:2])
				return
			}
			if !ok[1] && !ok[2] {
				dfs([]int{valid[0], valid[2]})
				return
			}
			if ok[0] && ok[1] || ok[2] && ok[3] {
				dfs(valid[:2])
				return
			}
			if ok[0] && !ok[1] || !ok[2] && ok[3] {
				dfs([]int{valid[0], valid[2]})
				return
			}
			dfs(valid[1:])
			return
		}
		var q1 []int
		var q2 []int
		for i := 0; i < len(valid); i++ {
			if i&1 > 0 {
				q1 = append(q1, valid[i])
			}
			if i&2 > 0 {
				q2 = append(q2, valid[i])
			}
		}
		is := make([]bool, 2)
		is[0] = ask(q1) == yes
		is[1] = ask(q2) == yes
		var next []int
		for i := 0; i < len(valid); i++ {
			if !(i&1 == 1) != is[0] || !(i&2 == 2) != is[1] {
				next = append(next, valid[i])
			}
		}
		dfs(next)
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	dfs(arr)
}
