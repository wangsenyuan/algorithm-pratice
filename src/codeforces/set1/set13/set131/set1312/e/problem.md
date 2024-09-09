You are given an array 𝑎1,𝑎2,…,𝑎𝑛
. You can perform the following operation any number of times:

Choose a pair of two neighboring equal elements 𝑎𝑖=𝑎𝑖+1
 (if there is at least one such pair).
Replace them by one element with value 𝑎𝑖+1
.
After each such operation, the length of the array will decrease by one (and elements are renumerated accordingly). What is the minimum possible length of the array 𝑎
 you can get?

### ideas
1. 假设前缀能够到达位置x, 但是如果变成x+1后，那么它是可以继续往前处理的，所以这样子不符合dp的调性
2. dp[l][r][x] 表示将l..r的数字处理为x (true / false)
3. dp[l][r][x] = dp[l+1][r][x-1] && a[l] == x - 1
4.           or = dp[l][r][x-1] && a[r] == x - 1
5. 应该是可行的
6. n * n * 2 * x = 500 * 500 * 2 * 1000 = 5 * 1e8 
7. too much.
8. 还需要优化
9. 对于l...r这一段，它的目标是明确的，就是最小值x + r - l (操作的次数)
10. dp[l][r] = true => 能够将 dp[l][r]变成这个区间内的最小值x， 变为 x + r - l
11. dp[l][r] = dp[l+1][r] && a[l] = (l+1...r)的最小 + r - (l + 1)
12. good