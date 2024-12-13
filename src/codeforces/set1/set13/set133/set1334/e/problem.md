You are given a positive integer 𝐷
. Let's build the following graph from it:

each vertex is a divisor of 𝐷
 (not necessarily prime, 1
 and 𝐷
 itself are also included);
two vertices 𝑥
 and 𝑦
 (𝑥>𝑦
) have an undirected edge between them if 𝑥
 is divisible by 𝑦
 and 𝑥𝑦
 is a prime;
the weight of an edge is the number of divisors of 𝑥
 that are not divisors of 𝑦
.
For example, here is the graph for 𝐷=12
:


Edge (4,12)
 has weight 3
 because 12
 has divisors [1,2,3,4,6,12]
 and 4
 has divisors [1,2,4]
. Thus, there are 3
 divisors of 12
 that are not divisors of 4
 — [3,6,12]
.

There is no edge between 3
 and 2
 because 3
 is not divisible by 2
. There is no edge between 12
 and 3
 because 123=4
 is not a prime.

Let the length of the path between some vertices 𝑣
 and 𝑢
 in the graph be the total weight of edges on it. For example, path [(1,2),(2,6),(6,12),(12,4),(4,2),(2,6)]
 has length 1+2+2+3+1+2=11
. The empty path has length 0
.

So the shortest path between two vertices 𝑣
 and 𝑢
 is the path that has the minimal possible length.

Two paths 𝑎
 and 𝑏
 are different if there is either a different number of edges in them or there is a position 𝑖
 such that 𝑎𝑖
 and 𝑏𝑖
 are different edges.

You are given 𝑞
 queries of the following form:

𝑣
 𝑢
 — calculate the number of the shortest paths between vertices 𝑣
 and 𝑢
.
The answer for each query might be large so print it modulo 998244353
.

### ideas
1. 考虑u和v的边的weight
2. u / v = p (一个prime数)， 假设v的 divisors的数是已知的，
3. 那么u-v的weight = 那些 v % (x * p) != 0 的部分
4. 比如 (12, 6) [1, 2, 3, 4, 6, 12] vs (1, 2, 3，6)
5. (4, 12) (也就是2, 6的倍数)
6. 也就是那些因数x，v / x 的部分，不包含p了
7. 假设u的质因数，这里找出来，v的质因数也找出
8. gcd(u, v)的质因数找出来
9. 但是怎么算边的数量呢？
10. 