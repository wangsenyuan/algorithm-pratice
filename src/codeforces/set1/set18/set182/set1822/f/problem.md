Two friends, Alisa and Yuki, planted a tree with 𝑛
vertices in their garden. A tree is an undirected graph without cycles, loops, or multiple edges. Each edge in this tree
has a length of 𝑘
. Initially, vertex 1
is the root of the tree.

Alisa and Yuki are growing the tree not just for fun, they want to sell it. The cost of the tree is defined as the
maximum distance from the root to a vertex among all vertices of the tree. The distance between two vertices 𝑢
and 𝑣
is the sum of the lengths of the edges on the path from 𝑢
to 𝑣
.

The girls took a course in gardening, so they know how to modify the tree. Alisa and Yuki can spend 𝑐
coins to shift the root of the tree to one of the neighbors of the current root. This operation can be performed any
number of times (possibly zero). Note that the structure of the tree is left unchanged; the only change is which vertex
is the root.

The friends want to sell the tree with the maximum profit. The profit is defined as the difference between the cost of
the tree and the total cost of operations. The profit is cost of the tree minus the total cost of operations.

Help the girls and find the maximum profit they can get by applying operations to the tree any number of times (possibly
zero).

### thoughts

1. 对于任何一个root，计算它的最大收益，cost也是固定的
2. 计算完u后，移动到v再计算