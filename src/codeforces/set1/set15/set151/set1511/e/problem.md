You have a large rectangular board which is divided into 𝑛×𝑚
 cells (the board has 𝑛
 rows and 𝑚
 columns). Each cell is either white or black.

You paint each white cell either red or blue. Obviously, the number of different ways to paint them is 2𝑤
, where 𝑤
 is the number of white cells.

After painting the white cells of the board, you want to place the maximum number of dominoes on it, according to the following rules:

each domino covers two adjacent cells;
each cell is covered by at most one domino;
if a domino is placed horizontally (it covers two adjacent cells in one of the rows), it should cover only red cells;
if a domino is placed vertically (it covers two adjacent cells in one of the columns), it should cover only blue cells.
Let the value of the board be the maximum number of dominoes you can place. Calculate the sum of values of the board over all 2𝑤
 possible ways to paint it. Since it can be huge, print it modulo 998244353
.

The first line contains two integers 𝑛
 and 𝑚
 (1≤𝑛,𝑚≤3⋅105
; 𝑛𝑚≤3⋅105
) — the number of rows and columns, respectively.

Then 𝑛
 lines follow, each line contains a string of 𝑚
 characters. The 𝑗
-th character in the 𝑖
-th string is * if the 𝑗
-th cell in the 𝑖
-th row is black; otherwise, that character is o.


### ideas
1. 懵逼了
2. 一种方案的value = 在这种方案下，能够使用的多米诺骨牌的最大数量
3. 因为骨牌只能cover相同的颜色，且垂直的时候是红色，水平的时候是蓝色
4. value = 垂直方向，红色对的数量 + 水平方向 蓝色对的数量
5. 而且这个貌似是可以贪心的，就是如果它的上面是红色，且自己是红色，就应该马上放置一个
6. dp[i][j][0] 表示i，j没有被占用时的贡献(i,j是很色，或者前面的被占用了)
7. 如果(i, j)是黑色 dp[i][j][0] = 0, dp[i][j][1] = dp[i-1][j][1] + dp[i][j-1][1]
8. 如果(i, j)是白色, dp[i][j][1] = (dp[i-1][j][0] 如果上面是白色 + dp[i][j-1][0] (如果左边是白色)) * pow(2, w - 2)
9. dp[i][j][0] = dp[i-1][j][1] + dp[i][j-1][1] 