You are given a tree∗
 of 𝑛
 vertices. You must perform the following operation exactly twice.

Select a vertex 𝑣
;
Remove all edges incident to 𝑣
, and also the vertex 𝑣
.
Please find the maximum number of connected components after performing the operation exactly twice.

Two vertices 𝑥
 and 𝑦
 are in the same connected component if and only if there exists a path from 𝑥
 to 𝑦
. For clarity, note that the graph with 0
 vertices has 0
 connected components by definition.†

### ideas
1. 如果只能选一个，那么必然是选择那个deg最大的那个
2. 因为删除deg=x的，将会制造出x个components
3. 所以，选择u以后，应该在剩余的里面选择那些最大的v
4. 这里v存在两种情况，一种是它和u相连，一种是不相连，分类讨论