You are given an undirected graph where each edge has one of two colors: black or red.

Your task is to assign a real number to each node so that:

for each black edge the sum of values at its endpoints is 1
;
for each red edge the sum of values at its endpoints is 2
;
the sum of the absolute values of all assigned numbers is the smallest possible.
Otherwise, if it is not possible, report that there is no feasible assignment of the numbers.

### ideas
1. 居然是real number。
2. 对于black来说，考虑一端连续的节点，组成的一条黑色的路径，如果这个路径的长度超过4，那么就有 a[1] = a[3], a[2] = a[4]
3. 且 a[1] + a[2] = 1.0
4. 所以，这里出现这样一个pattern，就是黑色的边组成的一个子图，红色的边组成的子图
5. 那些存在红色、黑色边的节点的value，不能超过1.0 （但似乎是可以低于1.0的）
6. 还可以给负值。。。
7. 先不要考虑值的问题
8. 只考虑一个黑色的子图，考虑一个长度为3的圈，似乎只有0.5这样的赋值才可以
9. 但是如果是长度为4的圈，似乎就没有限制，a[0] = a[2], a[1] = a[3], a[0] + a[1] = 1.0即可
10. 也就是说如果这个子图中，存在奇数的圈，那么只能全部赋值0.5；否则可以认为没有限制
11. 同样在红色的子图中，如果存在奇数的圈，那么只能全部赋值1.0， 否则也没有限制
12. 然后在没有限制的情况下，如何最小化sum呢？和能够赋值的个数有关系，在一个子图中，如果进行图色后，蓝色如果多于绿色，那么蓝色越小越好，就是0；
13. 现在考虑那些交界的点，如果冲突了，比如既要赋值0.5，又要赋值1.0， 那么显然是没有答案的
14. 否则的话，在没有限制的情况下，可以用上面的作色方案来统一处理吗？
15. 假设在一个完整的图中，有红色的，也有黑色的，且都没有值的限制（即只有偶数的圈）
16. 其中blue和green的数量已知，且有a + b = blue, c + d = green
17. 那么赋值后期望 x * a + y * b + (2 - x) * c + (1 - y) * d 最小
18. x * (a - c) + y * (b - d) 最小
19. b - d = blue - green - (a - c) = D - (a - c)
20. x * (a - c) + y * (D - (a - c)) = (x - y) * (a - c) + y * D
21. 如果这里y = 0, 或者 x = 0
22. 应该只需要检查几种可能的赋值情况，2，1，0