Being a physicist, Charlie likes to plan his life in simple and precise terms.

For the next 𝑚
 months, starting with no money, Charlie will work hard and earn 𝑥
 pounds per month. For the 𝑖
-th month (1≤𝑖≤𝑚)
, there'll be a single opportunity of paying cost 𝑐𝑖
 pounds to obtain happiness ℎ𝑖
.

Borrowing is not allowed. Money earned in the 𝑖
-th month can only be spent in a later 𝑗
-th month (𝑗>𝑖
).

Since physicists don't code, help Charlie find the maximum obtainable sum of happiness.

Input
The first line of input contains a single integer 𝑡
 (1≤𝑡≤1000
) — the number of test cases.

The first line of each test case contains two integers, 𝑚
 and 𝑥
 (1≤𝑚≤50
, 1≤𝑥≤108
) — the total number of months and the monthly salary.

The 𝑖
-th of the following 𝑚
 lines contains two integers, 𝑐𝑖
 and ℎ𝑖
 (0≤𝑐𝑖≤108
, 1≤ℎ𝑖≤103
) — the cost and happiness on offer for the 𝑖
-th month. Note that some happiness may be free (𝑐𝑖=0
 for some 𝑖
's).

It is guaranteed that the sum of ∑𝑖ℎ𝑖
 over all test cases does not exceed 105
.

### ideas
1. 一个想法是，假设要在第i个月购买，那么alice至少要有c[i]的money
2. dp[i][j]表示在第i天时，剩余j的钱时，所购买的最大happiness
3. dp[i][j- c[i] + x] = dp[i-1][j] + h[i]
4. 但是j的范围太大了，所以这里换成alice有几个月的工资在手
5. dp[i][j + 1] = dp[i-1][k] + h[i] 如果 (j - k) * x >= c[i]
6. +1是因为它这个月后有多了一次工资
7. 这里出现一个问题，就是(j - k) * x - c[i]这个东西不一定整除x
8. 考虑一个例子, 比如 x = 3
9.  [0, ?], 【1， 4], [5, 4]
10. 在第一个月后获得3元钱， 购买了4，剩余2元钱， 再挣到3月钱， 剩余5元
11. 然后第二个月可以购买4，总共8
12. 但是 dp[1][j?]这个j不存在
13. dp[i][x] 后面用个map来表示？