The fox Yae climbed the tree of the Sacred Sakura. A tree is a connected undirected graph that does not contain cycles.

The fox uses her magical powers to move around the tree. Yae can jump from vertex 𝑣
to another vertex 𝑢
if and only if the distance between these vertices does not exceed 2
. In other words, in one jump Yae can jump from vertex 𝑣
to vertex 𝑢
if vertices 𝑣
and 𝑢
are connected by an edge, or if there exists such vertex 𝑤
that vertices 𝑣
and 𝑤
are connected by an edge, and also vertices 𝑢
and 𝑤
are connected by an edge.

After Yae was able to get the sakura petal, she wondered if there was a cyclic route in the tree 𝑣1,𝑣2,…,𝑣𝑛
such that:

the fox can jump from vertex 𝑣𝑖
to vertex 𝑣𝑖+1
,
the fox can jump from vertex 𝑣𝑛
to vertex 𝑣1
,
all 𝑣𝑖
are pairwise distinct.
Help the fox determine if the required traversal exists.

Input
The first line contains one integer 𝑛
(2≤𝑛≤2⋅105
) —the number of vertices of the tree.

Each of the following 𝑛−1
lines contains two integers 𝑢
and 𝑣
(1≤𝑢,𝑣≤𝑛
, 𝑢≠𝑣
) — vertices connected by an edge. It is guaranteed that these edges form a tree.

Output
On the first line, print "Yes" (without quotes) if the required route of the tree exists, or "No" (without quotes)
otherwise.

If the required tree traversal exists, on the second line print 𝑛
integers of different integers 𝑣1,𝑣2,…,𝑣𝑛
(1≤𝑣𝑖≤𝑛
) — the vertices of the tree in traversal order.

If there are several correct traversals, output any of them.

### thoughts

1. 完全没想法
2. 假设通过u进入子树u，那么只能从u的某个child v离开子树u
3. 但是是否仍然需要进入子树u呢？（到达某个子节点w)
4. 如果存在某个合理的路径，必然存在w....x的一条路径，然后离开u
5. 那么通过将v -> w (它们是可达的，因为它们的距离必然是2)
6. 可以在完全访问完u的子树后，从x离开
7. 离开时后，貌似只能到达u的父节点p（其他的节点，都是u的子节点）
8. dp[p][0] 表示p越过p后，进入某个节点u
9. dp[u][1] 表示从u进入子树u后是否能到达p
10. dp[p][0] = some dp[u][1] && other (dp[v][0])
11. 越过p后进入u，返回p后，再越过v，然后返回v（不能有其他节点），然后从v离开，到达fa[p]
12. dp[u][1] = dp[x][0] & dp[y][1]
13. 越过某个子节点x，然后返回到x，并进入y，但y不能有子节点，从y离开
14. dp[u][1] = x表示落脚掉