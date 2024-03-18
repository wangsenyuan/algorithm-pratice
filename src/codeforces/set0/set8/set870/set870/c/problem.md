You are given several queries. In the i-th query you are given a single positive integer ni. You are to represent ni as
a sum of maximum possible number of composite summands and print this maximum number, or print -1, if there are no such
splittings.

An integer greater than 1 is composite, if it is not prime, i.e. if it has positive divisors not equal to 1 and the
integer itself.

### idea

1. 如果num % 4 == 0， 那么就是 num / 4
2. let r = num % 4, 如果 r = 1, 4 * k + 9
3. r = 2, 4 * k + 6
4. r = 3, 4 * k + 15 = 4 * k + 6 + 9