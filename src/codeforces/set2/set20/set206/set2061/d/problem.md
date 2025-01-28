Kevin wrote an integer sequence 𝑎
of length 𝑛
on the blackboard.

Kevin can perform the following operation any number of times:

Select two integers 𝑥,𝑦
on the blackboard such that |𝑥−𝑦|≤1
, erase them, and then write down an integer 𝑥+𝑦
instead.
Kevin wants to know if it is possible to transform these integers into an integer sequence 𝑏
of length 𝑚
through some sequence of operations.

Two sequences 𝑎
and 𝑏
are considered the same if and only if their multisets are identical. In other words, for any number 𝑥
, the number of times it appears in 𝑎
must be equal to the number of times it appears in 𝑏
.

### ideas

1. 和bitmask有啥关系？
2. 两个数能够合并，必须是相邻的两个数
3. 如果a[n] > b[m] => no (因为 a[i]只会变大)
4. 