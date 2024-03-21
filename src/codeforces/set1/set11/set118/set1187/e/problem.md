You are given a tree (an undirected connected acyclic graph) consisting of 𝑛
vertices. You are playing a game on this tree.

Initially all vertices are white. On the first turn of the game you choose one vertex and paint it black. Then on each
turn you choose a white vertex adjacent (connected by an edge) to any black vertex and paint it black.

Each time when you choose a vertex (even during the first turn), you gain the number of points equal to the size of the
connected component consisting only of white vertices that contains the chosen vertex. The game ends when all vertices
are painted black.

Vertices 1
and 4
are painted black already. If you choose the vertex 2
, you will gain 4
points for the connected component consisting of vertices 2,3,5
and 6
. If you choose the vertex 9
, you will gain 3
points for the connected component consisting of vertices 7,8
and 9
.

Your task is to maximize the number of points you gain.

### ideas

1. f(u) = 以u为root的子树的score
2. f(u) = sz(u) + sum of f(v) where v is children of u
3. 假设计算出了dp(0) = f(0) (0是root)
4. 如何计算dp(1) (1是0的子节点)
5. f(1) - sz(1) = 已经是节点1的子树的score
6. 增加0变成1的节点时的score即可 = f(0) - sz(0) - f(1) + n - sz(1)
7. 