Vasya has a tree consisting of 𝑛
vertices with root in vertex 1
. At first all vertices has 0
written on it.

Let 𝑑(𝑖,𝑗)
be the distance between vertices 𝑖
and 𝑗
, i.e. number of edges in the shortest path from 𝑖
to 𝑗
. Also, let's denote 𝑘
-subtree of vertex 𝑥
— set of vertices 𝑦
such that next two conditions are met:

𝑥
is the ancestor of 𝑦
(each vertex is the ancestor of itself);
𝑑(𝑥,𝑦)≤𝑘
.
Vasya needs you to process 𝑚
queries. The 𝑖
-th query is a triple 𝑣𝑖
, 𝑑𝑖
and 𝑥𝑖
. For each query Vasya adds value 𝑥𝑖
to each vertex from 𝑑𝑖
-subtree of 𝑣𝑖
.

Report to Vasya all values, written on vertices of the tree after processing all queries.

Input
The first line contains single integer 𝑛
(1≤𝑛≤3⋅105
) — number of vertices in the tree.

Each of next 𝑛−1
lines contains two integers 𝑥
and 𝑦
(1≤𝑥,𝑦≤𝑛
) — edge between vertices 𝑥
and 𝑦
. It is guarantied that given graph is a tree.

Next line contains single integer 𝑚
(1≤𝑚≤3⋅105
) — number of queries.

Each of next 𝑚
lines contains three integers 𝑣𝑖
, 𝑑𝑖
, 𝑥𝑖
(1≤𝑣𝑖≤𝑛
, 0≤𝑑𝑖≤109
, 1≤𝑥𝑖≤109
) — description of the 𝑖
-th query.

Output
Print 𝑛
integers. The 𝑖
-th integers is the value, written in the 𝑖
-th vertex after processing all queries.

### ideas

1. 和高度有关系