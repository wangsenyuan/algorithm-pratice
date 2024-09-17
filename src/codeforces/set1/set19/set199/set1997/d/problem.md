You are given a rooted tree, consisting of 𝑛
 vertices. The vertices in the tree are numbered from 1
 to 𝑛
, and the root is the vertex 1
. The value 𝑎𝑖
 is written at the 𝑖
-th vertex.

You can perform the following operation any number of times (possibly zero): choose a vertex 𝑣
 which has at least one child; increase 𝑎𝑣
 by 1
; and decrease 𝑎𝑢
 by 1
 for all vertices 𝑢
 that are in the subtree of 𝑣
 (except 𝑣
 itself). However, after each operation, the values on all vertices should be non-negative.

Your task is to calculate the maximum possible value written at the root using the aforementioned operation.