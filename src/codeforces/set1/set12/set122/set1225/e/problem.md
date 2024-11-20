You are at the top left cell (1,1)
 of an 𝑛×𝑚
 labyrinth. Your goal is to get to the bottom right cell (𝑛,𝑚)
. You can only move right or down, one cell per step. Moving right from a cell (𝑥,𝑦)
 takes you to the cell (𝑥,𝑦+1)
, while moving down takes you to the cell (𝑥+1,𝑦)
.

Some cells of the labyrinth contain rocks. When you move to a cell with rock, the rock is pushed to the next cell in the direction you're moving. If the next cell contains a rock, it gets pushed further, and so on.

The labyrinth is surrounded by impenetrable walls, thus any move that would put you or any rock outside of the labyrinth is illegal.

Count the number of different legal paths you can take from the start to the goal modulo 109+7
. Two paths are considered different if there is at least one cell that is visited in one path, but not visited in the other.



### ideas
1. 如果没有rock， 那就相当简单了nCr(n + m - 1, n-1)
2. 有rock咋处理呢？dp[i][j] = dp[i+1][j] (如果 i+1有一块石头) ？
3. 如果可以把这块石头push下去。但是问题是，有可能已经push了几块石头下来