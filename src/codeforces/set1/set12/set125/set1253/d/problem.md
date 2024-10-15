You're given an undirected graph with 𝑛
 nodes and 𝑚
 edges. Nodes are numbered from 1
 to 𝑛
.

The graph is considered harmonious if and only if the following property holds:

For every triple of integers (𝑙,𝑚,𝑟)
 such that 1≤𝑙<𝑚<𝑟≤𝑛
, if there exists a path going from node 𝑙
 to node 𝑟
, then there exists a path going from node 𝑙
 to node 𝑚
.
In other words, in a harmonious graph, if from a node 𝑙
 we can reach a node 𝑟
 through edges (𝑙<𝑟
), then we should able to reach nodes (𝑙+1),(𝑙+2),…,(𝑟−1)
 too.

What is the minimum number of edges we need to add to make the graph harmonious?



### ideas
1. 这个题目有点绕
2. 假设l, r 是一个区间，有连接的最大的区间，在这个区间内，假设有m个component，那么就需要 m - 1 条边，把它们连起来