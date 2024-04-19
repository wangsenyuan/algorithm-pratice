Vasya has two arrays 𝐴
and 𝐵
of lengths 𝑛
and 𝑚
, respectively.

He can perform the following operation arbitrary number of times (possibly zero): he takes some consecutive subsegment
of the array and replaces it with a single element, equal to the sum of all elements on this subsegment. For example,
from the array [1,10,100,1000,10000]
Vasya can obtain array [1,1110,10000]
, and from array [1,2,3]
Vasya can obtain array [6]
.

Two arrays 𝐴
and 𝐵
are considered equal if and only if they have the same length and for each valid 𝑖
𝐴𝑖=𝐵𝑖
.

Vasya wants to perform some of these operations on array 𝐴
, some on array 𝐵
, in such a way that arrays 𝐴
and 𝐵
become equal. Moreover, the lengths of the resulting arrays should be maximal possible.

Help Vasya to determine the maximum length of the arrays that he can achieve or output that it is impossible to make
arrays 𝐴
and 𝐵
equal.

Input
The first line contains a single integer 𝑛 (1≤𝑛≤3⋅105)
— the length of the first array.

The second line contains 𝑛
integers 𝑎1,𝑎2,⋯,𝑎𝑛 (1≤𝑎𝑖≤109)
— elements of the array 𝐴
.

The third line contains a single integer 𝑚 (1≤𝑚≤3⋅105)
— the length of the second array.

The fourth line contains 𝑚
integers 𝑏1,𝑏2,⋯,𝑏𝑚 (1≤𝑏𝑖≤109)

- elements of the array 𝐵
  .

Output
Print a single integer — the maximum length of the resulting arrays after some operations were performed on arrays 𝐴
and 𝐵
in such a way that they became equal.

If there is no way to make array equal, print "-1".

### observations

1. 如果 sum(A) != sum(B) => -1
2. else 至少是1
3. 然后如果存在，那么某个前缀和必然相等的，剩余的部分也是相等的
4. 所以可以贪心处理