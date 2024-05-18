You are given an array of integers 𝑎1,𝑎2,…,𝑎𝑛
.

A pair of integers (𝑖,𝑗)
, such that 1≤𝑖<𝑗≤𝑛
, is called good, if there does not exist an integer 𝑘
 (1≤𝑘≤𝑛
) such that 𝑎𝑖
 is divisible by 𝑎𝑘
 and 𝑎𝑗
 is divisible by 𝑎𝑘
 at the same time.

Please, find the number of good pairs.




### solution

整体思路
考虑有多少对 (a[i],a[j]) 的 GCD 恰好等于 k，将其记作 res[k]。
如果 a 中没有数能整除 k，我们就可以把 res[k] 加入答案。


如何计算 res[k]
统计 a 中每个元素的出现次数 cnt。
枚举 k 及其倍数（k,2k,3k,...），累加这些数的出现次数，记作 c。
这 c 个数中，任选两个数的 GCD，一定是 k 的倍数。所以 c*(c-1)/2 就是 GCD 等于 k,2k,3k,... 的数对的个数。
但我们要计算的是【恰好】等于 k 的数对个数，所以要减去 GCD 恰好等于 2k,3k,... 的数对个数，得
res[k] = c*(c-1)/2 - res[2k] - res[3k] - ...
这个过程可以用 O(nlogn) 的枚举完成。（调和级数）
注意要倒序枚举 k，因为转移来源是比 k 大的数。

如何计算 a 中是否有整除 k 的数（a 中是否有 k 的因子）
如果 cnt[i] > 0，那么对于 i,2i,3i,... 这些数来说，a 中都有这些数的因子。标记这些数。
这个过程也可以用 O(nlogn) 的枚举完成。

注意使用 64 位整数。
如果觉得代码有些慢，可以加个快读来提速。
