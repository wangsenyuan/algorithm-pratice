You are given a forest of 𝑘
 rooted trees∗
. Lumberjack Timofey wants to cut down the entire forest by applying the following operation:

Select a subtree†
 of any vertex of one of the trees and remove it from the tree.
Timofey loves bitwise operations, so he wants the bitwise OR of the sizes of the subtrees he removed to be maximum. Help him and find the maximum result he can obtain.

∗
 A tree is a connected graph without cycles, loops, or multiple edges. In a rooted tree, a selected vertex is called a root. A forest is a collection of one or more trees.

†
 The subtree of a vertex 𝑣
 is the set of vertices for which 𝑣
 lies on the shortest path from this vertex to the root, including 𝑣
 itself.

Input
Each test consists of multiple test cases. The first line contains an integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Then follows the description of the test cases.

The first line of each test case contains a single integer 𝑘
 (1≤𝑘≤106
) — the number of trees in the forest.

This is followed by a description of each of the 𝑘
 trees:

The first line contains a single integer 𝑛
 (1≤𝑛≤106
) — the size of the tree. The vertices of the tree are numbered with integers from 1
 to 𝑛
. The root of the tree is vertex number 1
.

The second line contains 𝑛−1
 integers 𝑝2,𝑝3,…𝑝𝑛
 (1≤𝑝𝑖<𝑖
), where 𝑝𝑖
 — the parent of vertex 𝑖
.

It is guaranteed that the sum of 𝑘
 and 𝑛
 for all sets of input data does not exceed 106
.

Output
For each test case, output a single integer — the maximum result that can be obtained.

### ideas
1. 考虑只有一棵树，且这个树的size = n时的最大值
2. 