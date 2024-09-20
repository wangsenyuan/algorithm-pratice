Given two integers n
 and x
, find the number of triplets (a,b,c
) of positive integers such that ab+ac+bc≤n
 and a+b+c≤x
.

Note that order matters (e.g. (1,1,2
) and (1,2,1
) are treated as different) and a
, b
, c
 must be strictly greater than 0
.

### ideas
1. a * b + b * c + a * c <= n
2. a + b + c <= x
3. 不妨设 a <= b <= c
4. 3 * a * a <= n
5. a的上限可以确定是 n/3 的平方根
6. 那么b的范围就是 a...n/a (这里a越大，这个范围会越小)
7. c的范围是b...x - a - b
8. 