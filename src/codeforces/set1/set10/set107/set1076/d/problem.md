You are given an undirected connected weighted graph consisting of 𝑛
vertices and 𝑚
edges. Let's denote the length of the shortest path from vertex 1
to vertex 𝑖
as 𝑑𝑖
.

You have to erase some edges of the graph so that at most 𝑘
edges remain. Let's call a vertex 𝑖
good if there still exists a path from 1
to 𝑖
with length 𝑑𝑖
after erasing the edges.

Your goal is to erase the edges in such a way that the number of good vertices is maximized.

Input
The first line contains three integers 𝑛
, 𝑚
and 𝑘
(2≤𝑛≤3⋅105
, 1≤𝑚≤3⋅105
, 𝑛−1≤𝑚
, 0≤𝑘≤𝑚
) — the number of vertices and edges in the graph, and the maximum number of edges that can be retained in the graph,
respectively.

Then 𝑚
lines follow, each containing three integers 𝑥
, 𝑦
, 𝑤
(1≤𝑥,𝑦≤𝑛
, 𝑥≠𝑦
, 1≤𝑤≤109
), denoting an edge connecting vertices 𝑥
and 𝑦
and having weight 𝑤
.

The given graph is connected (any vertex can be reached from any other vertex) and simple (there are no self-loops, and
for each unordered pair of vertices there exists at most one edge connecting these vertices).

Output
In the first line print 𝑒
— the number of edges that should remain in the graph (0≤𝑒≤𝑘
).

In the second line print 𝑒
distinct integers from 1
to 𝑚
— the indices of edges that should remain in the graph. Edges are numbered in the same order they are given in the
input. The number of good vertices should be as large as possible.

### ideas

1. 保留k条边 => 保留 k+1个节点