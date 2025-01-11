Let's call an undirected graph 𝐺=(𝑉,𝐸)
 relatively prime if and only if for each edge (𝑣,𝑢)∈𝐸
  𝐺𝐶𝐷(𝑣,𝑢)=1
 (the greatest common divisor of 𝑣
 and 𝑢
 is 1
). If there is no edge between some pair of vertices 𝑣
 and 𝑢
 then the value of 𝐺𝐶𝐷(𝑣,𝑢)
 doesn't matter. The vertices are numbered from 1
 to |𝑉|
.

Construct a relatively prime graph with 𝑛
 vertices and 𝑚
 edges such that it is connected and it contains neither self-loops nor multiple edges.

If there exists no valid graph with the given number of vertices and edges then output "Impossible".

If there are multiple answers then print any of them.

### ideas
1. gcd(u, v) = 1 = u - 1 - 欧拉数？
2. 如果 sum(u...) = m就有解，否则就无解
3. 