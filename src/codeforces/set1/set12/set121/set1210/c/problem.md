Kamil likes streaming the competitive programming videos. His MeTube channel has recently reached 100
 million subscribers. In order to celebrate this, he posted a video with an interesting problem he couldn't solve yet. Can you help him?

You're given a tree — a connected undirected graph consisting of 𝑛
 vertices connected by 𝑛−1
 edges. The tree is rooted at vertex 1
. A vertex 𝑢
 is called an ancestor of 𝑣
 if it lies on the shortest path between the root and 𝑣
. In particular, a vertex is an ancestor of itself.

Each vertex 𝑣
 is assigned its beauty 𝑥𝑣
 — a non-negative integer not larger than 1012
. This allows us to define the beauty of a path. Let 𝑢
 be an ancestor of 𝑣
. Then we define the beauty 𝑓(𝑢,𝑣)
 as the greatest common divisor of the beauties of all vertices on the shortest path between 𝑢
 and 𝑣
. Formally, if 𝑢=𝑡1,𝑡2,𝑡3,…,𝑡𝑘=𝑣
 are the vertices on the shortest path between 𝑢
 and 𝑣
, then 𝑓(𝑢,𝑣)=gcd(𝑥𝑡1,𝑥𝑡2,…,𝑥𝑡𝑘)
. Here, gcd
 denotes the greatest common divisor of a set of numbers. In particular, 𝑓(𝑢,𝑢)=gcd(𝑥𝑢)=𝑥𝑢
.

Your task is to find the sum

∑𝑢 is an ancestor of 𝑣𝑓(𝑢,𝑣).

As the result might be too large, please output it modulo 109+7
.

Note that for each 𝑦
, gcd(0,𝑦)=gcd(𝑦,0)=𝑦
. In particular, gcd(0,0)=0
.

### ideas
1. 考虑一条直线，从左到右， f(i) = x[i]和前面的所有的点的x[j...i]的gcd之和
2. 对于x[i]来说，gcd(x[i], ?)肯定是x[i]的因数，这样的数，最多有lg(x[i])个
3. 假设g(i) = [(d1, c1), (d2, c2), (d3, c3)...]
4. 那么f(i) = d1 * c1 + d2 * c2 + ... d3 * c3
5. 那么到i+1后，g(i+1) = 对这些数字，进行merge就可以了