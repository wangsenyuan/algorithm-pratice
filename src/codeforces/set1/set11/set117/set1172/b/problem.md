Nauuo is a girl who loves drawing circles.

One day she has drawn a circle and wanted to draw a tree on it.

The tree is a connected undirected graph consisting of 𝑛
 nodes and 𝑛−1
 edges. The nodes are numbered from 1
 to 𝑛
.

Nauuo wants to draw a tree on the circle, the nodes of the tree should be in 𝑛
 distinct points on the circle, and the edges should be straight without crossing each other.

"Without crossing each other" means that every two edges have no common point or the only common point is an endpoint of both edges.

Nauuo wants to draw the tree using a permutation of 𝑛
 elements. A permutation of 𝑛
 elements is a sequence of integers 𝑝1,𝑝2,…,𝑝𝑛
 in which every integer from 1
 to 𝑛
 appears exactly once.

After a permutation is chosen Nauuo draws the 𝑖
-th node in the 𝑝𝑖
-th point on the circle, then draws the edges connecting the nodes.

The tree is given, Nauuo wants to know how many permutations are there so that the tree drawn satisfies the rule (the edges are straight without crossing each other). She only wants to know the answer modulo 998244353
, can you help her?

It is obvious that whether a permutation is valid or not does not depend on which 𝑛
 points on the circle are chosen.



### ideas
1. 没想法～
2. 如果 p[u] <-> p[v] 之间有一条线，且u, v 中间的某个x和v后面的某个y有线，那么这是无效的
3. 反过来讲, u....(v....w) 间有连线， 那么 v....w可以看作是u的子树，且v和w的子树之间不能有线
4. 感觉和 dfs order 有关系？
5. 一颗子树的位置，必须是连续的
6. 假设有x颗子树，那么有 x!个选择（这些子树的顺序决定了占位）