You are given an array 𝑎
 of length 𝑛
 and 𝑞
 queries 𝑙
, 𝑟
.

For each query, find the maximum possible 𝑚
, such that all elements 𝑎𝑙
, 𝑎𝑙+1
, ..., 𝑎𝑟
 are equal modulo 𝑚
. In other words, 𝑎𝑙mod𝑚=𝑎𝑙+1mod𝑚=⋯=𝑎𝑟mod𝑚
, where 𝑎mod𝑏
 — is the remainder of division 𝑎
 by 𝑏
. In particular, when 𝑚
 can be infinite, print 0
.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

The first line of each test case contains two integers 𝑛
, 𝑞
 (1≤𝑛,𝑞≤2⋅105
) — the length of the array and the number of queries.

The second line of each test case contains 𝑛
 integers 𝑎𝑖
 (1≤𝑎𝑖≤109
) — the elements of the array.

In the following 𝑞
 lines of each test case, two integers 𝑙
, 𝑟
 are provided (1≤𝑙≤𝑟≤𝑛
) — the range of the query.

It is guaranteed that the sum of 𝑛
 across all test cases does not exceed 2⋅105
, and the sum of 𝑞
 does not exceed 2⋅105
.

### ideas
1. 如果l...r 都相等 => 0
2. 否则的话，1 肯定是满足取模后相等的条件
3. 更进一步，他们的gcd也是满足条件的
4. 那是不是就是gcd呢？
5. 好像不是，比如 14, 2, 当 m = 4时，也是ok的
6. 假设 x = m * k + r
7. y = m * l + r
8. 且 x < y
9. (y - x) % m = 0
10.  如果是x, y, z 呢
11.  x = m * x1 + r
12.  y = m * y1 + r
13.  z = m * z1 + r
14.  (y - z) % m = 0
15.  (x - z) % m = 0
16.  (x - y) % m = 0
17.  也就是它们的差的绝对值的gcd（连续的情况下是不是也是一样的）