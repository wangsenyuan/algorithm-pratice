You are given a grid with 𝑛
 rows and 𝑚
 columns of non-negative integers and an integer 𝑘
. Let (𝑖,𝑗)
 denote the cell in the 𝑖
-th row from the top and 𝑗
-th column from the left (1≤𝑖≤𝑛
, 1≤𝑗≤𝑚
). For every cell (𝑖,𝑗)
, the integer 𝑎𝑖,𝑗
 is written on the cell (𝑖,𝑗)
.

You are initially at (1,1)
 and want to go to (𝑛,𝑚)
. You may only move down or right. That is, if you are at (𝑖,𝑗)
, you can only move to (𝑖+1,𝑗)
 or (𝑖,𝑗+1)
 (if the corresponding cell exists).

Before you begin moving, you may do the following operation any number of times:

Choose an integer 𝑖
 between 1
 and 𝑛
 and cyclically shift row 𝑖
 to the left by 1
. Formally, simultaneously set 𝑎𝑖,𝑗
 to 𝑎𝑖,(𝑗mod𝑚)+1
 for all integers 𝑗
 (1≤𝑗≤𝑚
).
Note that you may not do any operation after you start moving.
After moving from (1,1)
 to (𝑛,𝑚)
, let 𝑥
 be the number of operations you have performed before moving, and let 𝑦
 be the sum of the integers written on visited cells (including (1,1)
 and (𝑛,𝑚)
). Then the cost is defined as 𝑘𝑥+𝑦
.

Find the minimum cost to move from (1,1)
 to (𝑛,𝑚)
.

### ideas
1. dp[i][j] 表示到达（i，j）时的最大值。
2. 然后在(i, j)处支付一个k的成本，可以得到 a[i][j-1] - a[i][j]的收益
