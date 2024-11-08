Assume that you have 𝑘
 one-dimensional segments 𝑠1,𝑠2,…𝑠𝑘
 (each segment is denoted by two integers — its endpoints). Then you can build the following graph on these segments. The graph consists of 𝑘
 vertexes, and there is an edge between the 𝑖
-th and the 𝑗
-th vertexes (𝑖≠𝑗
) if and only if the segments 𝑠𝑖
 and 𝑠𝑗
 intersect (there exists at least one point that belongs to both of them).

For example, if 𝑠1=[1,6],𝑠2=[8,20],𝑠3=[4,10],𝑠4=[2,13],𝑠5=[17,18]
, then the resulting graph is the following:


A tree of size 𝑚
 is good if it is possible to choose 𝑚
 one-dimensional segments so that the graph built on these segments coincides with this tree.

You are given a tree, you have to find its good subtree with maximum possible size. Recall that a subtree is a connected subgraph of a tree.

Note that you have to answer 𝑞
 independent queries.

### ideas
1. 有点乱～
2. 这是一棵树
3. 对于每个节点u，考虑3个状态， dp[u][0] 表示它的左右两边还是空的情况
4. dp[u][1]表示，有一边区间重合了
5. dp[u][2] 表示，右边两边区间和u的区间重合了
6. 如果u只有一个子树 v, 那么 dp[u][0] = 1 + dp[v][0] 完全和v重合
7. dp[u][1] = 1 + max(dp[v][0], dp[v][1])
8. dp[u][2] = 1 + max(dp[v][0], dp[v][2])
9. 
