You are given an undirected tree consisting of 𝑛
 vertices. An undirected tree is a connected undirected graph with 𝑛−1
 edges.

Your task is to add the minimum number of edges in such a way that the length of the shortest path from the vertex 1
 to any other vertex is at most 2
. Note that you are not allowed to add loops and multiple edges.

### ideas
1. dp[u][0] 表示将u和1节点直接相连时，在子树u中的最优解
2. dp[u][1] 表示将u的父节点和1相连时，在子树u中的最优解
3. dp[u][2] 表示将u的祖父节点和1相连时的最优解
4. dp[u][0] = sum of (min(dp[v][1], 1 + dp[v][0]), v is direct child of u)
5. dp[u][1] = sum of (min(dp[v][2], 1 + dp[v][0]))
6. dp[u][2] = sum of (dp[v][0] + 1)
7. 