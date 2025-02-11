Tsovak and Narek are playing a game. They have an array 𝑎
 and a matrix 𝑏
 of integers with 𝑛
 rows and 𝑚
 columns, numbered from 1
. The cell in the 𝑖
-th row and the 𝑗
-th column is (𝑖,𝑗)
.

They are looking for the elements of 𝑎
 in turns; Tsovak starts first. Each time a player looks for a cell in the matrix containing the current element of 𝑎
 (Tsovak looks for the first, then Narek looks for the second, etc.). Let's say a player has chosen the cell (𝑟,𝑐)
. The next player has to choose his cell in the submatrix starting at (𝑟+1,𝑐+1)
 and ending in (𝑛,𝑚)
 (the submatrix can be empty if 𝑟=𝑛
 or 𝑐=𝑚
). If a player cannot find such a cell (or the remaining submatrix is empty) or the array ends (the previous player has chosen the last element), then he loses.

Your task is to determine the winner if the players play optimally.

### ideas
1. a[i] <= 7
2. b[i] <= 7
3. 考虑a[0] = x, 这时候它可以对应的状态是一个位置列表
4. dp[i][r][c] 表示可以在a[i]，当前用户选择后的最后结果
5. dp[i][r][c] = win 如果全部 dp[i+1][x][y] = lose
6. dp[i][r][c] = lose 如果存在 dp[i+1][x][y] = win
7. 所以可以求和吗？全部lose = sum = 0， 存在win sum > 0
8. 所以是子区域求和？