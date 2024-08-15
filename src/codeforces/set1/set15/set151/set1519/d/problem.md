You are given two integer arrays 𝑎
 and 𝑏
 of length 𝑛
.

You can reverse at most one subarray (continuous subsegment) of the array 𝑎
.

Your task is to reverse such a subarray that the sum ∑𝑖=1𝑛𝑎𝑖⋅𝑏𝑖
 is maximized.

Input
The first line contains one integer 𝑛
 (1≤𝑛≤5000
).

The second line contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤107
).

The third line contains 𝑛
 integers 𝑏1,𝑏2,…,𝑏𝑛
 (1≤𝑏𝑖≤107
).

Output
Print single integer — maximum possible sum after reversing at most one subarray (continuous subsegment) of 𝑎
.

### ideas
1. dp[l][r] = 在区间l...r中间进行了reverse后的，最大值
2. dp[l][r] = max(dp[l][r-1], dp[l+1][r]) 或者对l...r进行反转后的结果
3. 需要计算出pref, 和 suff，以加快dp[l][r]的计算