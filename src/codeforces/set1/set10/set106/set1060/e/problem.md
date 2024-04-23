Sergey Semyonovich is a mayor of a county city N and he used to spend his days and nights in thoughts of further
improvements of Nkers' lives. Unfortunately for him, anything and everything has been done already, and there are no
more possible improvements he can think of during the day (he now prefers to sleep at night). However, his assistants
have found a solution and they now draw an imaginary city on a paper sheet and suggest the mayor can propose its
improvements.

Right now he has a map of some imaginary city with 𝑛
subway stations. Some stations are directly connected with tunnels in such a way that the whole map is a tree (
assistants were short on time and enthusiasm). It means that there exists exactly one simple path between each pair of
station. We call a path simple if it uses each tunnel no more than once.

One of Sergey Semyonovich's favorite quality objectives is the sum of all pairwise distances between every pair of
stations. The distance between two stations is the minimum possible number of tunnels on a path between them.

Sergey Semyonovich decided to add new tunnels to the subway map. In particular, he connected any two stations 𝑢
and 𝑣
that were not connected with a direct tunnel but share a common neighbor, i.e. there exists such a station 𝑤
that the original map has a tunnel between 𝑢
and 𝑤
and a tunnel between 𝑤
and 𝑣
. You are given a task to compute the sum of pairwise distances between all pairs of stations in the new map.

### ideas

1. 如果只是tree，那么应该是比较简单的，就是通过dfs tree，然后更改root，利用之前的信息进行计算
2. 现在距离为2的被连起来了，要怎么计算呢？首先，对于节点x，所有距离等于1的的没有变化，但是距离>1的，都减小了1
3. 所以，可以在原来的基础上，进行2次计算？
4. 不对， 考虑 a -> b -> c -> d -> e
5. a 到e的距离变成了2 (a -> c -> e)， 而不是3
6. 看起来是距离x是奇数的部分变成了（比如 b, e) 变成了 dist/2 + 1
7. 考虑root = 0， 如何计算到0的距离之和？
8. 假设到0的距离为1的节点有cnt[1]个，距离为2的节点有cnt[2]个。。。
9. 那么结果就等于 1 * cnt[1] + 2/2 * cnt[2] + (3 + 1) / 2 * cnt[3]....
10. 统一的形式就是 sum ((i + 1)) / 2 * cnt[i]
11. dp[u] = 表示子树中到连接到u的最小值
12. dp[u] = sum of dp[v] (v 是u的child)
13. 对于v来说，所有的距离 + 1, 但是cnt没有变化
14. dp[u] = sum of [i + 1] / 2 * cnt[i] (i 表示的是到v的距离)
15. 如果i是奇数 [i+1] / 2 == [i]/2  (这里是向上取整)
16. 如果i是偶数 [i+1] / 2 = [i]/2 + 1
17. 所以 dp[u] = dp[v] + sum of cnt[i] where i is even (i 从0开始) 