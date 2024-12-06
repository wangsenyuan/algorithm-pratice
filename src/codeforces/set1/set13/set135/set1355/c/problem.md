Like any unknown mathematician, Yuri has favourite numbers: 𝐴
, 𝐵
, 𝐶
, and 𝐷
, where 𝐴≤𝐵≤𝐶≤𝐷
. Yuri also likes triangles and once he thought: how many non-degenerate triangles with integer sides 𝑥
, 𝑦
, and 𝑧
 exist, such that 𝐴≤𝑥≤𝐵≤𝑦≤𝐶≤𝑧≤𝐷
 holds?

Yuri is preparing problems for a new contest now, so he is very busy. That's why he asked you to calculate the number of triangles with described property.

The triangle is called non-degenerate if and only if its vertices are not collinear.

### ideas
1. x + y > z (只要满足这个条件就可以了)
2. (x <= y <= z)
3. 在给定x、y的时候，很容易计算z
4. 如何在给定z的时候，计算x和y？
5. z => x + y > z
6. 迭代x，那么x可以作为一个确定的参数
7. z = x0 + y的一个函数
8. y的取值范围是 (b, c)
9. res += (d - (b + x0 + 1)) * (c - b - 1) / 2