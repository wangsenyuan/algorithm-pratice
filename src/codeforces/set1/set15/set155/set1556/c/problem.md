William has a favorite bracket sequence. Since his favorite sequence is quite big he provided it to you as a sequence of positive integers 𝑐1,𝑐2,…,𝑐𝑛
 where 𝑐𝑖
 is the number of consecutive brackets "(" if 𝑖
 is an odd number or the number of consecutive brackets ")" if 𝑖
 is an even number.

For example for a bracket sequence "((())()))" a corresponding sequence of numbers is [3,2,1,3]
.

You need to find the total number of continuous subsequences (subsegments) [𝑙,𝑟]
 (𝑙≤𝑟
) of the original bracket sequence, which are regular bracket sequences.

A bracket sequence is called regular if it is possible to obtain correct arithmetic expression by inserting characters "+" and "1" into this sequence. For example, sequences "(())()", "()" and "(()(()))" are regular, while ")(", "(()" and "(()))(" are not.

### ideas
1. pref(sum) = c[i] * (1 if odd, -1 if even)
2. regular =  pref(r) - pref(l-1) = 0, 且在l...r中间，没有出现负值
3. 一旦出现一个负值，说明右括号太多了。应该直接舍弃掉