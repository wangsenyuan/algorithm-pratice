Given three numbers 𝑛,𝑎,𝑏
. You need to find an adjacency matrix of such an undirected graph that the number of components in it is equal to 𝑎
, and the number of components in its complement is 𝑏
. The matrix must be symmetric, and all digits on the main diagonal must be zeroes.

In an undirected graph loops (edges from a vertex to itself) are not allowed. It can be at most one edge between a pair of vertices.

The adjacency matrix of an undirected graph is a square matrix of size 𝑛
 consisting only of "0" and "1", where 𝑛
 is the number of vertices of the graph and the 𝑖
-th row and the 𝑖
-th column correspond to the 𝑖
-th vertex of the graph. The cell (𝑖,𝑗)
 of the adjacency matrix contains 1
 if and only if the 𝑖
-th and 𝑗
-th vertices in the graph are connected by an edge.

A connected component is a set of vertices 𝑋
 such that for every two vertices from this set there exists at least one path in the graph connecting this pair of vertices, but adding any other vertex to 𝑋
 violates this rule.

The complement or inverse of a graph 𝐺
 is a graph 𝐻
 on the same vertices such that two distinct vertices of 𝐻
 are adjacent if and only if they are not adjacent in 𝐺
.

