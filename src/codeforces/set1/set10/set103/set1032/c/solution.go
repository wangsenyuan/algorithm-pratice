package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(a []int) []int {
	n := len(a)
	if n == 1 {
		return []int{1}
	}
	prev := make([][]int, n)
	for i := 0; i < n; i++ {
		prev[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			prev[i][j] = -1
		}
	}

	for j := 0; j < 5; j++ {
		prev[0][j] = j
	}

	for i := 1; i < n; i++ {
		if a[i-1] < a[i] {
			it := -1
			for j := 0; j < 5; j++ {
				prev[i][j] = it
				if prev[i-1][j] >= 0 {
					it = j
				}
			}
		} else if a[i-1] > a[i] {
			it := -1
			for j := 4; j >= 0; j-- {
				prev[i][j] = it
				if prev[i-1][j] >= 0 {
					it = j
				}
			}
		} else {
			var flag int
			for j := 0; j < 5; j++ {
				if prev[i-1][j] >= 0 {
					flag |= 1 << j
				}
			}
			if flag == 0 {
				// no answer
				return nil
			}
			for j := 0; j < 5; j++ {
				tmp := flag
				if (tmp>>j)&1 == 1 {
					tmp ^= 1 << j
				}

				prev[i][j] = findFirstSetBit(tmp)
			}
		}
	}
	ans := make([]int, n)
	ans[n-1] = -1
	for j := 0; j < 4; j++ {
		if prev[n-1][j] >= 0 {
			ans[n-1] = j
			break
		}
	}
	if ans[n-1] < 0 {
		return nil
	}

	for i := n - 1; i > 0; i-- {
		ans[i-1] = prev[i][ans[i]]
	}

	for i := 0; i < n; i++ {
		ans[i]++
	}

	return ans
}

func findFirstSetBit(num int) int {
	var res int
	for num > 0 {
		if num&1 == 1 {
			return res
		}
		res++
		num >>= 1
	}
	return -1
}
