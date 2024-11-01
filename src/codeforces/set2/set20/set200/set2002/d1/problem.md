You are given a perfect binary tree†
 consisting of 𝑛
 vertices. The vertices are numbered from 1
 to 𝑛
, and the root is the vertex 1
. You are also given a permutation 𝑝1,𝑝2,…,𝑝𝑛
 of [1,2,…,𝑛]
.

You need to answer 𝑞
 queries. For each query, you are given two integers 𝑥
, 𝑦
; you need to swap 𝑝𝑥
 and 𝑝𝑦
 and determine if 𝑝1,𝑝2,…,𝑝𝑛
 is a valid DFS order‡
 of the given tree.

Please note that the swaps are persistent through queries.

†
 A perfect binary tree is a tree with vertex 1
 as its root, with size 𝑛=2𝑘−1
 for a positive integer 𝑘
, and where the parent of each vertex 𝑖
 (1<𝑖≤𝑛
) is ⌊𝑖2⌋
. Thus, all leaves of this tree are at a distance 𝑘−1
 from the root.

‡
 A DFS order is found by calling the following 𝚍𝚏𝚜
 function on the given tree.

 ### ideas
 1. 怎么确定是 dfs order的？
 2. 对于所有的节点x，它的所有的子节点都在它的内部
 3. 先生成一个dfs order + in, out
 4. 然后对于一开始的p，记录它们最小的in，最大的out（先不管咋搞）
 5. 然后能够得到一些bad/good的点
 6. bad就是那些不满足dfs order的父节点（有子节点在它的外部了）
 7. 交换完以后，就可以去更新bad/good的点
 8. 首先它们自己会被影响到，然后就是它们的父节点
 9. 好像可以搞