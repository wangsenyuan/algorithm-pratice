You are given two bracket sequences (not necessarily regular) 𝑠
 and 𝑡
 consisting only of characters '(' and ')'. You want to construct the shortest regular bracket sequence that contains both given bracket sequences as subsequences (not necessarily contiguous).

Recall what is the regular bracket sequence:

() is the regular bracket sequence;
if 𝑆
 is the regular bracket sequence, then (𝑆
) is a regular bracket sequence;
if 𝑆
 and 𝑇
 regular bracket sequences, then 𝑆𝑇
 (concatenation of 𝑆
 and 𝑇
) is a regular bracket sequence.
Recall that the subsequence of the string 𝑠
 is such string 𝑡
 that can be obtained from 𝑠
 by removing some (possibly, zero) amount of characters. For example, "coder", "force", "cf" and "cores" are subsequences of "codeforces", but "fed" and "z" are not.



### ideas
1. dp[i][j][k] 表示包含前i，j且level = k的最短的字符串
2. dp[n][m][0] 就是答案。然后再重构出来