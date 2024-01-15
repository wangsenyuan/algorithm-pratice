A compass points directly toward the morning star. It can only point in one of eight directions: the four cardinal
directions (N, S, E, W) or some combination (NW, NE, SW, SE). Otherwise, it will break.

The directions the compass can point.
There are 𝑛
distinct points with integer coordinates on a plane. How many ways can you put a compass at one point and the morning
star at another so that the compass does not break?

### thoughts

1. 假设在(x0, y0)放置了morning star, 那么指南针，只能位于4条线上
2. 和 y = y0 的水平线，
3. 和 x = x0 的竖直线
4. y = x - x0 + y0 的斜线
5. y = -x + x0 + y0的斜线
6. 可以把这些线按照水平，竖直，斜线等进行分类
7. 但是，剩下就有问题了，因为两条线，相交得到一个点
8. 不可能迭代任意两类，进行判断
9. 那还不如直接迭代这个平面上的点
10. 是不是看起来是n * n 的, 是不是只需要迭代其中比较少的一些组合就可以了？
11. 看错题目了
12. 它问的是在这些位置中，选择一对点，分别放置指南针和star，的有效的个数