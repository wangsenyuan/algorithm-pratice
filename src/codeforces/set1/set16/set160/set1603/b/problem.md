YouKn0wWho has two even integers 𝑥
and 𝑦
. Help him to find an integer 𝑛
such that 1≤𝑛≤2⋅1018
and 𝑛mod𝑥=𝑦mod𝑛
. Here, 𝑎mod𝑏
denotes the remainder of 𝑎
after division by 𝑏
. If there are multiple such integers, output any. It can be shown that such an integer always exists under the given
constraints.

### ideas

1. n % x = r => n = x * m + r
2. y % n = r => y = n * k + r
3. r < x and r < n
4. 先分析一下n的奇偶性
5. 如果n是奇数，那么r必须也是奇数， y可以是偶数（只要k是奇数）
6. 如果n是偶数, r是偶数， y还是偶数（对k和m没有要求）
7. 可以先处理n = x/y的特殊情况
8. 如果x = y, 那么n = x
9. n < x, n % x = n, y % n = n 不行 => n > x
10. 如果 y < n => y % n = y, n % x = y => n = x * k + y => n = x + y (n > y)
11. x = 4, y = 2, n = 10
12. 6 % 4 = 2, 2 % 6 = 2
13. n < y
14. x < n < y
15. n %x = y % n
16. y % n = n % x < x
17. 