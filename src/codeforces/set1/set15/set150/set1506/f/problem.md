Consider an infinite triangle made up of layers. Let's number the layers, starting from one, from the top of the triangle (from top to bottom). The 𝑘
-th layer of the triangle contains 𝑘
 points, numbered from left to right. Each point of an infinite triangle is described by a pair of numbers (𝑟,𝑐)
 (1≤𝑐≤𝑟
), where 𝑟
 is the number of the layer, and 𝑐
 is the number of the point in the layer. From each point (𝑟,𝑐)
 there are two directed edges to the points (𝑟+1,𝑐)
 and (𝑟+1,𝑐+1)
, but only one of the edges is activated. If 𝑟+𝑐
 is even, then the edge to the point (𝑟+1,𝑐)
 is activated, otherwise the edge to the point (𝑟+1,𝑐+1)
 is activated. Look at the picture for a better understanding.

Activated edges are colored in black. Non-activated edges are colored in gray.
From the point (𝑟1,𝑐1)
 it is possible to reach the point (𝑟2,𝑐2)
, if there is a path between them only from activated edges. For example, in the picture above, there is a path from (1,1)
 to (3,2)
, but there is no path from (2,1)
 to (1,1)
.

Initially, you are at the point (1,1)
. For each turn, you can:

Replace activated edge for point (𝑟,𝑐)
. That is if the edge to the point (𝑟+1,𝑐)
 is activated, then instead of it, the edge to the point (𝑟+1,𝑐+1)
 becomes activated, otherwise if the edge to the point (𝑟+1,𝑐+1)
, then instead if it, the edge to the point (𝑟+1,𝑐)
 becomes activated. This action increases the cost of the path by 1
;
Move from the current point to another by following the activated edge. This action does not increase the cost of the path.
You are given a sequence of 𝑛
 points of an infinite triangle (𝑟1,𝑐1),(𝑟2,𝑐2),…,(𝑟𝑛,𝑐𝑛)
. Find the minimum cost path from (1,1)
, passing through all 𝑛
 points in arbitrary order.

 ### ideas
 1. 首先路径，要按照r排序，且r不可能重复（否则就没法访问完所有的点）
 2. 然后，两个点之间的结果，不影响其他的部分。所以，可以每段单独处理
 3. 然后对于(a, b) (c, d) 
 4. 假设c-a 的距离是 3， 也就是差了3层，再假设b = 1
 5. 那么 b变成2，以后，如果不付出额外的cost，就必须一直保留成2
 6. 如果 d > b => d / 2 - b / 2
 7. 如果 d < b， 这个是不可能的（因为它没法到前面去）
 8. 搞错了， 是看 r + c 的奇偶性的
 9. 如果 r + c 是偶数 (r + 1, c)
 10. 如果r+c是奇数, => (r + 1, c + 1)