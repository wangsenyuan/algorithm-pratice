You are given a tree consisting of 𝑛
 vertices. A number is written on each vertex; the number on vertex 𝑖
 is equal to 𝑎𝑖
.

Let's denote the function 𝑔(𝑥,𝑦)
 as the greatest common divisor of the numbers written on the vertices belonging to the simple path from vertex 𝑥
 to vertex 𝑦
 (including these two vertices). Also let's denote 𝑑𝑖𝑠𝑡(𝑥,𝑦)
 as the number of vertices on the simple path between vertices 𝑥
 and 𝑦
, including the endpoints. 𝑑𝑖𝑠𝑡(𝑥,𝑥)=1
 for every vertex 𝑥
.

Your task is calculate the maximum value of 𝑑𝑖𝑠𝑡(𝑥,𝑦)
 among such pairs of vertices that 𝑔(𝑥,𝑦)>1
.

### ideas
1. gcd(x, y) > 1 最大的 dist(x, y)
2. 对于节点u来说, 它的质因数的个数不超过 lgn
3. 而且，如果是通过它的，只要找最大的两个，对于同一个质因数