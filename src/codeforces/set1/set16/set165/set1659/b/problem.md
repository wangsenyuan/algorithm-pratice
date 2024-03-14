You are given a binary string of length 𝑛
. You have exactly 𝑘
moves. In one move, you must select a single bit. The state of all bits except that bit will get flipped (0
becomes 1
, 1
becomes 0
). You need to output the lexicographically largest string that you can get after using all 𝑘
moves. Also, output the number of times you will select each bit. If there are multiple ways to do this, you may output
any of them.

A binary string 𝑎
is lexicographically larger than a binary string 𝑏
of the same length, if and only if the following holds:

in the first position where 𝑎
and 𝑏
differ, the string 𝑎
contains a 1
, and the string 𝑏
contains a 0
.

### ideas

1. 从高到低，希望获取到足够多的1
2. 搞错了，是翻转除i外的其他位置
3. 如果考虑一堆pair(i, j) s[i] = 0, s[j] = 1
4. 那么在(i, j)上经过操作后，相当于交换了它们的位置
5. 如果剩余1次，就是把剩余的除了最后一个0，其他都变成1即可