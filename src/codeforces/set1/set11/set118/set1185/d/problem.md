A sequence 𝑎1,𝑎2,…,𝑎𝑘
 is called an arithmetic progression if for each 𝑖
 from 1
 to 𝑘
 elements satisfy the condition 𝑎𝑖=𝑎1+𝑐⋅(𝑖−1)
 for some fixed 𝑐
.

For example, these five sequences are arithmetic progressions: [5,7,9,11]
, [101]
, [101,100,99]
, [13,97]
 and [5,5,5,5,5]
. And these four sequences aren't arithmetic progressions: [3,1,2]
, [1,2,4,8]
, [1,−1,1,−1]
 and [1,2,3,3,3]
.

You are given a sequence of integers 𝑏1,𝑏2,…,𝑏𝑛
. Find any index 𝑗
 (1≤𝑗≤𝑛
), such that if you delete 𝑏𝑗
 from the sequence, you can reorder the remaining 𝑛−1
 elements, so that you will get an arithmetic progression. If there is no such index, output the number -1.
 