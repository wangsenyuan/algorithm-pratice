You are given a string 𝑠
 and a string 𝑡
, both consisting only of lowercase Latin letters. It is guaranteed that 𝑡
 can be obtained from 𝑠
 by removing some (possibly, zero) number of characters (not necessary contiguous) from 𝑠
 without changing order of remaining characters (in other words, it is guaranteed that 𝑡
 is a subsequence of 𝑠
).

For example, the strings "test", "tst", "tt", "et" and "" are subsequences of the string "test". But the strings "tset", "se", "contest" are not subsequences of the string "test".

You want to remove some substring (contiguous subsequence) from 𝑠
 of maximum possible length such that after removing this substring 𝑡
 will remain a subsequence of 𝑠
.

If you want to remove the substring 𝑠[𝑙;𝑟]
 then the string 𝑠
 will be transformed to 𝑠1𝑠2…𝑠𝑙−1𝑠𝑟+1𝑠𝑟+2…𝑠|𝑠|−1𝑠|𝑠|
 (where |𝑠|
 is the length of 𝑠
).

Your task is to find the maximum possible length of the substring you can remove so that 𝑡
 is still a subsequence of 𝑠
.

### ideas
1. dp[l]表示s[..l]能够保护的t最长的前缀
2. fp[r]表示s[r..]能够保护的t的最长的后缀
3. 当dp[l] + fp[r] == len(t)的时候，删除l...r即可
4. 这样显然不大行(n * n的复杂性)
5. 反过来考虑t, dp[l]表示t[..l]需要的最短的前缀s[..], fp[r]表示t[r...]需要的最短的后缀s[..]
6. 那么 dp[l], fp[l+1] 如果不超过s的长度,那么就是一个可行的方案