Diana loves playing with numbers. She's got n
 cards with positive integer numbers ai
 written on them. She spends her free time multiplying the numbers on the cards. She picks a non-empty subset of the cards and multiplies all the numbers ai
 written on them.

Diana is happy when the product of the numbers ends with her favorite digit d
. Now she is curious what cards she should pick so that the product of the numbers on them is the largest possible and the last decimal digit of the product is d
. Please, help her.

Input
The first line contains the integers n
 and d
 (1≤n≤105
, 0≤d≤9
). The second line contains n
 integers ai
 (1≤ai≤1000
).

Output
On the first line, print the number of chosen cards k
 (1≤k≤n
). On the next line, print the numbers written on the chosen cards in any order.

If it is impossible to choose a subset of cards with the product that ends with the digit d
, print the single line with −1
.

### ideas
1. a[i] <= 1000 可能是个突破口
2. 所有的数乘在一起， 也就是 1e5 * 1e3 = 1e8
3. 所以，可以用 dp[i][x] = 余数为x时的最大值 