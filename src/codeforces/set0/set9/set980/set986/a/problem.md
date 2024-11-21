Some company is going to hold a fair in Byteland. There are n
 towns in Byteland and m
 two-way roads between towns. Of course, you can reach any town from any other town using roads.

There are k
 types of goods produced in Byteland and every town produces only one type. To hold a fair you have to bring at least s
 different types of goods. It costs d(u,v)
 coins to bring goods from town u
 to town v
 where d(u,v)
 is the length of the shortest path from u
 to v
. Length of a path is the number of roads in this path.

The organizers will cover all travel expenses but they can choose the towns to bring goods from. Now they want to calculate minimum expenses to hold a fair in each of n
 towns.

 ### ideas
 1. 共有k种货物，至少需要s种，对每个城市需要计算最小的cost
 2. dp[u][i] 表示要把货物i运送到城市u的最短距离
 3. 这个可以用bfs计算（从有货物i的城市出发）