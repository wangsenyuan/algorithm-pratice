You are given an array 𝑎
of length 𝑛
, consisting of positive integers.

You can perform the following operation on this array any number of times (possibly zero):

choose three integers 𝑙
, 𝑟
and 𝑥
such that 1≤𝑙≤𝑟≤𝑛
, and multiply every 𝑎𝑖
such that 𝑙≤𝑖≤𝑟
by 𝑥
.
Note that you can choose any integer as 𝑥
, it doesn't have to be positive.

You have to calculate the minimum number of operations to make the array 𝑎
sorted in strictly ascending order (i. e. the condition 𝑎1<𝑎2<⋯<𝑎𝑛
must be satisfied).

### thoughts

1. a[i] > 0 => 肯定有答案
2. 如果 a[i-1] < a[i] （通过变换，或者不变换）不需要操作
3. 如果 a[i-1] > a[i] 如果a[i-1]之前现在是正数，通过 *-1， 可以变成负数
4. 但是这样有个问题，会改变前面的关系
5. 比如 1, 3, 2 => -1, -3, 2
6. 但是如果前面是 4, 3, 2 => -4 < -3 < 2 （貌似2 * -1似乎更好）
7. 在保证前面如果是正数，它必须是递增的，如果是负数，它必须是递降的前提下 分情况讨论
    * a. 前面递增 < a[i] 时，不需要额外的操作
    * b. 前面递增 >= a[i] 时，比如 1 < 3 > 2，a[i] * x 保证a[i]足够的大
    * c. 前面递减 且绝对值 < a[i] 时，前面 * -1 进行翻转
    * d. 前面递减 且绝对值 >= a[i]时，a[i] *= -x，保证新的a[i] 的绝对值更大
8. 这里缺少一个将正数变成负数的机会，且x的取值依赖于前面的变换；且怎么证明这样子能够获得最优解呢？
9. 如果 a[i] != a[i+1], 那么x的取值不重要，这是因为，如果前面的保证是正数的情况下，仍然可以保证它们的相对关系
10. 比如在上面的b、d步里，可以让所有后面的数都*x
11. 所以此时只需要判断相邻数的大小就可以了。相同的数，始终需要一个操作