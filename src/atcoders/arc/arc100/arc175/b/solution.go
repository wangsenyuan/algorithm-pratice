package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, A, B := readThreeNums(reader)
	s := readString(reader)

	res := solve(s, A, B)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string, A int, B int) int {
	n := len(s)
	// 不大对
	// 不得要领
	// 如果靠左有一个左括号，显然暂时不能替换和交换它
	// 把左右变成+1/-1, 那么是不是可以变成前缀和的形式
	// 假设交换 i, j, 那么对于j后面的前缀和来说，不会有影响
	// 如果替换i, 那么就会赢到i后面所有的数
	// 目的是使得，所有的前缀和>= 0, 且最后的前缀和 = 0
	// 可以使用两个replace 代替一个 swap
	A = min(A, 2*B)

	var level int
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			level++
		} else {
			level--
		}
	}
	// 如果level不等于0， 这些只能通过swap替换掉
	// 如果 level > 0, 那么必须替换左括号为右括号
	var res int

	buf := []byte(s)

	if level > 0 {
		// 左括号太多了，从后面往前replace
		for i := n - 1; level > 0; i-- {
			if buf[i] == '(' {
				level -= 2
				res += B
				buf[i] = ')'
			}
		}
	}
	if level < 0 {
		// 右括号太多了，从前面replace
		for i := 0; level < 0; i++ {
			if buf[i] == ')' {
				level += 2
				res += B
				buf[i] = '('
			}
		}
	}

	// level == 0
	// 然后就是用左括号和右括号交换
	var arr []int

	for i := 0; i < n; i++ {
		if buf[i] == '(' {
			arr = append(arr, i)
		}
	}
	// 肯定是各占一半的
	j := len(arr)

	for i := 0; i < n; i++ {
		if buf[i] == '(' {
			level++
		} else {
			level--
		}
		if level < 0 {
			buf[i], buf[arr[j-1]] = buf[arr[j-1]], buf[i]
			j--
			res += A
			level = 1
		}
	}

	return res
}


func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}