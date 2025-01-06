You are given a two-dimensional maze with a start and end position. Your task is to find the fastest way to get from the start to the end position. The fastest way is to make the minimum number of steps where one step is going left, right, up, or down. Of course, you cannot walk through walls.

There is, however, a catch: If you make more than three steps in the same direction, you lose balance and fall down. Therefore, it is forbidden to make more than three consecutive steps in the same direction. It is okay to walk three times to the right, then one step to the left, and then again three steps to the right. This has the same effect as taking five steps to the right, but is slower.

### ideas
1. dp[i][j][x][d] 表示到达i，j且在x方向已经走了d步时的最小值
2. 