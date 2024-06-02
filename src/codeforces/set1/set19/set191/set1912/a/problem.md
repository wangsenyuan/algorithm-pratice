Allyn is playing a new strategy game called "Accumulator Apex". In this game, Allyn is given the initial value of an integer 𝑥
, referred to as the accumulator, and 𝑘
 lists of integers. Allyn can make multiple turns. On each turn, Allyn can withdraw the leftmost element from any non-empty list and add it to the accumulator 𝑥
 if the resulting 𝑥
 is non-negative. Allyn can end the game at any moment. The goal of the game is to get the largest possible value of the accumulator 𝑥
. Please help Allyn find the largest possible value of the accumulator 𝑥
 they can get in this game.

 ### ideas
 1. max prefix sum?
 2. 还有个要求，是在任何一步操作后，都不能出现负数
 3. 那这样子就有难度了
 4. max prefix sum还有用吗？
 5. 如果任何一个队列的head是正数，直接加上去就可了
 6. ok，现在所有的队列head都是负数，那似乎也应该加上绝对值最小的负数
 7. 好像这个策略是ok的？
 8. 不对的。可以考虑这样一个情况，就是 [-2, -3, 6], [-1, -2]
 9. 假设现在x = 5， 如果只使用第一个list，那么收益 = 6, 但是如果使用上面的算法，把-1加入后，就无法达到收益6
 10. 所以这里还有一个因素要考虑，就是在保证当前，且保证x >= 0 的前提下，可以得到的最大收益
 11. 