Chaneka has an array [𝑎1,𝑎2,…,𝑎𝑛]
. Initially, all elements are white. Chaneka will choose one or more different indices and colour the elements at those
chosen indices black. Then, she will choose all white elements whose indices are multiples of the index of at least one
black element and colour those elements green. After that, her score is the maximum value of 𝑎𝑖
out of all black and green elements.

There are 2𝑛−1
ways for Chaneka to choose the black indices. Find the sum of scores for all possible ways Chaneka can choose the black
indices. Since the answer can be very big, print the answer modulo 998244353
.

### thoughts

1. 考虑贡献性吗？
2. 如果下标i被涂抹成了黑色，假设j也被涂成了黑色
3. 假设 i < j,
4. 先考虑一个例子, (3, 4) 那么被涂成绿色的下标为 6,8,9,12,15,16,.... 等等
5. 假设要让 a[k]作为最大值，有多少种选择可以使得它们成为最大值呢？
6. 如果 k的倍数存在比它大的，显然它的贡献是0，
7. 假设不存在这样的数，一种是k被涂成黑色，还有一种是，它的某个因子被涂成黑色，但是如何保证它肯定是最大值呢？
8. 假设按照a递减不断放入数组，所以那些在k前面出现的下标不能被涂成黑色
9. 然后，看那些没有被标记的位置，且为k的因子的位置，剩余多少个，