You and your 𝑛−1
 friends have found an array of integers 𝑎1,𝑎2,…,𝑎𝑛
. You have decided to share it in the following way: All 𝑛
 of you stand in a line in a particular order. Each minute, the person at the front of the line chooses either the first or the last element of the array, removes it, and keeps it for himself. He then gets out of line, and the next person in line continues the process.

You are standing in the 𝑚
-th position in the line. Before the process starts, you may choose up to 𝑘
 different people in the line, and persuade them to always take either the first or the last element in the array on their turn (for each person his own choice, not necessarily equal for all people), no matter what the elements themselves are. Once the process starts, you cannot persuade any more people, and you cannot change the choices for the people you already persuaded.

Suppose that you're doing your choices optimally. What is the greatest integer 𝑥
 such that, no matter what are the choices of the friends you didn't choose to control, the element you will take from the array will be greater than or equal to 𝑥
?

Please note that the friends you don't control may do their choice arbitrarily, and they will not necessarily take the biggest element available.

### ideas
1. 如果k >= m - 1, 就可以获得前m个、后m个中的最大值
2. 现在 k < m - 1,
3. 假设player想获得a[i], 那么他需要保证多少个人听他的才行呢？
4. 一头懵
5. 这k个人，肯定是头部的那些 dp[l][r] = player能够获得的最大值 l + n - r - 1 == k
6. 选定l（r就确定了）然后再l...r中间，从k+1...m 开始选择，能够得到的最小值
7. i - l + r - j = m - k， max(a[i], a[j])
8. 要取这个里面的最小值