You are given an integer array 𝑎
 of length 𝑛
.

You can perform the following operation: choose an element of the array and replace it with any of its neighbor's value.

For example, if 𝑎=[3,1,2]
, you can get one of the arrays [3,3,2]
, [3,2,2]
 and [1,1,2]
 using one operation, but not [2,1,2
] or [3,4,2]
.

Your task is to calculate the minimum possible total sum of the array if you can perform the aforementioned operation at most 𝑘
 times.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

The first line of each test case contains two integers 𝑛
 and 𝑘
 (1≤𝑛≤3⋅105
; 0≤𝑘≤10
).

The second line contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤109
).

Additional constraint on the input: the sum of 𝑛
 over all test cases doesn't exceed 3⋅105
.


### ideas
1. 如果只有一次操作，那么最优的方案是，选择 a[i] - a[i+1] 最大的进行替换(a[i] - a[i-1])
2. k是10次，那可以换个dp的思路了
3. dp[i][j] 表示到i为止进行了j次替换后的最优解
4. 还需要知道a[i]现在等于多少，它不可能等于很远的数，距离不会超过j
5. dp[i][j][l] 表示到i为止，进行j次替换后的最后解，且a[i] 现在的值等于附近 i + l 处的值，l可以是负值
6. dp[i+1][j][-1] 表示i+1处的值更改为a[i+2]
7. dp[i+1][j+x][-x] 表示 i+1处的值更改为 a[i+1-x]，但是这个时候，中间的都必须被修改掉
8. dp[i+1][j+x][-x] = dp[i][j+x][-(x+1)] (i+1不进行操作) 且a[i]的值 = a[i+x+1]
9. dp[i+1][j+x][x] = dp[i][j+x-1][x-1] 如果 a[i] = a[i-x]
10. dp[i+1][j][0] = min(dp[i][j][x], x >= 0) a[i]等于它前面的值，且i+1不操作
11. 这个dp这么复杂吗？
12. 再想想，上面的做法不大对
13. 假设a[i]用来替换别的数，那么它没有被替换的道理
14. dp[i][j]表示在i前面使用了j次替换后的最小的前缀和
15. 假设现在要计算 dp[i][j], 那么就是使用i（包括i）前面的某个下标l 且a[l]是[l...i]的最小值
16. 去替换 a[l...i]
17. 并使用 a[l] 去替换前面的一个序列