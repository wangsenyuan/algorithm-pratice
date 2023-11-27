package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
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

func solve(s string) string {
	// 要把1尽可能的放在2的前面，把0放在1的前面
	n := len(s)
	/// 21 不能直接交换，因为后面可能遇到0
	// 先交换后，0不能往前了
	// 201
	// 用0去交换1，再用1去交换2
	buf := []byte(s)
	p1 := -1
	for i := 0; i < n; i++ {
		if buf[i] == '0' {
			if p1 < 0 {
				// no 1 to swap
				continue
			}
			buf[i], buf[p1] = buf[p1], buf[i]
			for p1 < i && buf[p1] == '0' {
				// move to next 1
				p1++
			}
		} else if buf[i] == '1' {
			if p1 < 0 {
				p1 = i
			}
		} else {
			// buf[i] = 2, a block, can't swap before it
			p1 = -1
		}
	}

	// swap 1 and 2, 1 always can swap with 2, but not actually swap
	// kind of insert
	var res []int
	vis := make([]bool, n)
	for i, r := 0, 0; i < n; i++ {
		if vis[i] {
			continue
		}
		if buf[i] != '2' {
			res = append(res, i)
			continue
		}
		// buf[i] = 2
		// 要把后面所有的1都移动到这里
		if r <= i {
			r = i + 1
		}
		for ; r < n; r++ {
			if buf[r] == '1' {
				res = append(res, r)
				vis[r] = true
			}
		}
		res = append(res, i)
	}
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = buf[res[i]]
	}

	return string(ans)
}
