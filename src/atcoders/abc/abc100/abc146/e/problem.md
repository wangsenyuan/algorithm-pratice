https://atcoder.jp/contests/abc146/tasks/abc146_e

输入 n(1≤n≤2e5) k(1≤k≤1e9) 和长为 n 的数组 a(1≤a[i]≤1e9)。

输出 a 的非空连续子数组 b 的个数，满足 sum(b) % k = len(b)。

输入
5 4
1 4 2 3 5
输出 4

输入
8 4
4 2 4 2 4 2 4 2
输出 7

输入
10 7
14 15 92 65 35 89 79 32 38 46
输出 8

### ideas

1. sum(b) = pref(r) - pref(l)
2. 可以改写为 （pref(r) - pref(l)) % k = r - l
3. (pref(r) - pref(l)) % k = pref(r) % k - pref(l) % k
4. or pref(l) % k - pref(r) % k
5. pref(r) % k - r 和 pref(l) % k - l
6. 分别考虑别pref(r) % k 大的范围和小的范围