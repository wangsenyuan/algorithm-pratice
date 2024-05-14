Sasha likes programming. Once, during a very long contest, Sasha decided that he was a bit tired and needed to relax. So he did. But since Sasha isn't an ordinary guy, he prefers to relax unusually. During leisure time Sasha likes to upsolve unsolved problems because upsolving is very useful.

Therefore, Sasha decided to upsolve the following problem:

You have an array 𝑎
 with 𝑛
 integers. You need to count the number of funny pairs (𝑙,𝑟)
 (𝑙≤𝑟)
. To check if a pair (𝑙,𝑟)
 is a funny pair, take 𝑚𝑖𝑑=𝑙+𝑟−12
, then if 𝑟−𝑙+1
 is an even number and 𝑎𝑙⊕𝑎𝑙+1⊕…⊕𝑎𝑚𝑖𝑑=𝑎𝑚𝑖𝑑+1⊕𝑎𝑚𝑖𝑑+2⊕…⊕𝑎𝑟
, then the pair is funny. In other words, ⊕
 of elements of the left half of the subarray from 𝑙
 to 𝑟
 should be equal to ⊕
 of elements of the right half. Note that ⊕
 denotes the bitwise XOR operation.

It is time to continue solving the contest, so Sasha asked you to solve this task.


### idea
1. 当l...r的长度为偶数时， (l + r - 1) / 2 = (l + r) / 2
2. 正好前面一半，后面一半
3. 假设知道中点mid, 且mid+1...r 的xor = x 那么如果a[l][mid] = x, 那么 (l...r)就是一个答案
4. 显然不能这么去算
5. 这里可以确定的一个点时， a[l] ^ .... ^ a[r] = 0
6. 所以有 s[r] = s[l-1]
7. xor时，这个式子一定时成立的，唯一需要确定的是，r和 l-1 的奇偶性要相同， 才能保证 l...r 的长度为偶数