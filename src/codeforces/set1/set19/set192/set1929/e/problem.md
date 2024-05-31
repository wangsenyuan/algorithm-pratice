Sasha was given a tree†
 with 𝑛
 vertices as a prize for winning yet another competition. However, upon returning home after celebrating his victory, he noticed that some parts of the tree were missing. Sasha remembers that he colored some of the edges of this tree. He is certain that for any of the 𝑘
 pairs of vertices (𝑎1,𝑏1),…,(𝑎𝑘,𝑏𝑘)
, he colored at least one edge on the simple path‡
 between vertices 𝑎𝑖
 and 𝑏𝑖
.

Sasha does not remember how many edges he exactly colored, so he asks you to tell him the minimum number of edges he could have colored to satisfy the above condition.

†
A tree is an undirected connected graph without cycles.

‡
A simple path is a path that passes through each vertex at most once.

Input
Each test consists of multiple test cases. The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (2≤𝑛≤105
) — the number of vertices in the tree.

The next (𝑛−1)
 lines describe the edges of the tree. The 𝑖
-th line contains two integers 𝑢𝑖
 and 𝑣𝑖
 (1≤𝑢𝑖,𝑣𝑖≤𝑛
, 𝑢𝑖≠𝑣𝑖
) — the numbers of the vertices connected by the 𝑖
-th edge.

The next line contains a single integer 𝑘
 (1≤𝑘≤20
) — the number of pairs of vertices between which Sasha colored at least one edge on a simple path.

The next 𝑘
 lines describe pairs. The 𝑗
-th line contains two integers 𝑎𝑗
 and 𝑏𝑗
 (1≤𝑎𝑗,𝑏𝑗≤𝑛,𝑎𝑗≠𝑏𝑗
) — the vertices in the 𝑗
-th pair.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 105
. It is guaranteed that the sum of 2𝑘
 over all test cases does not exceed 220
.

### ideas
1. 对于每条边，是可以知道它被哪些pair所cover的（这个可以用lca来做到）
2. dp[state] 表示要cover住state表示的pair所需的最小的edge的数量
3. 问题出在计算dp[state] 这是因为，不可能去迭代所有的边，这个是 state * n的
4. 如果一对pair，它的lca = u, 那么这条边必然只可能在u的子树中
5. 且如果有一对pair，它的一端在u中，一端在u外面，且它还没有被cover到，那么有很大可能要在u的入边上进行着色
6. 不迭代state，但是迭代edge可以吗？
7. edge的状态为cur (cur > 0) 如果state包含cur, dp[state] = 1 + dp[x] x & (state ^ cur) == (state ^ cur)
8. x要包含state中那些未被cur包含的位
9. x是 (state ^ cur) 的超集，且是state的子集
10. 通过不断添加新的节点进去吗？
11. 完全摸不着头脑。看答案了
12. 关键点，居然是， cover每条边的状态不会超过O(k)...

### solution

Let's consider each edge 𝑖
 and mark the set of pairs 𝑆𝑖
 that it covers. Then the claim is: we have a total of 𝑂(𝑘)
 different sets. This is because we are only interested in the edges that are present in the compressed tree on these 𝑘
 pairs of vertices. And as it is known, the number of edges in the compressed tree is 𝑂(𝑘)
.

Then we need to find the minimum number of sets among these sets such that each pair is present in at least one of them. This can be achieved by dynamic programming on sets as follows:

Let 𝑑𝑝[𝑚𝑎𝑠𝑘]
 be the minimum number of edges that need to be removed in order for at least one edge to be removed among the pairs corresponding to the individual set bits in 𝑚𝑎𝑠𝑘
.

Then the update is as follows: 𝑑𝑝[𝑚𝑎𝑠𝑘|𝑆𝑖]=min(𝑑𝑝[𝑚𝑎𝑠𝑘|𝑆𝑖],𝑑𝑝[𝑚𝑎𝑠𝑘]+1)
 for all distinct sets 𝑆
, where 𝑆𝑖
 is the mask corresponding to the pairs passing through the edge. This update is performed because we are adding one more edge to this mask.

As a result, we obtain a solution with a time complexity of 𝑂(𝑛𝑘+2𝑘𝑘)
, where 𝑂(𝑛𝑘)
 is for precomputing the set of pairs removed by each edge for each edge, and 𝑂(2𝑘𝑘)
 is for updating the dynamic programming.

