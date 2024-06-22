You are given a tree with 𝑛
 vertices.

Your task is to find the maximum number 𝑥
 such that it is possible to remove exactly 𝑘
 edges from this tree in such a way that the size of each remaining connected component†
 is at least 𝑥
.

†
 Two vertices 𝑣
 and 𝑢
 are in the same connected component if there exists a sequence of numbers 𝑡1,𝑡2,…,𝑡𝑘
 of arbitrary length 𝑘
, such that 𝑡1=𝑣
, 𝑡𝑘=𝑢
, and for each 𝑖
 from 1
 to 𝑘−1
, vertices 𝑡𝑖
 and 𝑡𝑖+1
 are connected by an edge.


### ideas
1. when k = n - 1, x = 1
2. 且如果能够达到 x, 那么同样 x - 1 也是ok的
3. 所以符合2分的性质，所以可以从底部开始，满足x的时候，删除一条边，看看最后剩下的部分
4. 但是有个问题就是，比如有个子树，它的size > x 但是size < 2 * x，它有3个子树，a,b,c且每一个都 < x
5. 所以要删除它，最好是找个size = x的位置， 似乎只能浪费了
6. dfs不大行，要bfs才行。从外围开始，选中最少的点，这样浪费的少