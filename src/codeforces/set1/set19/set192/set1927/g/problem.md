A horizontal grid strip of 𝑛
 cells is given. In the 𝑖
-th cell, there is a paint charge of size 𝑎𝑖
. This charge can be:

either used to the left — then all cells to the left at a distance less than 𝑎𝑖
 (from max(𝑖−𝑎𝑖+1,1)
 to 𝑖
 inclusive) will be painted,
or used to the right — then all cells to the right at a distance less than 𝑎𝑖
 (from 𝑖
 to min(𝑖+𝑎𝑖−1,𝑛)
 inclusive) will be painted,
or not used at all.
Note that a charge can be used no more than once (that is, it cannot be used simultaneously to the left and to the right). It is allowed for a cell to be painted more than once.

What is the minimum number of times a charge needs to be used to paint all the cells of the strip?

Input
The first line of the input contains an integer 𝑡
 (1≤𝑡≤100
) — the number of test cases in the test. This is followed by descriptions of 𝑡
 test cases.

Each test case is specified by two lines. The first one contains an integer 𝑛
 (1≤𝑛≤100
) — the number of cells in the strip. The second line contains 𝑛
 positive integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤𝑛
), where 𝑎𝑖
 is the size of the paint charge in the 𝑖
-th cell from the left of the strip.

It is guaranteed that the sum of the values of 𝑛
 in the test does not exceed 1000
.

Output
For each test case, output the minimum number of times the charges need to be used to paint all the cells of the strip.

### ideas
1. 如果i被用来进行向右边的charge，那么i...i + a[i]-1 中间的，要么不使用，要么只能向右
2. 如果i和i+1，一个向左，一个向右, 那么问题就变成了两个更小的问题
3. dp[l][r] 表示的是在区间l...r中间使用的最小的charge，以完整覆盖区间l...r
4. dp[l][r] = 1 + dp[l+1...r].... dp[l + a[i]...r] (如果在使用l处的)
5.          = 1 + dp[l...r - a[r]] .... dp[l...a[r]-1] 如果使用r处的
6.          = dp[l...i] + dp[i+1....r] 如果使用(i, i+1) 
7. 【2, 2, 5, 1, 6, 1, 8, 2, 8, 2】 expect 2
8. 这个例子说明上面的算法不对
9.  这里的问题在于，这个dp的方向是不定的（正常情况下是递增的）
10. dp[i] = 不管使用什么方式，将[1...i]都cover住的最优解
11. dp[i] = min(dp[i-a[i]...i-1]) + 1 或者 min(dp[j], dp[j+2...i-1]) + 1 如果j + 1 + a[j+1] > i
12. 还需要知道i处的有没有被用掉
13. dp[i][0] = 没有使用i处的，但是cover住了[1...i]的最优解
14. dp[i][1] = 使用i处，cover住了
15. dp[i][1] = 1 + min(dp[i-a[i]...i-1][0/1])
16.          = 2 + min(dp[j][0]) if j+a[j] >= i - a[i] + 1
17. dp[i][0] = 1 + min(dp[j][0]) if j + a[j] - 1 >= a[i]
18. dp[i] 表示前i个被cover住的cost
19. dp[i] = dp[j] + 1 如果 j >= i - a[i]
20.       = dp[j] + 2 如果 j+1...i 能够被完全cover， 且使用j+1...i中的两个，f，s
21.  f向右，s向左, f <= s + 1
22. 这里固定f的情况下，s从 f - 1到n进行判断，是否能够cover住 min(f, s - a[s] + 1), max(f + a[f] - 1, s)
23. 