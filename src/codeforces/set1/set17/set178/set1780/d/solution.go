package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(x int) int {
		fmt.Println(fmt.Sprintf("- %d", x))
		return readNum(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		cnt := readNum(reader)
		res := solve(cnt, ask)
		fmt.Println(fmt.Sprintf("! %d", res))
	}
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

func solve(cnt int, ask func(x int) int) int {
	// 假设数n有x个1
	// 如果用1去减它会怎么样？
	// 如果 n最后一位是1，那么减去1以后， cnt =》 cnt - 1
	// 否则的话 cnt会不变 (10), 或者变大 (1???0)
	//  diff = 表示的就是下一个1的位置-1
	var ans int
	pos := 0
	for cnt > 0 {
		// 这里可能是不对的
		x := (1 << pos)
		tmp := ask(x)
		if tmp < 0 {
			panic("fail")
		}
		ans += x
		if tmp == cnt-1 {
			// 最后一位是1
			pos++
			cnt = tmp
		} else if tmp == cnt && cnt == 1 {
			// only one digit
			ans += 1 << pos
			cnt--
			break
		} else {
			diff := tmp - cnt
			cnt = tmp
			// tmp >= cnt
			// 开始时n = ???100
			//  -x后n = ???011
			// 要把这些1减掉
			// 即使diff = 0 也没有关系
			y := ((1 << (diff + 1)) - 1) << pos
			cnt = ask(y)
			// ???0111
			// pos + 1不需要检查
			pos += diff + 2

			ans += y
		}

	}

	return ans
}
