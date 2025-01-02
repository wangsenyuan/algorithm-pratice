A subsequence is a string that can be derived from another string by deleting some or no symbols without changing the order of the remaining symbols. Characters to be deleted are not required to go successively, there can be any gaps between them. For example, for the string "abaca" the following strings are subsequences: "abaca", "aba", "aaa", "a" and "" (empty string). But the following strings are not subsequences: "aabaca", "cb" and "bcaa".

You are given a string 𝑠
 consisting of 𝑛
 lowercase Latin letters.

In one move you can take any subsequence 𝑡
 of the given string and add it to the set 𝑆
. The set 𝑆
 can't contain duplicates. This move costs 𝑛−|𝑡|
, where |𝑡|
 is the length of the added subsequence (i.e. the price equals to the number of the deleted characters).

Your task is to find out the minimum possible total cost to obtain a set 𝑆
 of size 𝑘
 or report that it is impossible to do so.

Input
The first line of the input contains two integers 𝑛
 and 𝑘
 (1≤𝑛≤100,1≤𝑘≤1012
) — the length of the string and the size of the set, correspondingly.

The second line of the input contains a string 𝑠
 consisting of 𝑛
 lowercase Latin letters.


 ### ideas
 1. dp[i][j] 表示前缀，且集合且s的长度为j时的数量
 2.  dp[i][j] = dp[i-1][j] + dp[i-1][j-1] + (s[i])
 3.  但是问题是 asssa  那么长度为3的asa 只有一个（而不是3个）怎么把这个给挑出来？
 4.  但是 ass (ssa) 也要算进去
 5.  dp[i][j] = dp[i-1][j] + ?
 6.   长度为j的元素有多少个？
 7. 主要是字符相同的时候，怎么搞
 8. 如果 s[i] != s[i-1], dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
 9. 如果 s[i] = s[i-1], dp[i][j] = dp[i-1][j-1] (只需要计算加上去的)