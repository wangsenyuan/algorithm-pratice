You are given two arrays of integers — 𝑎1,…,𝑎𝑛
of length 𝑛
, and 𝑏1,…,𝑏𝑚
of length 𝑚
. You can choose any element 𝑏𝑗
from array 𝑏
(1≤𝑗≤𝑚
), and for all 1≤𝑖≤𝑛
perform 𝑎𝑖=𝑎𝑖|𝑏𝑗
. You can perform any number of such operations.

After all the operations, the value of 𝑥=𝑎1⊕𝑎2⊕…⊕𝑎𝑛
will be calculated. Find the minimum and maximum values of 𝑥
that could be obtained after performing any set of operations.

### thoughts

1. (a | b) ^ (c | b)
2. 在什么情况下，可以取最大值，什么情况下取最小值
3. 如果 a[i] = 1, c[i] = 1, 无论b怎么取值，对结果是无影响的 res[i] = 0
4. 如果 a[i] = 0, c[i] = 1, 如果 b[i] = 1, 那么res[i] = 0, 如果b[i] = 0, res[i] = 1
5. 如果 a[i] = 0, c[i] = 0, 无论b怎么取值, 结果都是 res[i] = 0
6. 为了取最大值，就是要从高往低位取，在保证高位不变的情况下，当前位能成为1
7. 要考虑a中当前位1的个数

### solution

Note that after performing the operation on 𝑏𝑗
, which has some bit set to 1, this bit will become 1 for all numbers in 𝑎
(and will remain so, as a bit cannot change from 1 to 0 in the result of an OR operation). If 𝑛
is even, then in the final XOR, this bit will become 0, as it will be equal to the XOR of an even number of ones. If 𝑛
is odd, then this bit will be 1 in the final XOR.

Therefore, if 𝑛
is even, by performing the operation on 𝑏𝑗
, we set all the bits that are 1 in 𝑏𝑗
to 0 in the final XOR. If 𝑛
is odd, we do the opposite and set these bits to 1. So, if 𝑛
is even, the XOR does not increase when applying the operation, which means that to obtain the minimum possible XOR, we
need to apply the operation to all the numbers in 𝑏
, and the maximum XOR will be the original XOR. For odd 𝑛
, it is the opposite: the minimum is the original XOR, and the maximum is obtained after applying the operation to all
elements in 𝑏
.

To apply the operation to all elements in 𝑏
, it is sufficient to calculate their bitwise OR and apply the operation to the array 𝑎
with it.