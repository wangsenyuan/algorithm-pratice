You have array a that contains all integers from 1 to n twice. You can arbitrary permute any numbers in a.

Let number i be in positions xi, yi (xi < yi) in the permuted array a. Let's define the value di = yi - xi — the distance between the positions of the number i. Permute the numbers in array a to minimize the value of the sum .

### ideas
1. sum((n - i) * abs(di +i - n))
2. 如果 di + i >= n
3. （n - i) * (di + i - n)
4. di 等于数字i的两个位置的距离
5. 考虑1，2，3.。。。n
6. 考虑n， n - n = 0, 所以n的位置没有关系
7. 考虑i，(n - i) * abs(di + i - n)
8. 如果某个数的位置正好时 n - i, 那么 di + i = n, 那么后半部分 = 0
9. 是否能够始终让 di = n - i 呢？
10. 考虑1， 2， 3
11. 1, 3, 1, 2, 2, 3
12. 考虑1， 3， 4， 4
13. 1, 3, 3, 1, 2, 4, 2, 4
14. 似乎是可以的
15. pow(i) = i, i + n - i (=n, 但是位置n只能放一个数)
16. 如果在n+1放置一个=n + 1 - i 位置放置i
17. 1(?) (2)。。。(n) (n+1) 