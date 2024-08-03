Reca company makes monitors, the most popular of their models is AB999 with the screen size a × b centimeters. Because of some production peculiarities a screen parameters are integer numbers. Recently the screen sides ratio x: y became popular with users. That's why the company wants to reduce monitor AB999 size so that its screen sides ratio becomes x: y, at the same time they want its total area to be maximal of all possible variants. Your task is to find the screen parameters of the reduced size model, or find out that such a reduction can't be performed.

### ideas
1. 假设h，w
2. h <= a,w <= b, h : w = x : y
3. and h * w 最大
4. h * y - w * x = 0
5. a * x + b * y = 0的一组整数解，可用extgcd(a, b)计算出来
6. 假设得到了一组 (u, v) u * x + v * y = 0 (u > 0, v < 0)
7. 然后增加u (以y的倍数增加), v也以x的倍数减少
8. 可以让x和y先除以gcd(x, y), 然后这两个必须在