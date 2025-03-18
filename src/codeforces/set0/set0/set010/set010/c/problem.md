Not long ago Billy came across such a problem, where there were given three natural numbers A, B and C from the range [1, N], and it was asked to check whether the equation AB = C is correct. Recently Billy studied the concept of a digital root of a number. We should remind you that a digital root d(x) of the number x is the sum s(x) of all the digits of this number, if s(x) ≤ 9, otherwise it is d(s(x)). For example, a digital root of the number 6543 is calculated as follows: d(6543) = d(6 + 5 + 4 + 3) = d(18) = 9. Billy has counted that the digital root of a product of numbers is equal to the digital root of the product of the factors' digital roots, i.e. d(xy) = d(d(x)d(y)). And the following solution to the problem came to his mind: to calculate the digital roots and check if this condition is met. However, Billy has doubts that this condition is sufficient. That's why he asks you to find out the amount of test examples for the given problem such that the algorithm proposed by Billy makes mistakes.

### ideas
1. C = AB
2. d(C) = d(AB) = d(d(A) * d(B)) 这个是成立的
3. 但是反过来不成立，就是说如果d(d(A) * d(B)) = d(C)
4. A * B = C 不一定成立
5. 不知道咋入手啊
6. N比较小 1e6
7. 所有的数都可以，根据d(num) 分类到 1....9这些组中
8. 那么 d(C) = c, 就可以有几种分类, 比如 a * b (这个可以全部计算出来)
9. C = A * B (这个数量 = 它的质因数分解后，可以计算出来)
10. 可以搞了