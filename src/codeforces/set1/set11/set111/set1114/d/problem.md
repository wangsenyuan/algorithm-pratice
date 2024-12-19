You are given a line of 𝑛
 colored squares in a row, numbered from 1
 to 𝑛
 from left to right. The 𝑖
-th square initially has the color 𝑐𝑖
.

Let's say, that two squares 𝑖
 and 𝑗
 belong to the same connected component if 𝑐𝑖=𝑐𝑗
, and 𝑐𝑖=𝑐𝑘
 for all 𝑘
 satisfying 𝑖<𝑘<𝑗
. In other words, all squares on the segment from 𝑖
 to 𝑗
 should have the same color.

For example, the line [3,3,3]
 has 1
 connected component, while the line [5,2,4,4]
 has 3
 connected components.

The game "flood fill" is played on the given line as follows:

At the start of the game you pick any starting square (this is not counted as a turn).
Then, in each game turn, change the color of the connected component containing the starting square to any other color.
Find the minimum number of turns needed for the entire line to be changed into a single color.

### ideas
1. 假设选中了位置i,作为起始位置
2. 那么需要多少次调整才行呢？
3. 假设i = 1, 这个策略是不是，只要遇到一个新的segment，就要转换呢？
4. 从后面开始，确实是这样的。问题出在中间的i
5. 假设从i开始，往前后处理到了l...r, [l+1...r-1]的颜色，变化后可以吧l...r给连接起来
6. 那么这种情况下，就可以省掉很多操作
7. dp[l][r] = dp[l+1][r-1] + 1 if a[l] = a[r]
8.  else = min(dp[l+1][r], dp[l][r-1]) + 1