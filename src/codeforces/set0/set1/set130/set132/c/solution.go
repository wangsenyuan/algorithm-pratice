package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	s := readString(reader)
	n := readNum(reader)
	return solve(s, n)
}

const inf = 1 << 30

func solve(command string, n int) int {
	dp := make([]int, (n+1)*2)
	for i := range len(dp) {
		dp[i] = -inf
	}
	dp[0] = 0
	dp[1] = 0
	fp := make([]int, len(dp))
	move := []int{1, -1}

	for _, x := range []byte(command) {
		for j := range len(fp) {
			fp[j] = -inf
		}
		var v int
		if x == 'T' {
			v++
		}
		for j := 0; j <= n; j++ {
			for f := 0; f <= 1; f++ {
				// 在当前节点执行k次操作
				for k := 0; k+j <= n; k++ {
					if (k&1 + v) == 1 {
						// 交换后，这里是将F变成T
						fp[(j+k)*2+f] = max(fp[(j+k)*2+f], dp[j*2+f^1])
					} else {
						// 交换后，这里是F
						fp[(j+k)*2+f] = max(fp[(j+k)*2+f], dp[j*2+f]+move[f])
					}
				}
			}

		}
		copy(dp, fp)
	}

	return max(dp[n*2+0], dp[n*2+1])
}

func abs(num int) int {
	return max(num, -num)
}
