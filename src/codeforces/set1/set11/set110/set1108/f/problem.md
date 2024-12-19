You are given an undirected weighted connected graph with 𝑛
 vertices and 𝑚
 edges without loops and multiple edges.

The 𝑖
-th edge is 𝑒𝑖=(𝑢𝑖,𝑣𝑖,𝑤𝑖)
; the distance between vertices 𝑢𝑖
 and 𝑣𝑖
 along the edge 𝑒𝑖
 is 𝑤𝑖
 (1≤𝑤𝑖
). The graph is connected, i. e. for any pair of vertices, there is at least one path between them consisting only of edges of the given graph.

A minimum spanning tree (MST) in case of positive weights is a subset of the edges of a connected weighted undirected graph that connects all the vertices together and has minimum total cost among all such subsets (total cost is the sum of costs of chosen edges).

You can modify the given graph. The only operation you can perform is the following: increase the weight of some edge by 1
. You can increase the weight of each edge multiple (possibly, zero) times.

Suppose that the initial MST cost is 𝑘
. Your problem is to increase weights of some edges with minimum possible number of operations in such a way that the cost of MST in the obtained graph remains 𝑘
, but MST is unique (it means that there is only one way to choose MST in the obtained graph).

Your problem is to calculate the minimum number of operations required to do it.

### ideas
1. 没想法呐～～～
2. 有些边是必须选的（有些边是可选的）；必须选的那些边被改变，会增加k（所以不能动）
3. 那些可选边
4. 可选边是不是只需要增加一次？好像是的（增加一次后，它肯定不会再被选中，也不会被作为新的candidate使用）