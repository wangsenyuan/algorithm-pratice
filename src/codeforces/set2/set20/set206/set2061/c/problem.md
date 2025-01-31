Kevin enjoys logic puzzles.

He played a game with 𝑛
classmates who stand in a line. The 𝑖
-th person from the left says that there are 𝑎𝑖
liars to their left (not including themselves).

Each classmate is either honest or a liar, with the restriction that no two liars can stand next to each other. Honest
classmates always say the truth. Liars can say either the truth or lies, meaning their statements are considered
unreliable.

Kevin wants to determine the number of distinct possible game configurations modulo 998244353
. Two configurations are considered different if at least one classmate is honest in one configuration and a liar in the
other.

### ideas

1. 一头雾水
2. 考虑一个dp的算法
3. dp[i][j][0/1] 表示到i为止，有j个liar时，且i是一个liar 0, or not
4. dp[i][j][0] = dp[i-1][j-1][1] （不能有连续两个liar）
5. dp[i][j][1] = dp[i-1][j][0] + dp[i-1][j][1] 这里 j = a[i]
6. 从后往前看，如果 n是一个liar（那么它的答案没有参考意义）
7. 如果n不是liar， 那么意味着前面有 a[n]个liar
8. 感觉骗子就是隔板，在同一个区间内，a[?]等于它左边隔板的数量
9. dp[i][j][0] = dp[i-1][j-1][1] => dp[i+1][j][1] = dp[i-1][j-1][1] (中间多了个骗子)
10. dp[i][j] = 到i为止共有i个骗子，且i不是骗子时的计数
11. dp[i][j], 这里j = a[i]
12. 同时 dp[i+1][a[i+1]] += dp[i-1][a[i+1]-1]