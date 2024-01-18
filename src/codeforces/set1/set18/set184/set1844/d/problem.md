The row-major order of an 𝑟×𝑐
grid of characters 𝐴
is the string obtained by concatenating all the rows, i.e.
𝐴11𝐴12…𝐴1𝑐𝐴21𝐴22…𝐴2𝑐…𝐴𝑟1𝐴𝑟2…𝐴𝑟𝑐.
A grid of characters 𝐴
is bad if there are some two adjacent cells (cells sharing an edge) with the same character.

You are given a positive integer 𝑛
. Consider all strings 𝑠
consisting of only lowercase Latin letters such that they are not the row-major order of any bad grid. Find any string
with the minimum number of distinct characters among all such strings of length 𝑛
.

It can be proven that at least one such string exists under the constraints of the problem.

### thoughts

1. 考虑n的所有因子1, ...a, b....n
2. factor = 1时，也就是只有一行时，2个就可以了
3. factor = 2时，
4. 那么除了左上角和右上角的两个（两条边），其他点，都有3条边，
5. factor = 3时，最多时4条边
6. 是不是地图着色问题？

### solution

The condition is equivalent to a graph of pairs of characters in 𝑠
that need to be different. In graph-theoretic language, we need to find the chromatic number of this graph.

By considering the 1×𝑛
and 𝑛×1
grids, there is an edge between character 𝑢
and 𝑢+1
for all 1≤𝑢≤𝑛−1
. By considering a 𝑛𝑑×𝑑
grid (where 𝑑
divides 𝑛
), there is an edge between character 𝑢
and 𝑢+𝑑
for all 1≤𝑢≤𝑛−𝑑
whenever 𝑑
divides 𝑛
.

Let 𝑐
be the smallest positive integer that does not divide 𝑛
. There is an edge between every pair of the characters 1,…,𝑐
(in graph-theoretic language, this is a clique), so the answer is at least 𝑐
. On the other hand, consider the string obtained by letting 𝑠1,…,𝑠𝑐
be all distinct characters and repeating this pattern periodically (𝑠𝑖=𝑠𝑖+𝑐
for all 1≤𝑖≤𝑛−𝑐
). Any pair of equal characters have an index difference that is a multiple of 𝑐
, say 𝑘𝑐
. But since 𝑐
does not divide 𝑛
, 𝑘𝑐
also does not divide 𝑛
, so these characters are not connected by an edge. Therefore this construction gives a suitable string with 𝑐
distinct characters.

The time complexity is (𝑛)
.