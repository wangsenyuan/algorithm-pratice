QED is given a permutation∗
 𝑝
 of length 𝑛
. He also has a string 𝑠
 of length 𝑛
 containing only characters 𝙻
 and 𝚁
. QED only likes permutations that are sorted in non-decreasing order. To sort 𝑝
, he can select any of the following operations and perform them any number of times:

Choose an index 𝑖
 such that 𝑠𝑖=𝙻
. Then, swap 𝑝𝑖
 and 𝑝𝑖−1
. It is guaranteed that 𝑠1≠𝙻
.
Choose an index 𝑖
 such that 𝑠𝑖=𝚁
. Then, swap 𝑝𝑖
 and 𝑝𝑖+1
. It is guaranteed that 𝑠𝑛≠𝚁
.
He is also given 𝑞
 queries. In each query, he selects an index 𝑖
 and changes 𝑠𝑖
 from 𝙻
 to 𝚁
 (or from 𝚁
 to 𝙻
). Note that the changes are persistent.

After each query, he asks you if it is possible to sort 𝑝
 in non-decreasing order by performing the aforementioned operations any number of times. Note that before answering each query, the permutation 𝑝
 is reset to its original form.