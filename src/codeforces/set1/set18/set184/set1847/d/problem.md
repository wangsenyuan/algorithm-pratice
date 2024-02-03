Let 𝑠
be a binary string of length 𝑛
. An operation on 𝑠
is defined as choosing two distinct integers 𝑖
and 𝑗
(1≤𝑖<𝑗≤𝑛
), and swapping the characters 𝑠𝑖,𝑠𝑗
.

Consider the 𝑚
strings 𝑡1,𝑡2,…,𝑡𝑚
, where 𝑡𝑖
is the substring †
of 𝑠
from 𝑙𝑖
to 𝑟𝑖
. Define 𝑡(𝑠)=𝑡1+𝑡2+…+𝑡𝑚
as the concatenation of the strings 𝑡𝑖
in that order.

There are 𝑞
updates to the string. In the 𝑖
-th update 𝑠𝑥𝑖
gets flipped. That is if 𝑠𝑥𝑖=1
, then 𝑠𝑥𝑖
becomes 0
and vice versa. After each update, find the minimum number of operations one must perform on 𝑠
to make 𝑡(𝑠)
lexicographically as large‡
as possible.

Note that no operation is actually performed. We are only interested in the number of operations.

### thoughts

1. 先不考虑flip，要怎么操作使得t(s)最大呢？最大的t(s) = 111...1000...0
2. 似乎也挺难的～
3. 貌似还不能对区间排序，因为它就是按照这个区间连接的
4. 在处理第i个区间时，前面0/1的个数是可以计算出来的
5. 假设这个区间内有x个1，它们都需要移动到前面去 如果 x <= cnt[0] 前面的个数，这时不需要考虑这个区间内的情况
6. 因为这x个1可以和前面的0进行交换
7. 如果 x > cnt[0]呢？ cnt[0]个1可以交换，剩下的 x - cnt[0]个1也需要交换，也可能不需要交换？

### solution

Lets assume you know string 𝑡
. String 𝑡
is made by positions in 𝑠
. Lets denote 𝑓(𝑖)=
position in 𝑠
from which 𝑡𝑖
is made. For maximising 𝑡
you need to make the starting elements in 𝑡
as large as possible. Now, to make 𝑡
lexicographically as large as possible, we need to swap positions in 𝑠
. We can swap positions greedily. We first try making 𝑠𝑓(1)=1
. Then we try making 𝑠𝑓(2)=1
and so on.

Now, suppose for two indices 𝑖
,𝑗
(1≤𝑖<𝑗≤|𝑡|
) such that 𝑓(𝑖)=𝑓(𝑗)
, we know that index 𝑗
is waste. 𝑡
is basically the preference of indices in 𝑠
which should be equal to 1
. If 𝑠𝑓(𝑗)
is to be set 1
, then it would already be set 1
because before setting 𝑠𝑓(𝑗)
equal to 1
we would have already set 𝑠𝑓(𝑖)
equal to 1
because 𝑓(𝑖)
is equal to 𝑓(𝑗)
. Hence, for each index 𝑖
in 𝑠
, we only add its first occurrence in 𝑡
. This makes the size of 𝑡
bound by size of 𝑠
. Now, this 𝑡
can be found using various pbds like set,dsu,segment tree,etc.

Now, before answering the queries, we find the answer for the current string 𝑠
. We know that there are 𝑥
ones and 𝑛−𝑥
zeros in 𝑠
. So, for each 𝑖
(1≤𝑖≤𝑚𝑖𝑛(𝑥,|𝑡|)
), we make 𝑠𝑓(𝑖)=1
. Hence, the number of swaps needed will become number of positions 𝑖
(1≤𝑖≤𝑚𝑖𝑛(𝑥,|𝑡|)
) such that 𝑠𝑖=0
.

Now, in each query, there is exactly one positions that is either getting flipped from 0
to 1
or from 1
to 0
. That is, 𝑥
is either getting changed to 𝑥+1
or 𝑥−1
. You already know the answer for 𝑥
. Now, if 𝑥
is getting reduced, then you need to decrease the answer by one if 𝑥<=|𝑡|
and 𝑠𝑓(𝑥)=0
. If 𝑥
is increasing by one, then you need to add one to the answer if 𝑥<|𝑡|
and 𝑠𝑓(𝑥+1)=0
.

Time complexity - 𝑂(𝑛log𝑛+𝑞)
.