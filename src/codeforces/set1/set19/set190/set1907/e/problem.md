Given a non-negative integer number 𝑛
(𝑛≥0
). Let's say a triple of non-negative integers (𝑎,𝑏,𝑐)
is good if 𝑎+𝑏+𝑐=𝑛
, and 𝑑𝑖𝑔𝑠𝑢𝑚(𝑎)+𝑑𝑖𝑔𝑠𝑢𝑚(𝑏)+𝑑𝑖𝑔𝑠𝑢𝑚(𝑐)=𝑑𝑖𝑔𝑠𝑢𝑚(𝑛)
, where 𝑑𝑖𝑔𝑠𝑢𝑚(𝑥)
is the sum of digits of number 𝑥
.

For example, if 𝑛=26
, then the pair (4,12,10)
is good, because 4+12+10=26
, and (4)+(1+2)+(1+0)=(2+6)
.

Your task is to find the number of good triples for the given number 𝑛
. The order of the numbers in a triple matters. For example, the triples (4,12,10)
and (10,12,4)
are two different triples.

### thoughts

1. 对于n来说， digsum 是确定的假设是 dsn
2. dsn <= 9 * 7 = 63
3. 可以迭代ds吗？ da + db + dc = dsn
4. 对于给定的ds, 可以计算有多少个数的digits sum = ds (可以预处理)
5. 然后找出那些 a <= n, b <= n - a, c = n - a - b
6. 可以先计算 n = a + b吗？假设这样的pair的计数已经知道了
7. 每一位不能有carry