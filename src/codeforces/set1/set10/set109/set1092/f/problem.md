You are given a tree consisting exactly of 𝑛
 vertices. Tree is a connected undirected graph with 𝑛−1
 edges. Each vertex 𝑣
 of this tree has a value 𝑎𝑣
 assigned to it.

Let 𝑑𝑖𝑠𝑡(𝑥,𝑦)
 be the distance between the vertices 𝑥
 and 𝑦
. The distance between the vertices is the number of edges on the simple path between them.

Let's define the cost of the tree as the following value: firstly, let's fix some vertex of the tree. Let it be 𝑣
. Then the cost of the tree is ∑𝑖=1𝑛𝑑𝑖𝑠𝑡(𝑖,𝑣)⋅𝑎𝑖
.

Your task is to calculate the maximum possible cost of the tree if you can choose 𝑣
 arbitrarily.



 ### ideas
 1. dist(i, v) = dist(i) * a[i] 以v为root
 2. dp[u] += a[v] + (dist[i] + 1) * a[i] 