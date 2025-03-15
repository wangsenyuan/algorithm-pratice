package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		n := readNum(reader)
		res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(num int) int {

	// 最大是7
	var digits []int
	bases := []int{1}

	for n := num; n > 0; n /= 10 {
		digits = append(digits, n%10)
		bases = append(bases, 10*bases[len(bases)-1])
	}
	ans := 7

	for i, x := range digits {
		if x == 7 {
			return 0
		}
		if i == 0 {
			for j := 0; j < 7; j++ {
				if (x+9*j)%10 == 7 {
					ans = min(ans, j)
				}
			}
			continue
		}
		r := num % bases[i]
		if x < 7 {
			// 在低位使用 999
			tmp := 7 - x
			if r < tmp {
				tmp++
			}
			ans = min(ans, tmp)
		}
		// 在当前位使用9999
		for j := 0; j < ans; j++ {
			y := x*bases[i] + r + j*(bases[i+1]-1)
			y /= bases[i]
			if y%10 == 7 {
				ans = min(ans, j)
			}
		}
	}

	return ans
}
