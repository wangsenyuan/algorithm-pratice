Let there be a set that contains distinct positive integers. To expand the set to contain as many integers as possible,
Eri can choose two integers 𝑥≠𝑦
from the set such that their average 𝑥+𝑦2
is still a positive integer and isn't contained in the set, and add it to the set. The integers 𝑥
and 𝑦
remain in the set.

Let's call the set of integers consecutive if, after the elements are sorted, the difference between any pair of
adjacent elements is 1
. For example, sets {2}
, {2,5,4,3}
, {5,6,8,7}
are consecutive, while {2,4,5,6}
, {9,7}
are not.

Eri likes consecutive sets. Suppose there is an array 𝑏
, then Eri puts all elements in 𝑏
into the set. If after a finite number of operations described above, the set can become consecutive, the array 𝑏
will be called brilliant.

Note that if the same integer appears in the array multiple times, we only put it into the set once, as a set always
contains distinct positive integers.

Eri has an array 𝑎
of 𝑛
positive integers. Please help him to count the number of pairs of integers (𝑙,𝑟)
such that 1≤𝑙≤𝑟≤𝑛
and the subarray 𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟
is brilliant.

### ideas

1. (4 + 7) / 2 = 5
2. (5 + 7) / 2 = 6
3. (4 + 8) / 2 = 6, ...good
4. (3 + 7) / 2 = 5 (3, 5) = 4, (5, 7) = 6
5. (3 + 6) / 2 = 4,
6. 似乎 l, r 如果 l < r, 那么就可以得到这个区间内所有的数？
7. (l + r) / 2 = mid, 如果 l < mid, 可以继续进行
8. 如果 mid = l, 那么 r = mid + 1
9. 这里不对。 它要求是 （x + y) / 2 是positive integer
10. 也就是说它们的parity必须是一致的
11. 1 3 6 10 15 21
12. 长度=1的，6
13. 长度=2， [1,3], [1, 3, 6] 也可以的, 可以先得到2，然后得到4
14. (6 + 1) * 6 / 2 = 21
15. [10, 15], [3, 6], [15, 21] 也不行
16. 如果只是两位，很好判断，它们parity不一致 => bad，
17. 或者parity一致，但是中位数parity不一致 =》 bad
18. 如果有3个不同的数，a,b,c是不是一定是good的？
19. 如果a < b < c and a, b, c都是奇数
20. (a, b) 操作到无法继续时，必然是出现了第一个偶数，同样的，b,c操作到无法继续时，必然也出现了一个偶数，且d < e
21. 有两个偶数就可以操作下去了
22. 如果两个奇数一个偶数, 如果第一个出现的偶数不是给定的这个偶数，那么就可以一直处理下去，否则也是不行的
23. 比如15, 21, 18
24. 三个偶数也是ok的，2欧一个奇数也和2奇一偶数类似
25. 所以，总结一下，一个子序列，如果只有一个数 => good
26. 如果有两个数, 判断奇偶性
27. 