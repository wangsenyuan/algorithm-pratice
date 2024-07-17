A permutation is a sequence of 𝑛
 integers from 1
 to 𝑛
, in which all the numbers occur exactly once. For example, [1]
, [3,5,2,1,4]
, [1,3,2]
 are permutations, and [2,3,2]
, [4,3,1]
, [0]
 are not.

Polycarp was given four integers 𝑛
, 𝑙
, 𝑟
 (1≤𝑙≤𝑟≤𝑛)
 and 𝑠
 (1≤𝑠≤𝑛(𝑛+1)2
) and asked to find a permutation 𝑝
 of numbers from 1
 to 𝑛
 that satisfies the following condition:

𝑠=𝑝𝑙+𝑝𝑙+1+…+𝑝𝑟
.
For example, for 𝑛=5
, 𝑙=3
, 𝑟=5
, and 𝑠=8
, the following permutations are suitable (not all options are listed):

𝑝=[3,4,5,2,1]
;
𝑝=[5,2,4,3,1]
;
𝑝=[5,2,1,3,4]
.
But, for example, there is no permutation suitable for the condition above for 𝑛=4
, 𝑙=1
, 𝑟=1
, and 𝑠=5
.
Help Polycarp, for the given 𝑛
, 𝑙
, 𝑟
, and 𝑠
, find a permutation of numbers from 1
 to 𝑛
 that fits the condition above. If there are several suitable permutations, print any of them.

### ideas
1. 最小的sum是 1 + 2 + ... + r - l + 1
2. 最大的是 n + n - 1 + ... + n - (r - l)
3. 如果不在这个范围内， no answer，else应该是肯定存在答案的
4. 假设 和 最大值的 diff，那么就用 1 替换 n  -n + 1