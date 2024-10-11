You are given a sequence of 𝑛
 pairs of integers: (𝑎1,𝑏1),(𝑎2,𝑏2),…,(𝑎𝑛,𝑏𝑛)
. This sequence is called bad if it is sorted in non-descending order by first elements or if it is sorted in non-descending order by second elements. Otherwise the sequence is good. There are examples of good and bad sequences:

𝑠=[(1,2),(3,2),(3,1)]
 is bad because the sequence of first elements is sorted: [1,3,3]
;
𝑠=[(1,2),(3,2),(1,2)]
 is bad because the sequence of second elements is sorted: [2,2,2]
;
𝑠=[(1,1),(2,2),(3,3)]
 is bad because both sequences (the sequence of first elements and the sequence of second elements) are sorted;
𝑠=[(1,3),(3,3),(2,2)]
 is good because neither the sequence of first elements ([1,3,2])
 nor the sequence of second elements ([3,3,2])
 is sorted.
Calculate the number of permutations of size 𝑛
 such that after applying this permutation to the sequence 𝑠
 it turns into a good sequence.

A permutation 𝑝
 of size 𝑛
 is a sequence 𝑝1,𝑝2,…,𝑝𝑛
 consisting of 𝑛
 distinct integers from 1
 to 𝑛
 (1≤𝑝𝑖≤𝑛
). If you apply permutation 𝑝1,𝑝2,…,𝑝𝑛
 to the sequence 𝑠1,𝑠2,…,𝑠𝑛
 you get the sequence 𝑠𝑝1,𝑠𝑝2,…,𝑠𝑝𝑛
. For example, if 𝑠=[(1,2),(1,3),(2,3)]
 and 𝑝=[2,3,1]
 then 𝑠
 turns into [(1,3),(2,3),(1,2)]
.

### ideas
1. good = tot - bad?
2. tot 比较好处理
3. bad = bad first + bad second - bad first & second