You are given an array 𝑎
 of 𝑛
 (𝑛≥2
) positive integers and an integer 𝑝
. Consider an undirected weighted graph of 𝑛
 vertices numbered from 1
 to 𝑛
 for which the edges between the vertices 𝑖
 and 𝑗
 (𝑖<𝑗
) are added in the following manner:

If 𝑔𝑐𝑑(𝑎𝑖,𝑎𝑖+1,𝑎𝑖+2,…,𝑎𝑗)=𝑚𝑖𝑛(𝑎𝑖,𝑎𝑖+1,𝑎𝑖+2,…,𝑎𝑗)
, then there is an edge of weight 𝑚𝑖𝑛(𝑎𝑖,𝑎𝑖+1,𝑎𝑖+2,…,𝑎𝑗)
 between 𝑖
 and 𝑗
.
If 𝑖+1=𝑗
, then there is an edge of weight 𝑝
 between 𝑖
 and 𝑗
.
Here 𝑔𝑐𝑑(𝑥,𝑦,…)
 denotes the greatest common divisor (GCD) of integers 𝑥
, 𝑦
, ....

Note that there could be multiple edges between 𝑖
 and 𝑗
 if both of the above conditions are true, and if both the conditions fail for 𝑖
 and 𝑗
, then there is no edge between these vertices.

The goal is to find the weight of the minimum spanning tree of this graph.

### ideas
1. 如果 gcd(ai, ai+1, ai+2, ... aj) = ax (肯定是最小值)
2. 如果 i != x, 且 i != y，猜想用 ax, 连接 (i, x) (x, j) 是一个可行的方案
3. 考虑在mst中连接到x的时候，假设此时已经链接了i,或者j，那么使用a[x]去链接(x, i)应该是可以的
4. 除非，在i的前面有一个k cost(k, x) = a[k] < a[x]
5. 所以，按照a[i]升序排，
6. 然后把i左右两边的，a[i]的倍数，都是用a[i]的cost加进来
7. 