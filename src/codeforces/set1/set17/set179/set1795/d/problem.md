You are given an undirected graph consisting of 𝑛
vertices and 𝑛
edges, where 𝑛
is divisible by 6
. Each edge has a weight, which is a positive (greater than zero) integer.

The graph has the following structure: it is split into 𝑛3
triples of vertices, the first triple consisting of vertices 1,2,3
, the second triple consisting of vertices 4,5,6
, and so on. Every pair of vertices from the same triple is connected by an edge. There are no edges between vertices
from different triples.

You have to paint the vertices of this graph into two colors, red and blue. Each vertex should have exactly one color,
there should be exactly 𝑛2
red vertices and 𝑛2
blue vertices. The coloring is called valid if it meets these constraints.

The weight of the coloring is the sum of weights of edges connecting two vertices with different colors.

Let 𝑊
be the maximum possible weight of a valid coloring. Calculate the number of valid colorings with weight 𝑊
, and print it modulo 998244353
.

### ideas

1. 如果一个三角形里面有一条最大的边，那么它的两个端点的颜色必须是相反的
2. 如果有两条边，或者三条边最大，只要保证每个三角内有red/blue即可
3. 第一种三角形挑出来，假设有m个，那么还需要 n/2 - m个红色，n/2 - m个蓝色
4. (2 ^^ m) * ((n/2 - m) * 2, (n/2 - m)) 即可