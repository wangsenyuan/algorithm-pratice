https://codeforces.com/problemset/problem/1896/E

### problem

You are given a permutation†
𝑎
of size 𝑛
. We call an index 𝑖
good if 𝑎𝑖=𝑖
is satisfied. After each second, we rotate all indices that are not good to the right by one position. Formally,

Let 𝑠1,𝑠2,…,𝑠𝑘
be the indices of 𝑎
that are not good in increasing order. That is, 𝑠𝑗<𝑠𝑗+1
and if index 𝑖
is not good, then there exists 𝑗
such that 𝑠𝑗=𝑖
.
For each 𝑖
from 1
to 𝑘
, we assign 𝑎𝑠(𝑖%𝑘+1):=𝑎𝑠𝑖
all at once.
For each 𝑖
from 1
to 𝑛
, find the first time that index 𝑖
becomes good.

†
A permutation is an array consisting of 𝑛
distinct integers from 1
to 𝑛
in arbitrary order. For example, [2,3,1,5,4]
is a permutation, but [1,2,2]
is not a permutation (2
appears twice in the array) and [1,3,4]
is also not a permutation (𝑛=3
but there is 4
in the array).

### thoughts

1. 一旦一个数到达了它的位置，它就退出了
2. 比如对于序列 ... 5, 3, 6...
3. 如果3要到达的位置在当前位置的后面，那么对于5来说，3就没有影响
4. 这是因为，3肯定会在5到达之前到达
5. 同时，这里6对5也没有影响
6. 所以总结出来就是，5要达到的位置pos 5之间，排除掉比5大的数，以及比5小，但是也会在5前面到达的那些
7. 那这里考虑3如果要跑到前面去，要怎么计算它的移动距离呢？可以考虑它也是往后的，那些比它大的忽略

### solution

For each index 𝑖
from 1
to 𝑛
, let ℎ𝑖
denote the number of cyclic shifts needed to move 𝑎𝑖
to its correct spot. In other words, ℎ𝑖
is the minimum value such that (𝑖+ℎ𝑖−1) % 𝑛+1=𝑎𝑖
. How can we get the answer from ℎ𝑖
?

For convenience, we will assume that the array is cyclic, so 𝑎𝑗=𝑎(𝑗−1) % 𝑛+1
. The answer for each index 𝑖
from 1
to 𝑛
is ℎ𝑖
(defined in hint 1) minus the number of indices 𝑗
where 𝑖<𝑗<𝑖+ℎ𝑖
and 𝑖<𝑎𝑗<𝑖+ℎ𝑖
(or 𝑖<𝑎𝑗+𝑛<𝑖+ℎ𝑖
to handle cyclic case when 𝑖+ℎ𝑖>𝑛
). This is because the value that we are calculating is equal to the number of positions that 𝑎𝑖
will skip during the rotation as the index is already good.

To calculate the above value, it is convenient to define an array 𝑏
of size 2𝑛
where 𝑏𝑖=𝑎𝑖
for all 𝑖
between 1
to 𝑛
, and 𝑏𝑖=𝑎𝑖−𝑛+𝑛
for all 𝑖
between 𝑛+1
to 2𝑛
to handle cyclicity. We will loop from 𝑖=2𝑛
to 𝑖=1
, and do a point increment to position 𝑎𝑖
if 𝑎𝑖≥𝑖
, otherwise, do a point increment to position 𝑎𝑖+𝑛
. Then, to get the answer for index 𝑖
, we do a range sum query from 𝑖+1
to 𝑖+ℎ𝑖−1
. Point increment and range sum query can be done using a binary indexed tree in 𝑂(log𝑛)
time per query/update. Hence, the problem can be solved in 𝑂(𝑛log𝑛)
time.