package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	fmt.Println(solve(n))
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

const inf int = 1e18

func solve(n int) int {
	primes := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = inf
	}
	dp[1] = 1

	for j := 1; j < 11; j++ {
		for i := n; i > 0; i-- {
			tmp := inf
			pw := 1
			for k := 0; k < i; k++ {
				if i%(k+1) == 0 {
					if tmp/pw >= dp[i/(k+1)] {
						tmp = pw * dp[i/(k+1)]
					}
				}
				if pw >= inf/primes[j] {
					break
				}
				pw *= primes[j]
			}
			dp[i] = tmp
		}
	}
	return dp[n]
}

func abs(num int) int {
	return max(num, -num)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
