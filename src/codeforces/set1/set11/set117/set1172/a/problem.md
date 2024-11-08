Nauuo is a girl who loves playing cards.

One day she was playing cards but found that the cards were mixed with some empty ones.

There are 𝑛
 cards numbered from 1
 to 𝑛
, and they were mixed with another 𝑛
 empty cards. She piled up the 2𝑛
 cards and drew 𝑛
 of them. The 𝑛
 cards in Nauuo's hands are given. The remaining 𝑛
 cards in the pile are also given in the order from top to bottom.

In one operation she can choose a card in her hands and play it — put it at the bottom of the pile, then draw the top card from the pile.

Nauuo wants to make the 𝑛
 numbered cards piled up in increasing order (the 𝑖
-th card in the pile from top to bottom is the card 𝑖
) as quickly as possible. Can you tell her the minimum number of operations?

### ideas
1. 考虑每次操作
2. 就是把手中的某张牌x,放到最后，并把头部的牌y抽到手中
3. 目标是使的桌子上的牌，由1，2，。。。。n组成
4. 考虑最后一次操作，必然是把手中的n，放到最后，并从开始拿走0
5. 首先，有一个2*n步的方式
6. 就是先把数字牌，都收集到手中，如果头部是数字，那么手中肯定有0牌，把0牌放在尾巴上，然后把数字牌拿到手中
7. 这样子，可以在最多n步内，把所有的数字牌拿到手中
8. 然后再用数字牌去换0牌就可以了
9. 如果一开始的状态是???? 1, 2, 3... i
10. 如果i+1在手中，就把i+1换过去， 然后把头部的牌拿到手中
11. 但是如果 i+1在前面的?中呢？就麻烦了
12. 所以，如果这个策略不可以，就用2*n的策略就好了
13. 也不对
14. 其实1可以在比较早的时候就放进去
15. 然后放置2，3.。。。n
16. 也就是说，假设能够在某一步的时候开始放置1，然后2.。。。i
17. 那么在放置完i后，i+1必须在手中， 依次类推
18. 我们考虑1什么时候在手中？
19. 就是看1，2，3.。。。x这些牌什么时候可以在手中
20. 结果 = max(f(i)) + n - i + 1