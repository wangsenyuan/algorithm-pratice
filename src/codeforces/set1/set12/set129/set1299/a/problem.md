Anu has created her own function 𝑓
: 𝑓(𝑥,𝑦)=(𝑥|𝑦)−𝑦
where |
denotes the bitwise OR operation. For example, 𝑓(11,6)=(11|6)−6=15−6=9
. It can be proved that for any nonnegative numbers 𝑥
and 𝑦
value of 𝑓(𝑥,𝑦)
is also nonnegative.

She would like to research more about this function and has created multiple problems for herself. But she isn't able to
solve all of them and needs your help. Here is one of these problems.

A value of an array [𝑎1,𝑎2,…,𝑎𝑛]
is defined as 𝑓(𝑓(…𝑓(𝑓(𝑎1,𝑎2),𝑎3),…𝑎𝑛−1),𝑎𝑛)
(see notes). You are given an array with not necessarily distinct elements. How should you reorder its elements so that
the value of the array is maximal possible?

Input
The first line contains a single integer 𝑛
(1≤𝑛≤105
).

The second line contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(0≤𝑎𝑖≤109
). Elements of the array are not guaranteed to be different.

### ideas

1. 如果 x & y = x， x是y的子集, x | y = y， f(x, y) = (x | y) - y = 0
2. 所以如果最后的y包含所有的1的，那么最后的结果就会digits越少
3. 所以，应该尽量使这样的数字成为x
4. b2 = f(a1, a2) = a1 | a2 - a2
5. f(b2, a3) = b2 | a3 - a3
6. 如果一位出现了两次，会怎么样？那么这一位就会变成0，因为无论怎么安排顺序，它都会在第二次出现时被消除掉
7. 如果只出现了一次，那除非它在a1中，否则也会被无法保留
8. 所以，就看最大能保留的数时多少