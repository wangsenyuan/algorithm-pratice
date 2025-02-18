A tree is an undirected connected graph without cycles.

You are given a tree of 𝑛
 vertices. Find the number of ways to choose exactly 𝑘
 vertices in this tree (i. e. a 𝑘
-element subset of vertices) so that all pairwise distances between the selected vertices are equal (in other words, there exists an integer 𝑐
 such that for all 𝑢,𝑣
 (𝑢≠𝑣
, 𝑢,𝑣
 are in selected vertices) 𝑑𝑢,𝑣=𝑐
, where 𝑑𝑢,𝑣
 is the distance from 𝑢
 to 𝑣
).

Since the answer may be very large, you need to output it modulo 109+7
.

### ideas
1. n <= 100
2. 以某个节点x为中心点，在它的x个子树中，选择k个个在相同距离下的节点（按层选择）
3. 还有以两个点的连线为中心的（这个情况似乎只有k=2的时候才是有效的）一边选一个的情况
4. k=2的时候，特殊处理（随便选两个点，它们只有一个距离）
5. 剩下的都是选定一个中心点x，在它的每个分支中，根据距离选择节点(这个过程似乎要dp才行)