package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbb"))
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
	fmt.Println(longestPalindromeSubseq("euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew"))
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	//dp[i][k] means longest sub seq starting from i, having length k
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
	}

	for k := 2; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			if s[i] != s[i+k-1] {
				dp[i][k] = max(dp[i][k-1], dp[i+1][k-1])
			} else if k > 3 {
				dp[i][k] = max(2+dp[i+1][k-2], dp[i][k-1])
			} else {
				dp[i][k] = k
			}
		}
	}

	return dp[0][n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
