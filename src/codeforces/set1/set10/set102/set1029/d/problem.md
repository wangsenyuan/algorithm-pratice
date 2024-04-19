You are given an array 𝑎
, consisting of 𝑛
positive integers.

Let's call a concatenation of numbers 𝑥
and 𝑦
the number that is obtained by writing down numbers 𝑥
and 𝑦
one right after another without changing the order. For example, a concatenation of numbers 12
and 3456
is a number 123456
.

Count the number of ordered pairs of positions (𝑖,𝑗)
(𝑖≠𝑗
) in array 𝑎
such that the concatenation of 𝑎𝑖
and 𝑎𝑗
is divisible by 𝑘
.

Input
The first line contains two integers 𝑛
and 𝑘
(1≤𝑛≤2⋅105
, 2≤𝑘≤109
).

The second line contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(1≤𝑎𝑖≤109
).

Output
Print a single integer — the number of ordered pairs of positions (𝑖,𝑗)
(𝑖≠𝑗
) in array 𝑎
such that the concatenation of 𝑎𝑖
and 𝑎𝑗
is divisible by 𝑘
.

### ideas

1. interesting
2. 如果 a[j] % k = 0, 那么其他的数在它前面，就仍然是可以被整除
3. 如果y = a[j] % k != 0
4. (a[i] ++ a[j]) = a[i] * pow(10,length of a[j]) + a[j]
5. 如果 a[i] * pow(10, ...) % k = x
6. 那么如果 (x + y) = k 那么也是 ok的
7. 数字长度不会超过10，所以对于所有的数 a[i]， 可以按照长度1....10 计算%k的值