You are given a rooted tree, consisting of 𝑛
 vertices. The vertices in the tree are numbered from 1
 to 𝑛
, and the root is the vertex 1
. Let 𝑑𝑥
 be the distance (the number of edges on the shortest path) from the root to the vertex 𝑥
.

There is a chip that is initially placed at the root. You can perform the following operation as many times as you want (possibly zero):

move the chip from the current vertex 𝑣
 to a vertex 𝑢
 such that 𝑑𝑢=𝑑𝑣+1
. If 𝑣
 is the root, you can choose any vertex 𝑢
 meeting this constraint; however, if 𝑣
 is not the root, 𝑢
 should not be a neighbor of 𝑣
 (there should be no edge connecting 𝑣
 and 𝑢
).

For example, in the tree above, the following chip moves are possible: 1→2
, 1→5
, 2→7
, 5→3
, 5→4
, 3→6
, 7→6
.

A sequence of vertices is valid if you can move the chip in such a way that it visits all vertices from the sequence (and only them), in the order they are given in the sequence.

Your task is to calculate the number of valid vertex sequences. Since the answer might be large, print it modulo 998244353
.

### ideas
1. dp[u] 是到u的有效计数, dp[u] = sum(dp[v]) - dp[p] 就是到它的上一层，且排除父节点的有效计数
2. 