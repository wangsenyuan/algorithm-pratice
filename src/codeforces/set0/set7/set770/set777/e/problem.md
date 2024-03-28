Of course you have heard the famous task about Hanoi Towers, but did you know that there is a special factory producing
the rings for this wonderful game? Once upon a time, the ruler of the ancient Egypt ordered the workers of Hanoi Factory
to create as high tower as possible. They were not ready to serve such a strange order so they had to create this new
tower using already produced rings.

There are n rings in factory's stock. The i-th ring has inner radius ai, outer radius bi and height hi. The goal is to
select some subset of rings and arrange them such that the following conditions are satisfied:

Outer radiuses form a non-increasing sequence, i.e. one can put the j-th ring on the i-th ring only if bj ≤ bi.
Rings should not fall one into the the other. That means one can place ring j on the ring i only if bj>ai.
The total height of all rings used should be maximum possible.

### ideas

1. n <= 100000 (n * n will TLE)
2. sort by b[i]（倒序） first, 在相同b[i]时，按照a[i]倒序排（a[i]越小，对后面的j越有利）
3. 假设dp[i]=以i为顶的tower的高度
4. dp[j] = 前面那些 a[i] < b[j] 的dp的最大值
5. 这里要对a[i]离散化（压缩）处理