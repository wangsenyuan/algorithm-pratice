You are given two connected undirected graphs with the same number of vertices. In both graphs, there is a token located at some vertex. In the first graph, the token is initially at vertex 𝑠1
, and in the second graph, the token is initially at vertex 𝑠2
. The following operation is repeated an infinite number of times:

Let the token currently be at vertex 𝑣1
 in the first graph and at vertex 𝑣2
 in the second graph.
A vertex 𝑢1
, adjacent to 𝑣1
, is chosen in the first graph.
A vertex 𝑢2
, adjacent to 𝑣2
, is chosen in the second graph.
The tokens are moved to the chosen vertices: in the first graph, the token moves from 𝑣1
 to 𝑢1
, and in the second graph, from 𝑣2
 to 𝑢2
.
The cost of such an operation is equal to |𝑢1−𝑢2|
.
Determine the minimum possible total cost of all operations or report that this value will be infinitely large.


### ideas
1. 如果在某个阶段，出现了一个位置，接下来，在两个图上面，它们的序列一致，那么接下来就可以达到0
2. 假设v1 = v2, 且 u1 = u2, 那么就ok了（可以再返回再继续）
3. 就是说，找到那些i, 且i在两个图中都有相同的邻居(?)
4. 如果不存在这样的点 => -1
5. 如果存在这样的点，如果一旦到达这样的点 =》就是到到这样的点时的最小距离
6. dp[i][j] 表示，到达状态（i，j)的最短距离
7. dp[i][i] (i是特殊的点)
8. 状态怎么迁移？
9. i的邻居 * j的邻居？ 似乎是 n * n 的复杂性呐
10. 距离的奇偶性比较重要
11. dp[i][0] 表示在g1中，到达i，且距离为偶数时的最短距离
12. d1[i][0] & d2[i][0] 
13. 是求cost。。。
14. 