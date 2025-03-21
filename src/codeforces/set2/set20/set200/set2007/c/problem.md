Dora has just learned the programming language C++!

However, she has completely misunderstood the meaning of C++. She considers it as two kinds of adding operations on the array 𝑐
 with 𝑛
 elements. Dora has two integers 𝑎
 and 𝑏
. In one operation, she can choose one of the following things to do.

Choose an integer 𝑖
 such that 1≤𝑖≤𝑛
, and increase 𝑐𝑖
 by 𝑎
.
Choose an integer 𝑖
 such that 1≤𝑖≤𝑛
, and increase 𝑐𝑖
 by 𝑏
.
Note that 𝑎
 and 𝑏
 are constants, and they can be the same.

Let's define a range of array 𝑑
 as max(𝑑𝑖)−min(𝑑𝑖)
. For instance, the range of the array [1,2,3,4]
 is 4−1=3
, the range of the array [5,2,8,2,2,1]
 is 8−1=7
, and the range of the array [3,3,3]
 is 3−3=0
.

After any number of operations (possibly, 0
), Dora calculates the range of the new array. You need to help Dora minimize this value, but since Dora loves exploring all by herself, you only need to tell her the minimized value.

### ideas
1. d[i] = c[i] + n * a + m * b
2. d[i] - c[i] = n * a + m * b
3. 怎么感觉和gcd有关系呢？
4. 不管d[i] - c[i], w = a* x + b * y
5. 这个就是斜率为 - a/b 的直线
6. 不同的d[i] - c[i] = w 就是在不同线的位置
7. 那么就是看这些线的最短的距离？
8. 那么每次改变的单位 = gcd(a, b)?
9. 然后用c[i] 对 g求余