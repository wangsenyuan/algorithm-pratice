There is an edge-weighted complete binary tree with 𝑛
 leaves. A complete binary tree is defined as a tree where every non-leaf vertex has exactly 2 children. For each non-leaf vertex, we label one of its children as the left child and the other as the right child.

The binary tree has a very strange property. For every non-leaf vertex, one of the edges to its children has weight 0
 while the other edge has weight 1
. Note that the edge with weight 0
 can be connected to either its left or right child.

You forgot what the tree looks like, but luckily, you still remember some information about the leaves in the form of an array 𝑎
 of size 𝑛
. For each 𝑖
 from 1
 to 𝑛
, 𝑎𝑖
 represents the distance†
 from the root to the 𝑖
-th leaf in dfs order‡
. Determine whether there exists a complete binary tree which satisfies array 𝑎
. Note that you do not need to reconstruct the tree.

†
 The distance from vertex 𝑢
 to vertex 𝑣
 is defined as the sum of weights of the edges on the path from vertex 𝑢
 to vertex 𝑣
.

‡
 The dfs order of the leaves is found by calling the following 𝚍𝚏𝚜
 function on the root of the binary tree.


 ### ideas
 1. 假设 n = (1 << ?) - 1， 比如 n = 7
 2. 那么 0, 1, 2, 3 （4个leaf）这几个数字都应该存在
 3. 上面那个条件是不成立的
 4. 首先可以肯定的是0, 肯定存在一个，且只能有一个, 
 5. 它和1组成1对，然后它们变成新的0（这两个节点删掉）