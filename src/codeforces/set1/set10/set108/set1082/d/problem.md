Graph constructive problems are back! This time the graph you are asked to build should match the following properties.

The graph is connected if and only if there exists a path between every pair of vertices.

The diameter (aka "longest shortest path") of a connected undirected graph is the maximum number of edges in the
shortest path between any pair of its vertices.

The degree of a vertex is the number of edges incident to it.

Given a sequence of 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
construct a connected undirected graph of 𝑛
vertices such that:

the graph contains no self-loops and no multiple edges;
the degree 𝑑𝑖
of the 𝑖
-th vertex doesn't exceed 𝑎𝑖
(i.e. 𝑑𝑖≤𝑎𝑖
);
the diameter of the graph is maximum possible.
Output the resulting graph or report that no solution exists.

### ideas

1. 如果 ai 都 等于 1 => -1
2. a[i] = 1 的只能是叶子节点，然后所有 > 1的为内部节点，
3. 然后在最大的a[i]上把叶子节点加上去，直到不能增加，第二个节点
4. 但是这样子，会有种可能性，是没法把所有1都加上去 =》 -1