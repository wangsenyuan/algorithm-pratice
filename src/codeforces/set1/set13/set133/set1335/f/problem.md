There is a rectangular grid of size 𝑛×𝑚
. Each cell of the grid is colored black ('0') or white ('1'). The color of the cell (𝑖,𝑗)
 is 𝑐𝑖,𝑗
. You are also given a map of directions: for each cell, there is a direction 𝑠𝑖,𝑗
 which is one of the four characters 'U', 'R', 'D' and 'L'.

If 𝑠𝑖,𝑗
 is 'U' then there is a transition from the cell (𝑖,𝑗)
 to the cell (𝑖−1,𝑗)
;
if 𝑠𝑖,𝑗
 is 'R' then there is a transition from the cell (𝑖,𝑗)
 to the cell (𝑖,𝑗+1)
;
if 𝑠𝑖,𝑗
 is 'D' then there is a transition from the cell (𝑖,𝑗)
 to the cell (𝑖+1,𝑗)
;
if 𝑠𝑖,𝑗
 is 'L' then there is a transition from the cell (𝑖,𝑗)
 to the cell (𝑖,𝑗−1)
.
It is guaranteed that the top row doesn't contain characters 'U', the bottom row doesn't contain characters 'D', the leftmost column doesn't contain characters 'L' and the rightmost column doesn't contain characters 'R'.

You want to place some robots in this field (at most one robot in a cell). The following conditions should be satisfied.

Firstly, each robot should move every time (i.e. it cannot skip the move). During one move each robot goes to the adjacent cell depending on the current direction.
Secondly, you have to place robots in such a way that there is no move before which two different robots occupy the same cell (it also means that you cannot place two robots in the same cell). I.e. if the grid is "RL" (one row, two columns, colors does not matter there) then you can place two robots in cells (1,1)
 and (1,2)
, but if the grid is "RLL" then you cannot place robots in cells (1,1)
 and (1,3)
 because during the first second both robots will occupy the cell (1,2)
.
The robots make an infinite number of moves.

Your task is to place the maximum number of robots to satisfy all the conditions described above and among all such ways, you have to choose one where the number of black cells occupied by robots before all movements is the maximum possible. Note that you can place robots only before all movements.

### ideas
1. R?L这类的只能在一个地方放，更进一步的，如果(u, v) 和 (x, y) 在同一时刻到位置(i, j)
2. 它们不能同时放置
3. 如果把鸽子按照棋盘进行划分，那么显然，只有同色的才有可能会出现相撞的情况
4. 不是同色的不会；
5. 从任何一个格子开始，反向向外运动，任何同一层的（同色的）只能放置一个（不需要考虑同色，不是同色的不会在同一层）
6. 但是这个有问题，从哪里开始，貌似是会影响到结果的
7. 如果 a[i][j] = 'R' (这里向右运动)，那么dp[i][j+1] = dp[i][j] + 1 (dp[i][j]表示层级)
8. 但是似乎不大对，因为是存在环的
9. 貌似所有的都是最后都进入了一个环（或者在一个环上）
10. 能够产生冲突的点，就在于那些有多个分支进入的地方
11. 一旦进入一个环，貌似能够容纳的最多的robot，就是这个环的长度
12. 能够放置的就是环上面格子的数量，但是黑色的robot的数量不是
13. 能够进入环的所有格子，都需要考虑
14. 然后以这个环中，任何一个格子开始计算，给所有能到达环的格子，计算它们到这里的距离
15. 在相同距离上，取黑色的部分
16. 考虑一个长度为2的圈，那么距离为偶数的那些，就是冲突的