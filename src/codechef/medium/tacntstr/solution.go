package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(solve(scanner.Bytes()))
}

const MOD = 1000000007

func solve(s []byte) int64 {
	n := len(s)

	var ans, prefix int64

	ans = int64('Z' - s[0])
	prefix = ans

	for i := 1; i < n; i++ {
		x := int64('Z' - s[i])
		ans = (ans + (prefix+1)*x) % MOD
		prefix = ((prefix * 26) + x) % MOD
	}
	return ans
}
