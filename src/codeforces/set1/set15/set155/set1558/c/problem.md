You have a permutation: an array 𝑎=[𝑎1,𝑎2,…,𝑎𝑛]
 of distinct integers from 1
 to 𝑛
. The length of the permutation 𝑛
 is odd.

You need to sort the permutation in increasing order.

In one step, you can choose any prefix of the permutation with an odd length and reverse it. Formally, if 𝑎=[𝑎1,𝑎2,…,𝑎𝑛]
, you can choose any odd integer 𝑝
 between 1
 and 𝑛
, inclusive, and set 𝑎
 to [𝑎𝑝,𝑎𝑝−1,…,𝑎1,𝑎𝑝+1,𝑎𝑝+2,…,𝑎𝑛]
.

Find a way to sort 𝑎
 using no more than 5𝑛2
 reversals of the above kind, or determine that such a way doesn't exist. The number of reversals doesn't have to be minimized.

 ### ideas
 1. 考虑n的位置，如果n不在最后一位，那么就必须把n先交换到1，然后交换到n（那么每个操作最多2次）
 2. 那最多 2 * n < 5 * n / 2呐？
 3. 这样不是很简单了？不值2000分吧～～～
 4. 问题出在奇数长度的选择
 5. 比如 [2, 1, 3] 这个就没法sort回来
 6. 假设一个原本应该在奇数位置的数在偶数位置了，那么答案 => -1
 7. 因为一个数所在的位置的奇偶性是不会变的，它一开始在偶数位置，选择一个奇数长的前缀，翻转后，它仍然在偶数位置
 8. 否则答案始终是存在的。
 9. 1，4，3，2，5
 10. 因为4不在位置，但是无法交换， 先选择5 [5, 2, 3, 4, 1]
 11. 然后选择3, [3, 2, 5, 4, 1]
 12. 再选择5，[1, 4, 5, 2, 3]
 13. 再选择3，[5, 4, 1, 2, 3]
 14. 再选择5, [3, 2, 1, 4, 5]
 15. 所以，要两个位置一起考虑
 16. 先把n换到位置1，然后换到n-1的前面（肯定是可以的，因为n-1肯定在偶数位置）
 17. 然后选择长度l, 这样就翻转了n-1, n的相对位置
 18. 然后在选择n所在的位置，再选择长度l