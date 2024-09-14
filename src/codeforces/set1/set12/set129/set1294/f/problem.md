You are given an unweighted tree with 𝑛
 vertices. Recall that a tree is a connected undirected graph without cycles.

Your task is to choose three distinct vertices 𝑎,𝑏,𝑐
 on this tree such that the number of edges which belong to at least one of the simple paths between 𝑎
 and 𝑏
, 𝑏
 and 𝑐
, or 𝑎
 and 𝑐
 is the maximum possible. See the notes section for a better understanding.

The simple path is the path that visits each vertex at most once.

In the first line print one integer 𝑟𝑒𝑠
 — the maximum number of edges which belong to at least one of the simple paths between 𝑎
 and 𝑏
, 𝑏
 and 𝑐
, or 𝑎
 and 𝑐
.


### ideas
1. 考虑直径是不是肯定在答案中？
2. 假设(a, b)是直径，但是不在答案中；
3. 这里分三种情况，（a,b) 两点都不在，只有a不在
4. 如果（a,b)都不在，假设答案是(u, v, w) 如果这三点的路径不经过直径，那么显然将直径替换进去，会得到更优的结果
5. 如果有部分重叠（x, y) 那么 (u -> ... -> x) -> .. -> (y -> ... -> v) 肯定小于 (a -> .. -> x) -> .. -> (y -> .. -> b)
6. 使用(a, b）替换是更优的选择
7. 考虑b在其中，但是a不在其中（u, v, b)
8. 如果（a, b)路径不在其中，那么使用（a, b)是更优的选择（似乎也不可能出现这种情况，因为b是一个叶子节点）
9. 考虑(a, b)的部分路径出现了， 这个点是c, 那么此时用(a, c) 去替换(u, c)是更好的选择
10. 这里证明了，在最优解中，肯定存在直径上的亮点(a, b)
11. 那么剩下就是迭代其他的叶子节点（c), 看哪个提供最优解即可
12. 找到c和直径的交点，然后计算这个长度即可（bfs）即可