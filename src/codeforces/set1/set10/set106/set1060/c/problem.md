You are given two arrays 𝑎
and 𝑏
of positive integers, with length 𝑛
and 𝑚
respectively.

Let 𝑐
be an 𝑛×𝑚
matrix, where 𝑐𝑖,𝑗=𝑎𝑖⋅𝑏𝑗
.

You need to find a subrectangle of the matrix 𝑐
such that the sum of its elements is at most 𝑥
, and its area (the total number of elements) is the largest possible.

Formally, you need to find the largest number 𝑠
such that it is possible to choose integers 𝑥1,𝑥2,𝑦1,𝑦2
subject to 1≤𝑥1≤𝑥2≤𝑛
, 1≤𝑦1≤𝑦2≤𝑚
, (𝑥2−𝑥1+1)×(𝑦2−𝑦1+1)=𝑠
, and
∑𝑖=𝑥1𝑥2∑𝑗=𝑦1𝑦2𝑐𝑖,𝑗≤𝑥.

### ideas

1. n * m
2. 假设选中了a的一段[x, y], b的一段[u, v]
3. 那么对于a[i]，它的贡献 = a[i] * (sb[v] - sb[u-1])
4. 那么整体的贡献 = (sa[y] - sa[x-1]) * (sb[v] - sb[u-1])
5. 那么对于固定的 sa[y] - sa[x-1] 越长越好
6. 所以，可以对sa[y] - sa[x-1] (y - x + 1)进行排序, 使用一个升序队列，sum越大，长度越大