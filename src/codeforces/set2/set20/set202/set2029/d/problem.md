You are given an undirected graph with 𝑛
 vertices and 𝑚
 edges.

You can perform the following operation at most 2⋅max(𝑛,𝑚)
 times:

Choose three distinct vertices 𝑎
, 𝑏
, and 𝑐
, then for each of the edges (𝑎,𝑏)
, (𝑏,𝑐)
, and (𝑐,𝑎)
, do the following:
If the edge does not exist, add it. On the contrary, if it exists, remove it.
A graph is called cool if and only if one of the following holds:

The graph has no edges, or
The graph is a tree.
You have to make the graph cool by performing the above operations. Note that you can use at most 2⋅max(𝑛,𝑚)
 operations.

It can be shown that there always exists at least one solution.

### ideas
1. union-find, 如果没有圈，那么就是一棵树 =》 good
2. 如果遇到圈里 ，那么把这个圈上的所有的边，都可以删除掉（和边的数量是相同的）
3. 比如(a, b) 造成了一个环，选择(a, b, fa(b)) 可以把b给移除出来（变成一个独立的点）
4. b = fa(b), 然后一直进行，直到剩下3个点
5. 如果最后的结果里面，既存在数，又存在点，然后把它们merge起来就可以了
6. 