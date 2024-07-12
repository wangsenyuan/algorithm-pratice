You are given a sequence of 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
.

Does there exist a sequence of 𝑛
 integers 𝑏1,𝑏2,…,𝑏𝑛
 such that the following property holds?

For each 1≤𝑖≤𝑛
, there exist two (not necessarily distinct) indices 𝑗
 and 𝑘
 (1≤𝑗,𝑘≤𝑛
) such that 𝑎𝑖=𝑏𝑗−𝑏𝑘
.
Input
The first line contains a single integer 𝑡
 (1≤𝑡≤20
) — the number of test cases. Then 𝑡
 test cases follow.

The first line of each test case contains one integer 𝑛
 (1≤𝑛≤10
).

The second line of each test case contains the 𝑛
 integers 𝑎1,…,𝑎𝑛
 (−105≤𝑎𝑖≤105
).


### ideas
1. 无从下手
2. 如果n=1, a[1] = b[1] - b[1] (只有这一种情况) 
3. n = 2, a[1] = b[1] - b[1] 或者 b[1] - b[2], 或者 b[2] - b[1]
4. a[2] = b[2] - b[2] 或者 b[1] - b[2], 或者 b[1] - b[2]
5. a[i] = 0的情况可以忽略掉，因为它们肯定是存在的
6. 但是会多出来一些free的节点，按需要设置值
7. 如果 a[i1] = -a[i2], 那么它们两个可以互反，也就是，只要 b[j] - b[k] = a[i1], 那么 b[k] - b[j] = a[i2]
8. 这样子的，它们只需要2
9. a[i1] + a[i2] + a[i3] = 0 
10. b[j1] - b[k1] + b[j2] - b[k2] + b[j3] - b[k3] = 0
11. 那么不妨设k1 = j2, k2 = j3, k3 = j1 => 组成了一个环
12. 如果是4个呢？
13. b[j1] - b[k1] + b[j2] - b[k2] + b[j3] - b[k3] + b[j4] - b[k4] = 0
14. 也正好需要4个即可，确定了一个数后，其他的也确定了
15. 所以，要把这n个数，分成和 = 0的多组（或者说，它们的就必须是0）？
16. 例子1就不成立
```
4 -7 -1 5 10

𝑎1=4=2−(−2)=𝑏2−𝑏5
;
𝑎2=−7=−9−(−2)=𝑏1−𝑏5
;
𝑎3=−1=1−2=𝑏3−𝑏2
;
𝑎4=5=3−(−2)=𝑏4−𝑏5
;
𝑎5=10=1−(−9)=𝑏3−𝑏1
.

a5 + a2 = (b3 - b1) + (b1 - b5) = b3 - b5
a3 + a1 = (b3 - b2) + (b2 - b5) = b3 - b5
a4 = b4 - b5


``` 

1. a[1] = b[2] - b[n], a[2] = b[2] - b[n] ... a[i] = b[i] - b[n]
2. 除了 a[n] 无法给定 如果 a[n] = 0, 则很容易
3. 