You are given an array 𝑎1,𝑎2,…,𝑎𝑛
 of 𝑛
 integers. Consider 𝑆
 as a set of segments satisfying the following conditions.

Each element of 𝑆
 should be in form [𝑥,𝑦]
, where 𝑥
 and 𝑦
 are integers between 1
 and 𝑛
, inclusive, and 𝑥≤𝑦
.
No two segments in 𝑆
 intersect with each other. Two segments [𝑎,𝑏]
 and [𝑐,𝑑]
 intersect if and only if there exists an integer 𝑥
 such that 𝑎≤𝑥≤𝑏
 and 𝑐≤𝑥≤𝑑
.
For each [𝑥,𝑦]
 in 𝑆
, 𝑎𝑥+𝑎𝑥+1+…+𝑎𝑦≥0
.
The length of the segment [𝑥,𝑦]
 is defined as 𝑦−𝑥+1
. 𝑓(𝑆)
 is defined as the sum of the lengths of every element in 𝑆
. In a formal way, 𝑓(𝑆)=∑[𝑥,𝑦]∈𝑆(𝑦−𝑥+1)
. Note that if 𝑆
 is empty, 𝑓(𝑆)
 is 0
.

What is the maximum 𝑓(𝑆)
 among all possible 𝑆
?

### ideas
1. 假设S中的两个元素x, y, 它们靠在一起，那么把它们合并在一起，结果不会更差（且肯定可以合并在一起）
2. 因为 sum(x) >= 0, sum(y) >= 0, 所以 sum(x + y) >= 0也成立
3. 那么对于所有的r，可以找到一个最远的l, sum(l...r) >= 0
4. dp[r] = dp[r-1] (如果r不参与)
5.       max dp[l-1] + (r - l + 1)
6. 剩下的问题怎么找到r的l
7. pref[r] - pref[l] >= 0
8. pref[r] >= pref[l]
9. 那么就是左边，最左边的下标l, pref[r] >= pref[l]