package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	s := readString(reader)[:n]
	res := solve(s, m)
	if len(res) == 0 {
		res = "NO"
	}
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

func solve(s string, p int) string {
	// 从后往前
	n := len(s)

	if p == 1 || n == 1 {
		buf := []byte(s)
		// n must be 1
		x := int(buf[0] - 'a')
		if x+1 < p {
			buf[0] = byte(x + 1 + 'a')
			return string(buf)
		}
		return ""
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(s[i] - 'a')
	}

	return solve2(n, arr, p)
}

func solve2(n int, arr []int, p int) string {
	for i := n - 1; i >= 0; i-- {
		x := arr[i]
		for j := x + 1; j < p; j++ {
			if i > 0 && arr[i-1] == j || i > 1 && arr[i-2] == j {
				// 和前面的冲突了
				continue
			}
			arr[i] = j
			// 检查后面是否可以使用ab
			if i+1 < n {
				if i == 0 || arr[i-1] != 0 {
					arr[i+1] = 0
				} else {
					// arr[i+1] 不能取0， 只能是1/2
					arr[i+1] = 1
					if j == 1 {
						// 也不能取1
						arr[i+1] = 2
					}
				}
			}
			if i+2 < n {
				for y := 0; y < 3; y++ {
					if y != j && y != arr[i+1] {
						arr[i+2] = y
						break
					}
				}
			}

			for k := i + 3; k < n; k++ {
				arr[k] = 3 - arr[k-1] - arr[k-2]
			}
			return convert(arr)
		}
	}

	return ""
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func convert(arr []int) string {
	buf := make([]byte, len(arr))
	for i := 0; i < len(arr); i++ {
		buf[i] = byte(arr[i] + 'a')
	}
	return string(buf)
}
