package main

import "fmt"

func minStickers(stickers []string, target string) int {
	n := len(target)
	m := 1 << uint(n)
	dp := make([]int, m)

	for i := 1; i < m; i++ {
		dp[i] = -1
	}

	for i := 0; i < m; i++ {
		if dp[i] == -1 {
			continue
		}
		for _, sticker := range stickers {
			now := i
			for j := 0; j < len(sticker); j++ {
				for k := 0; k < n; k++ {
					if target[k] == sticker[j] && (now>>uint(k))&1 == 0 {
						now |= 1 << uint(k)
						break
					}
				}
			}
			if dp[now] == -1 || dp[now] > dp[i]+1 {
				dp[now] = dp[i] + 1
			}
		}
	}
	return dp[m-1]
}

func main() {
	fmt.Println(minStickers([]string{"with", "example", "science"}, "thehat"))
	fmt.Println(minStickers([]string{"notice", "possible"}, "basicbasic"))
	fmt.Println(minStickers([]string{"these", "guess", "about", "garden", "him"}, "atomher"))
}
