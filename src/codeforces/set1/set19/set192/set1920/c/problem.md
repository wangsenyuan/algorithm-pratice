Allen has an array 𝑎1,𝑎2,…,𝑎𝑛
. For every positive integer 𝑘
 that is a divisor of 𝑛
, Allen does the following:

He partitions the array into 𝑛𝑘
 disjoint subarrays of length 𝑘
. In other words, he partitions the array into the following subarrays:
[𝑎1,𝑎2,…,𝑎𝑘],[𝑎𝑘+1,𝑎𝑘+2,…,𝑎2𝑘],…,[𝑎𝑛−𝑘+1,𝑎𝑛−𝑘+2,…,𝑎𝑛]
Allen earns one point if there exists some positive integer 𝑚
 (𝑚≥2
) such that if he replaces every element in the array with its remainder when divided by 𝑚
, then all subarrays will be identical.
Help Allen find the number of points he will earn.


### ideas
1. k is easy to find, but m not
2. a1 % m = ak % m
3. 假设a1 != ak
4. a % m = b % m => (a - b) % m = 0
5. 那么也就是说 m 是 gcd(ak - a1, ...) 的一个除数
6. 至少有了一个 n * n 的算法
7. 还有一个观察就是，如果k 成立，那么k的倍数，也是成立的，只需要使用相同的m即可
8. 但是如果k不成立，那么k的倍数一定不成立吗？好像不是的，
9. [1, 0, 1, 1, 3, 2, 3, 3]
10. 这里k = 2, 显然不成立，但是k = 4是成立的
11. a[i] <= n 这个信息有没有用呢？
12. 那么 a[i] - a[j] <= n - 1, 也就是说， m的范围似乎比较小
13. m会不会一定是2？
14. 也不一定
15. [1, 2, 3, 4, 5, 6]
16. 如果m= 2, 那么当k=3 [1, 0, 1, 0, 1, 0] 不满足条件
17. 但是m = 3, [1, 2, 0, 1, 2, 0] 是满足条件的
18. 