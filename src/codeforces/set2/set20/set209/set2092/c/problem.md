For her birthday, each of Asuna's 𝑛
 admirers gifted her a tower. The height of the tower from the 𝑖
-th admirer is equal to 𝑎𝑖
.

Asuna evaluates the beauty of the received gifts as max(𝑎1,𝑎2,…,𝑎𝑛)
. She can perform the following operation an arbitrary number of times (possibly, zero).

Take such 1≤𝑖≠𝑗≤𝑛
 that 𝑎𝑖+𝑎𝑗
 is odd and 𝑎𝑖>0
, then decrease 𝑎𝑖
 by 1
 and increase 𝑎𝑗
 by 1
.
It is easy to see that the heights of the towers remain non-negative during the operations.
Help Asuna find the maximum possible beauty of the gifts after any number of operations!

### ideas
1. a[i] + a[j]是奇数的时候才能增加
2. 如果只有两个数，且a[i] + a[j]是奇数，那么它们是一直可以变化，知道a[j] = s, a[i] = 0
3. a[i] + a[j] = 奇数，然后再和一个偶数，还是奇数 + 偶数。。。
4. 也就是说，一个奇数，可以把所有的偶数都累加上去
5. 【5 4 3 2 9】
6. (6, 3, 3, 9)
7. (7, 2, 3, 9)
8. (7, 1, 3, 10)
9. (8, 0, 3, 10)
10. (8, 0, 0, 13)
11. (21, 0, 0, 0)
12. 如果没有奇数/偶数，那么只能是最大值
13. 如果存在偶数/奇数， 奇数可以变成偶数，偶数变奇数
14. 6 + 6 + 9
15. 9, 2, 11
16. 10, 1, 11 => 21