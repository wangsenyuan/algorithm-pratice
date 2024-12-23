You are given three non-negative integers 𝑏
, 𝑐
, and 𝑑
.

Please find a non-negative integer 𝑎∈[0,261]
 such that (𝑎|𝑏)−(𝑎&𝑐)=𝑑
, where |
 and &
 denote the bitwise OR operation and the bitwise AND operation, respectively.

If such an 𝑎
 exists, print its value. If there is no solution, print a single integer −1
. If there are multiple solutions, print any of them.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤105
). The description of the test cases follows.

The only line of each test case contains three positive integers 𝑏
, 𝑐
, and 𝑑
 (0≤𝑏,𝑐,𝑑≤1018
).

Output
For each test case, output the value of 𝑎
, or −1
 if there is no solution. Please note that 𝑎
 must be non-negative and cannot exceed 261
.

### ideas
1. (a | b) - (a & c) = d
2. a & c = a + c - (a | c)
3. (a | b) + (a | c) - a - c = d
4. (a | b) + (a | c) - a = c + d = e
5. a | b >= b
6. c + d = ... <= b + c - a
7. b, c, d = 4 2 6
8. 2 + 6 <= 4 + 2 - a 不成立（-1)
9. 从低往高位处理,如果 b[i] = 0, c[i] = 0, 且 e[i] = 0 (a[i] = 0)
10. 如果 b[i] = 0, c[i] = 0, 且 e[i] = 1, a[i] = 1 (两个1 - 一个1)
11.  如果 b[i] = 0, c[i] = 1, e[i] = 0 (a[i] = 0, 似乎也不行(1 - 0), a[i] = 1 也不行，2个1 - 一个1)
12.  如果发生借位是可以的
13.  
方法一

从集合的角度看，a|b 是 a 的超集，a&c 是 a 的子集。
所以 a|b 一定是 a&c 的超集。
所以不可能出现同一个比特位上，a|b 是 0 而 a&c 是 1 的情况。
这意味着减法是没有【借位】的，所以每一位互相独立，我们可以逐位分析。

逐位分析（从高到低或者从低到高都可以）：
如果 d 这一位是 1，那么必须是 1 - 0 = 1。
1. 如果 b 这一位是 0 且 c 这一位是 1，那么没有这样的 a，输出 -1。
2. 否则，如果 b 这一位是 0，那么 a 这一位必须是 1。（注意此时 c 这一位是 0）
如果 d 这一位是 0，那么可以是 1 - 1 = 0 或者 0 - 0 = 0。
1. 如果 b 这一位是 1 且 c 这一位是 0，那么没有这样的 a，输出 -1。
2. 否则，如果 b 这一位是 1（说明 c 这一位是 1），那么 a 这一位必须是 1。
