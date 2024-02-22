You are given a 1
-indexed array 𝑎
of length 𝑛
where each element is 1
or 2
.

Process 𝑞
queries of the following two types:

"1 s": check if there exists a subarray†
of 𝑎
whose sum equals to 𝑠
.
"2 i v": change 𝑎𝑖
to 𝑣
.

### thoughts

1. 因为数字是1/2， 所以，sum的范围是n..2*n
2. 先不考虑更新，如何快速的查找sum呢？
3. 如何利用数字是1/2这个特性呢？sum / 2 <= sz <= min(n, sum)