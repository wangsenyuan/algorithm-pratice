This is the hard version of the problem. The difference between the versions is the constraint on 𝑛
 and the required number of operations. You can make hacks only if all versions of the problem are solved.

There are two binary strings 𝑎
 and 𝑏
 of length 𝑛
 (a binary string is a string consisting of symbols 0
 and 1
). In an operation, you select a prefix of 𝑎
, and simultaneously invert the bits in the prefix (0
 changes to 1
 and 1
 changes to 0
) and reverse the order of the bits in the prefix.

For example, if 𝑎=001011
 and you select the prefix of length 3
, it becomes 011011
. Then if you select the entire string, it becomes 001001
.

Your task is to transform the string 𝑎
 into 𝑏
 in at most 2𝑛
 operations. It can be proved that it is always possible.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤1000
)  — the number of test cases. Next 3𝑡
 lines contain descriptions of test cases.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤105
)  — the length of the binary strings.

The next two lines contain two binary strings 𝑎
 and 𝑏
 of length 𝑛
.

It is guaranteed that the sum of 𝑛
 across all test cases does not exceed 105
.

Output
For each test case, output an integer 𝑘
 (0≤𝑘≤2𝑛
), followed by 𝑘
 integers 𝑝1,…,𝑝𝑘
 (1≤𝑝𝑖≤𝑛
). Here 𝑘
 is the number of operations you use and 𝑝𝑖
 is the length of the prefix you flip in the 𝑖
-th operation.

### ideas
1. 感觉得从后往前处理。如果最后一位是相同的，即a[n] = b[n], 往前一位
2. 先从正面考虑， 如果在i处处理过了，那么a[1...i]的都取反了（先不考虑reverse的情况）
3. 假设在f(i) = 到i为止的操作次数，
4. 如果在i进行了一次取反，那么 i-1前面的所有数，多取反了一次，a[i]取反了一次
5. 也就是不考虑reverse的情况，a[i]的取反次数 等于 后缀和
6. 那么这样也就决定了a[i]的值是什么
7. 但是有办法，让a[i] = b[i]吗？ 也就是让a[i]回到i位置，且和b[i]的值相同
8. 如果a[i] = b[i] => go on
9. 如果 a = "0", b = "1" =》 操作一次即可
10. 如果 a = "00", b = "11", i = 2
11. a = "01", b = "10", (i = 2 不行， 变回了a) 
12.   i = 1, a = "11"
13.   i = 2, a = "00"
14.   i = 1, a = "10"
15. 如果a是全0，b[i] = 1, 那么可以先选择j = i, 此时a[1:i]全1，然后再选择 j = i - 1 (或者前一个0的位置)，就又变成了全0的前缀
16. 那怎么把a全变成0呢？
17. 如果a[i] = 1, 那么就是选择j= i-1, 然后选择后一个为0的前一个的位置（或者是n）
18. 这样子就可以得到全0（这个操作数，不会操作n次）
19. 操作最多的情况是010101这样交错的字符串
20. 1234
21. 2134
22. 3124
23. 4213