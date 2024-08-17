The difference between the versions is in the costs of operations. Solution for one version won't work for another!

Alice has a grid of size 𝑛×𝑚
, initially all its cells are colored white. The cell on the intersection of 𝑖
-th row and 𝑗
-th column is denoted as (𝑖,𝑗)
. Alice can do the following operations with this grid:

Choose any subrectangle containing cell (1,1)
, and flip the colors of all its cells. (Flipping means changing its color from white to black or from black to white).

This operation costs 1
 coin.

Choose any subrectangle containing cell (𝑛,1)
, and flip the colors of all its cells.

This operation costs 2
 coins.

Choose any subrectangle containing cell (1,𝑚)
, and flip the colors of all its cells.

This operation costs 4
 coins.

Choose any subrectangle containing cell (𝑛,𝑚)
, and flip the colors of all its cells.

This operation costs 3
 coins.

As a reminder, subrectangle is a set of all cells (𝑥,𝑦)
 with 𝑥1≤𝑥≤𝑥2
, 𝑦1≤𝑦≤𝑦2
 for some 1≤𝑥1≤𝑥2≤𝑛
, 1≤𝑦1≤𝑦2≤𝑚
.

Alice wants to obtain her favorite coloring with these operations. What's the smallest number of coins that she would have to spend? It can be shown that it's always possible to transform the initial grid into any other.


### ideas
1. 操作2和操作3，是多余的，是因为，可以用操作1，模拟操作2/3，且不会付出更多的代价
2. dp[i][j][0] 表示，使用操作1，使得左上角（i, j) 都为1的最优解
3. dp[i][j][1] ... = 1
4. dp[i][j][0] = dp[i][j-1][0] + dp[i-1][j][0] 如果grid[i][j] = 0 
5.             = dp[i][j-1][1] + dp[i-1][j][1] + 1 如果 grid[i][j] = 1