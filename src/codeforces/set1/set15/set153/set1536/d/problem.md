Uh oh! Ray lost his array yet again! However, Omkar might be able to help because he thinks he has found the OmkArray of Ray's array. The OmkArray of an array 𝑎
 with elements 𝑎1,𝑎2,…,𝑎2𝑘−1
, is the array 𝑏
 with elements 𝑏1,𝑏2,…,𝑏𝑘
 such that 𝑏𝑖
 is equal to the median of 𝑎1,𝑎2,…,𝑎2𝑖−1
 for all 𝑖
. Omkar has found an array 𝑏
 of size 𝑛
 (1≤𝑛≤2⋅105
, −109≤𝑏𝑖≤109
). Given this array 𝑏
, Ray wants to test Omkar's claim and see if 𝑏
 actually is an OmkArray of some array 𝑎
. Can you help Ray?

The median of a set of numbers 𝑎1,𝑎2,…,𝑎2𝑖−1
 is the number 𝑐𝑖
 where 𝑐1,𝑐2,…,𝑐2𝑖−1
 represents 𝑎1,𝑎2,…,𝑎2𝑖−1
 sorted in nondecreasing order.

Input
Each test contains multiple test cases. The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Description of the test cases follows.

The first line of each test case contains an integer 𝑛
 (1≤𝑛≤2⋅105
) — the length of the array 𝑏
.

The second line contains 𝑛
 integers 𝑏1,𝑏2,…,𝑏𝑛
 (−109≤𝑏𝑖≤109
) — the elements of 𝑏
.

It is guaranteed the sum of 𝑛
 across all test cases does not exceed 2⋅105
.

Output
For each test case, output one line containing YES if there exists an array 𝑎
 such that 𝑏𝑖
 is the median of 𝑎1,𝑎2,…,𝑎2𝑖−1
 for all 𝑖
, and NO otherwise. The case of letters in YES and NO do not matter (so yEs and No will also be accepted).

### ideas
1. b[1] = a[1]
2. b[2] = median of [a[1], a[2], a[3]]
3. b[3] = median of [a[1], a[2], a[3], a[4], a[5]]
4. 如果 b[1] > b[2], b[2]肯定不能等于a[1], 假设b[2] = a[2], 那么，比然有a[1] > a[2] >= a[3]
5.     b[1] < b[2]  => a[1] < a[2] <= a[3]
6.  如果b[1] != b[2] => 肯定增加了两个同侧 （大于小于b[1])的数
7.  如果b[1] = b[2] => 肯定在两侧各加了大于(小于)b[1]的数
8.  b[i+1] > b[i]的时候，增加两个 大于b[i]的数，且其中一个必须是 b[i+1]吗？另外一个数 >= b[i+1]即可
9.  b[i+1] = b[i]的时候，不妨就增加两个b[i]即可
10. 这个不大对， 因为这样子，假设后面需要一个比b[i]小的数；
11. 如果在i的时候，产生了两个数, 一个比b[i]小，一个比b[i]大，而后面那个数，就可以理由到这个比b[i]小的数
12. 否则，就是比b[i]大的数了
13. 也就是说，假设现在添加了一个小于b[i]的数，那么这个数事实上，可以起作用的范围是[0...b[i]]
14. 那这个计数，要怎么利用呢？
15. 假设处理b[i], 希望有 2 * (i-1) <= b[i]的数，和 2 * (i-1)个更大的数
16. 那假设时候，发现有超过 2 * (i - 1)个数 <= b[i], 似乎是没有关系的，因为那些多余的数，是不是可以移动的？
17. 所以，这个是range update + range query
18. 不大对
19. 如果 b[i] => b[i+1], 且 b[i] < b[i+1], 
20. 那么必然是添加了两个>= b[i+1]的数, 且在b[i]和b[i+1]中间没有其他的数
21. 当b[i] = b[i+1]时，就是在b[i+1]的上下，各添加了一个数
22. 且我们可以认为set(a) = set(b)
23. 考虑b[n], 那么需要 2 * (n-1)个数 <= b[n], 2 * (n - 1)个数 >= b[n]