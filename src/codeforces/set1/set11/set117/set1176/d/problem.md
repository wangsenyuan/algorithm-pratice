Authors guessed an array 𝑎
 consisting of 𝑛
 integers; each integer is not less than 2
 and not greater than 2⋅105
. You don't know the array 𝑎
, but you know the array 𝑏
 which is formed from it with the following sequence of operations:

Firstly, let the array 𝑏
 be equal to the array 𝑎
;
Secondly, for each 𝑖
 from 1
 to 𝑛
:
if 𝑎𝑖
 is a prime number, then one integer 𝑝𝑎𝑖
 is appended to array 𝑏
, where 𝑝
 is an infinite sequence of prime numbers (2,3,5,…
);
otherwise (if 𝑎𝑖
 is not a prime number), the greatest divisor of 𝑎𝑖
 which is not equal to 𝑎𝑖
 is appended to 𝑏
;
Then the obtained array of length 2𝑛
 is shuffled and given to you in the input.
Here 𝑝𝑎𝑖
 means the 𝑎𝑖
-th prime number. The first prime 𝑝1=2
, the second one is 𝑝2=3
, and so on.

Your task is to recover any suitable array 𝑎
 that forms the given array 𝑏
. It is guaranteed that the answer exists (so the array 𝑏
 is obtained from some suitable array 𝑎
). If there are multiple answers, you can print any.

Input
The first line of the input contains one integer 𝑛
 (1≤𝑛≤2⋅105
) — the number of elements in 𝑎
.

The second line of the input contains 2𝑛
 integers 𝑏1,𝑏2,…,𝑏2𝑛
 (2≤𝑏𝑖≤2750131
), where 𝑏𝑖
 is the 𝑖
-th element of 𝑏
. 2750131
 is the 199999
-th prime number.

### ideas
1. 2 * n个数，其中的那些质数，要找出来
2. 假设原来有m个质数，那么，最后肯定至少有2 * m 个质数
3. 如果a[i]是质数，那么新加入的也是质数
4. 如果a[i]不是质数，那么最大的除数，有可能是质数, 比如 15, 它最大的除数是5
5. 且如果a[i]是质数，那么加入的p[a[i]], 这个关系是确定的（和i没有关系）
6. 所以，这样子可以找出至少m对
7. 那么剩下的数，它们之间的关系，只能是第二种（也和i没有关系）
8. 所以，一个质数，可以和它的p[x]建立一条线（如果存在的）
9. 一个非质数，和它最大除数建立一条线（如果存在的话）
10. 那么就是最大匹配问题吗？
11. 最大的那个质数，肯定在后面（它如果在前面，必须存在一个更大的质数）
12. 然后，把它对应的数，给删掉，
13. 然后再找最大的质数
14. 剩下的，从前面开始找（因为后面的数更小）