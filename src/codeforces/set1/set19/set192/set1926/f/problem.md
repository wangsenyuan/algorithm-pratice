Vladislav has a grid of size 7×7
, where each cell is colored black or white. In one operation, he can choose any cell and change its color (black ↔
 white).

Find the minimum number of operations required to ensure that there are no black cells with four diagonal neighbors also being black.

### ideas
1. dp[state] = 前14个个位置的状态码，对应的最优值