Egor has an array 𝑎
of length 𝑛
, initially consisting of zeros. However, he wanted to turn it into another array 𝑏
of length 𝑛
.

Since Egor doesn't take easy paths, only the following operation can be used (possibly zero or several times):

choose an array 𝑙
of length 𝑘
(1≤𝑙𝑖≤𝑛
, all 𝑙𝑖
are distinct) and change each element 𝑎𝑙𝑖
to 𝑙(𝑖%𝑘)+1
(1≤𝑖≤𝑘
).
He became interested in whether it is possible to get the array 𝑏
using only these operations. Since Egor is still a beginner programmer, he asked you to help him solve this problem.

The operation %
means taking the remainder, that is, 𝑎%𝑏
is equal to the remainder of dividing the number 𝑎
by the number 𝑏
.

Input
The first line of the input contains an integer 𝑡
(1≤𝑡≤105
) - the number of test cases.

Each test case consists of two lines. The first line contains two integers 𝑛
and 𝑘
(1≤𝑘≤𝑛≤105
).

The second line contains the array 𝑏1,𝑏2,…,𝑏𝑛
(1≤𝑏𝑖≤𝑛
).

It is guaranteed that the sum of 𝑛
over all test cases does not exceed 2⋅105
.

### thoughts

1. 考虑如何得到b[n],
2. 假设选择了一个长度为3的l, 其中 l[2] = n
3. a[n] => l[2 % 3 + 1] = l[3] = b[n]
4. 这里就出现了l[1] = ?, l[2] = n, l[3] = b[n]
5. 如果长度为2 l[1] = n, l[2] = b[n]
6. 如果 n != b[n] 那么就可以用长度为2的l进行操作
7. i != b[i] 都可以用 l = [i, b[i]]
8. 剩下的就是 i = b[i]的情况， j = b[j]
9. l = [i] 是不是就可以了？
10. 似乎是可行的