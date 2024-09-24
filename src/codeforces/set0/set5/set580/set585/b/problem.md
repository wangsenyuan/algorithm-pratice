The mobile application store has a new game called "Subway Roller".

The protagonist of the game Philip is located in one end of the tunnel and wants to get out of the other one. The tunnel is a rectangular field consisting of three rows and n columns. At the beginning of the game the hero is in some cell of the leftmost column. Some number of trains rides towards the hero. Each train consists of two or more neighbouring cells in some row of the field.

All trains are moving from right to left at a speed of two cells per second, and the hero runs from left to right at the speed of one cell per second. For simplicity, the game is implemented so that the hero and the trains move in turns. First, the hero moves one cell to the right, then one square up or down, or stays idle. Then all the trains move twice simultaneously one cell to the left. Thus, in one move, Philip definitely makes a move to the right and can move up or down. If at any point, Philip is in the same cell with a train, he loses. If the train reaches the left column, it continues to move as before, leaving the tunnel.

Your task is to answer the question whether there is a sequence of movements of Philip, such that he would be able to get to the rightmost column.

### ideas
1. 总花费的时间是n-1, 必须往右移动一格
2. dp[i][j] = true, 表示在时刻i(也就是第i列时)，是否处于行j
3. dp[i][j] = (dp[i-1][j] || dp[i-1][j+1] || dp[i-1][j-1] ) and it is safe in 2 seconds if in road
4.  