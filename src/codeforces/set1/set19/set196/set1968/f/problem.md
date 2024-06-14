Let us call an array 𝑥1,…,𝑥𝑚
 interesting if it is possible to divide the array into 𝑘>1
 parts so that bitwise XOR of values from each part are equal.

More formally, you must split array 𝑥
 into 𝑘
 consecutive segments, each element of 𝑥
 must belong to exactly 1
 segment. Let 𝑦1,…,𝑦𝑘
 be the XOR of elements from each part respectively. Then 𝑦1=𝑦2=⋯=𝑦𝑘
 must be fulfilled.

For example, if 𝑥=[1,1,2,3,0]
, you can split it as follows: [1],[1],[2,3,0]
. Indeed 1=1=2⊕3⊕0
.

You are given an array 𝑎1,…,𝑎𝑛
. Your task is to answer 𝑞
 queries:

For fixed 𝑙
, 𝑟
, determine whether the subarray 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
 is interesting.