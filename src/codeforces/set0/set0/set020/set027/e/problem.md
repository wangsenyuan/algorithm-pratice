Given the number n, find the smallest positive integer which has exactly n divisors. It is guaranteed that for the given n the answer will not exceed 1018.

### ideas
1. 如果知道结果x，那么x的除数的数量 = 将x质因数分解后
2. 每个质数的指数，p1, p2, p3,
3. 数量 = (p1 + 1) * (p2 + 1) * (p3 + 1)...
4. 假设就由2，和3组成, 
5. n = (p1 + 1) * (p2 + 1)
6. x = pow(2, p1) * pow(3, p2)
7. 显然p1 = p2的时候， p1和p2的和越小
8. 所以，是把n分解的越多越好
9. 也就是把n值因数分解
10. 假设2,3,5组成
11. (c2 + 1) * (c3 + 1) * (c5 + 1)
12. pow(2, c2) * pow(5, c5)
13. pow(2, c2) * pow(5, c5) > pow(2, c2 + c5)
14. 所以，把5都当做2对n是没有影响的，但是对res是更优的