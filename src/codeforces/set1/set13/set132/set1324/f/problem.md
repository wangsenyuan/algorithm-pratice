You are given a tree consisting of 𝑛
 vertices. A tree is a connected undirected graph with 𝑛−1
 edges. Each vertex 𝑣
 of this tree has a color assigned to it (𝑎𝑣=1
 if the vertex 𝑣
 is white and 0
 if the vertex 𝑣
 is black).

You have to solve the following problem for each vertex 𝑣
: what is the maximum difference between the number of white and the number of black vertices you can obtain if you choose some subtree of the given tree that contains the vertex 𝑣
? The subtree of the tree is the connected subgraph of the given tree. More formally, if you choose the subtree that contains 𝑐𝑛𝑡𝑤
 white vertices and 𝑐𝑛𝑡𝑏
 black vertices, you have to maximize 𝑐𝑛𝑡𝑤−𝑐𝑛𝑡𝑏
.


### ideas
1. 先考虑节点0，以0为根，任何一个子树，它的black的节点数量>=white的节点数量，都应该考虑把它给舍弃掉
2. 然后汇总就可以了
3. 然后re-root即可