You are given a tree (an undirected connected acyclic graph) consisting of 𝑛
 vertices and 𝑛−1
 edges. A number is written on each edge, each number is either 0
 (let's call such edges 0
-edges) or 1
 (those are 1
-edges).

Let's call an ordered pair of vertices (𝑥,𝑦)
 (𝑥≠𝑦
) valid if, while traversing the simple path from 𝑥
 to 𝑦
, we never go through a 0
-edge after going through a 1
-edge. Your task is to calculate the number of valid pairs in the tree.

Input
The first line contains one integer 𝑛
 (2≤𝑛≤200000
) — the number of vertices in the tree.

Then 𝑛−1
 lines follow, each denoting an edge of the tree. Each edge is represented by three integers 𝑥𝑖
, 𝑦𝑖
 and 𝑐𝑖
 (1≤𝑥𝑖,𝑦𝑖≤𝑛
, 0≤𝑐𝑖≤1
, 𝑥𝑖≠𝑦𝑖
) — the vertices connected by this edge and the number written on it, respectively.

It is guaranteed that the given edges form a tree.



### ideas
1. (x, y) is good if all edges are 0, or 1, or 01, 01?0 is bad
2. dp[u][0] 表示和u通过0边相连的节点的数量，dp[1]表示只有1边相连的， dp[10]表示，先有1，然后0边的节点数量， dp[11] 表示先有0边，再有1边的情况
3. ans = dp[u][0] * dp[v][1] (先用0， 再用1) + dp[u][0] * dp[v][10] + dp[u][1] * dp[v][1]
4.    + dp[u][11] * dp[v][1]  (必须考虑u-v边的情况)