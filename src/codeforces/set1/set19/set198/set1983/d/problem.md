Given two arrays of distinct positive integers 𝑎
 and 𝑏
 of length 𝑛
, we would like to make both the arrays the same. Two arrays 𝑥
 and 𝑦
 of length 𝑘
 are said to be the same when for all 1≤𝑖≤𝑘
, 𝑥𝑖=𝑦𝑖
.

Now in one move, you can choose some index 𝑙
 and 𝑟
 in 𝑎
 (𝑙≤𝑟
) and swap 𝑎𝑙
 and 𝑎𝑟
, then choose some 𝑝
 and 𝑞
 (𝑝≤𝑞
) in 𝑏
 such that 𝑟−𝑙=𝑞−𝑝
 and swap 𝑏𝑝
 and 𝑏𝑞
.

Is it possible to make both arrays the same?

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤2⋅104
). The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤105
) — the length of the arrays 𝑎
 and 𝑏
.

The second line of each test case contains 𝑛
 distinct integers 𝑎1,𝑎2,𝑎3,…,𝑎𝑛
 (1≤𝑎𝑖≤2⋅105
) — the integers in the array 𝑎
.

The third line of each test case contains 𝑛
 distinct integers 𝑏1,𝑏2,𝑏3,…,𝑏𝑛
 (1≤𝑏𝑖≤2⋅105
) — the integers in the array 𝑏
.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 105
.

### ideas
1. 目标是让 a = b （distinct）
2. 所以，至少sorted(a) = sorted(b)
3. 单独考虑a，如果把a中的一些元素进行排序，那么这些元素本身会形成一个环
4. 在环之间进行交换，似乎是没有意义的
5. 在a里面存在一个环，那么在b中也应该有一个和其对应的环
6. 如果这两个环完全一样，当然没有问题。这两个环，应该是一样的，或者shift后，应该是一样的? 
7. 例子 [1, 2, 3, 4] [4, 3, 2, 1]不符合上面的模式
8. 为了让b中的4到达最后的位置, l = 1, 但是 r = 4, 是不行的，这样子会造成a中顺序无法保证
9. (1, 3) 得到 [3, 2, 1, 4] [2, 3, 4, 1]
10. 然后选择 （3, 4) (1, 2) 就可以了，如果一个环的长度是4，就可以这么弄
11. 如果一个环的长度是2 => no
12. 如果一个环的长度是3 => no
13. (a, b, c) (b, a, c)
14. 如果选择([1, 2], [2, 3]) (b, a, c) (b, c, a)
15. 不ok
16. 所以看起来这个环的长度至少得是4
17. 4的倍数，是否ok呢？也可以的吧～