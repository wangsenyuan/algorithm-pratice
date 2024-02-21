Define a set 𝐴
as a child of set 𝐵
if and only if for each element of value 𝑥
that is in 𝐴
, there exists an element of value 𝑥+1
that is in 𝐵
.

Given integers 𝑁
and 𝐾
. In order to make Chaneka happy, you must find the number of different arrays containing 𝑁
sets [𝑆1,𝑆2,𝑆3,…,𝑆𝑁]
such that: - Each set can only contain zero or more different integers from 1
to 𝐾
. - There exists a way to rearrange the order of the array of sets into [𝑆′1,𝑆′2,𝑆′3,…,𝑆′𝑁]
such that 𝑆′𝑖
is a child of 𝑆′𝑖+1
for each 1≤𝑖≤𝑁−1
.

Print the answer modulo 998244353
.

Two sets are considered different if and only if there is at least one value that only occurs in one of the two sets.

Two arrays of sets are considered different if and only if there is at least one index 𝑖
such that the sets at index 𝑖
in the two arrays are different.

### thoughts

1. 如果N = 1, 2 ** k
2. N > k + 1呢？ 0，
3. 首先这n个集合内的元素，必然是连续的，且长度要么是N, N -1, N - 2, ... 1, 0
4. 且集合的位置，根据它最小的元素进行排序
5. 比如1，2，3，，，，N
6. 2。。。。。。N + 1
7. 如果1开头的序列存在，这个序列内的数字，要放入的集合是确定的，它唯一能确定的就是，要么是n个长度
8. 要们是n-1个长度
9. 也就是说1有3中选择，要么不出现，要们出现在位置0，要们出现在位置1
10. 2，要4种选择。。。。
11. 3 * 4 * 5 。。。(i + 2) * f(n)