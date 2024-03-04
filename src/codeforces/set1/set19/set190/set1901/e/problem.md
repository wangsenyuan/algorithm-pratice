You are given a tree consisting of 𝑛
vertices. A number is written on each vertex; the number on vertex 𝑖
is equal to 𝑎𝑖
.

You can perform the following operation any number of times (possibly zero):

choose a vertex which has at most 1
incident edge and remove this vertex from the tree.
Note that you can delete all vertices.

After all operations are done, you're compressing the tree. The compression process is done as follows. While there is a
vertex having exactly 2
incident edges in the tree, perform the following operation:

delete this vertex, connect its neighbors with an edge.
It can be shown that if there are multiple ways to choose a vertex to delete during the compression process, the
resulting tree is still the same.

Your task is to calculate the maximum possible sum of numbers written on vertices after applying the aforementioned
operation any number of times, and then compressing the tree.

### thoughts

1. 删除外围子树后，剩下的树再进行压缩
2. 压缩后的树中，不存在度数 = 2的节点
3. 然后计算最后的树的sum
4. 求能得到的最大的sum
5. 考虑删除外围子树后，剩下的树，所有的叶子节点都会被计入最后的sum
6. 假设dp[u] = 某个子树（不考虑压缩时）的贡献， dp[u] = max(0, sum(dp[children])) + a[u]
7. 0 表示把它的子树都删除掉，只剩u作为一个叶子
8. 假设考虑压缩的情况，那么u如果要保留下来，那么必须至少有两个children
9. 将dp[children]排序后，找最大值
10. 如果它只有一个child v, max(0, dp[v]) (它肯定会被排除掉)
11. 如此，好像可以先压缩？

### solution

We can use dynamic programming on tree to solve this problem. After we're done removing vertices, a vertex will be left
in the tree after the compression process if and only if its degree is not 2
. So, we can try to choose the number of children for each vertex in dynamic programming, and depending on this number
of children, the vertex we're considering is either removed from the tree during the compression, or left in the tree.

But if we want to use dynamic programming, we have to root the tree at some vertex, so the degree of each vertex depends
not only on the number of its children we left in the tree, but also on whether there exists an edge leading from the
ancestor to that vertex. So, for each vertex, we will consider two values of dynamic programming: the best answer for
its subtree if there is an edge leading into it from the parent (i. e. we haven't deleted everything outside of that
subtree), and the answer if there is no edge leading from the parent (i. e. we deleted everything outside of that
subtree).

Let the first of these two values be 𝑑𝑝𝑣
— the maximum answer for subtree of 𝑣
if there is at least one non-deleted vertex not in the subtree of vertex 𝑣
(i.e there is an edge from vertex 𝑣
up the tree).

Let's look at the dp transitions, depending on the number of children of vertex 𝑣
that are not deleted:

0
: 𝑑𝑝𝑣=max(𝑑𝑝𝑣,𝑎𝑣)
;
1
: 𝑑𝑝𝑣=max(𝑑𝑝𝑣,max𝑢𝑑𝑝𝑢)
(𝑎𝑣
is not taken into account, because 𝑣
has 2
incident edges and will be compressed);
at least 2
: 𝑑𝑝𝑣=max(𝑑𝑝𝑣,𝑎𝑣+∑𝑢𝑑𝑝𝑢)
.
Then let us consider the case when the resulting tree is in the subtree of some vertex 𝑣
. We can update the global answer depending on the number of children of vertex 𝑣
that are not deleted:

0
: 𝑎𝑛𝑠=max(𝑎𝑛𝑠,𝑎𝑣)
;
1
: 𝑎𝑛𝑠=max(𝑎𝑛𝑠,𝑎𝑣+max𝑢𝑑𝑝𝑢)
;
2
: 𝑎𝑛𝑠=max(𝑎𝑛𝑠,max𝑢1,𝑢2𝑑𝑝𝑢1+𝑑𝑝𝑢2)
(𝑎𝑣
is not taken into account, because 𝑣
has 2
incident edges and will be compressed);
at least 3
: 𝑎𝑛𝑠=max(𝑎𝑛𝑠,𝑎𝑣+∑𝑢𝑑𝑝𝑢)
.
For convenience, we can calculate auxiliary dynamic programming for children of vertex 𝑣
: 𝑠𝑢𝑚𝑖
is the maximum sum of 𝑑𝑝𝑢
for 𝑖
children. From above written transitions, we can see that, 𝑠𝑢𝑚3
can store maximum sum for at least 3
children.

Don't forget that we can delete all vertices, so the answer is at least 0
.