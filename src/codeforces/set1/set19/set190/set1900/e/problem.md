You are given a directed graph 𝐺
with 𝑛
vertices and 𝑚
edges between them.

Initially, graph 𝐻
is the same as graph 𝐺
. Then you decided to perform the following actions:

If there exists a triple of vertices 𝑎
, 𝑏
, 𝑐
of 𝐻
, such that there is an edge from 𝑎
to 𝑏
and an edge from 𝑏
to 𝑐
, but there is no edge from 𝑎
to 𝑐
, add an edge from 𝑎
to 𝑐
.
Repeat the previous step as long as there are such triples.
Note that the number of edges in 𝐻
can be up to 𝑛2
after performing the actions.

You also wrote some values on vertices of graph 𝐻
. More precisely, vertex 𝑖
has the value of 𝑎𝑖
written on it.

Consider a simple path consisting of 𝑘
distinct vertices with indexes 𝑣1,𝑣2,…,𝑣𝑘
. The length of such a path is 𝑘
. The value of that path is defined as ∑𝑘𝑖=1𝑎𝑣𝑖
.

A simple path is considered the longest if there is no other simple path in the graph with greater length.

Among all the longest simple paths in 𝐻
, find the one with the smallest value.

### thoughts

1. 考虑G的components，每一个都会成为一个完全图
2. 在每个完全图中，最长的路径就是连接所有节点的那条路径
3. 这个不是无向图，是有向图，所有要找强两通分量