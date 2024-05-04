You have an array a with length n, you can perform operations. Each operation is like this: choose two adjacent elements
from a, say x and y, and replace one of them with gcd(x, y), where gcd denotes the greatest common divisor.

What is the minimum number of operations you need to make all of the elements equal to 1?

### ideas

1. 显然，如果gcd(a...) > 1 => -1
2. 否则，肯定是可以的
3. 其次应该找到最快的能够变成1的区间，然后用这个区间去处理其他的？
4. n <= 2000
5. dp[i][j]表示将i..j处理成它们最小g的最小次数
6. 似乎进行不下去
7. 如果已经有一个1了，那么就可以用这个1去处理其他的数
8. 如果不存在这样一个数，那要怎样获得这样的一个数呢？
9. 