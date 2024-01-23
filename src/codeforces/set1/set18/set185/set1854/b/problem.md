He has a deck of 𝑛
cards with values 𝑎1,…,𝑎𝑛
from top to bottom. Each card can be either locked or unlocked. Initially, only the topmost card is unlocked.

The game proceeds in turns. In each turn, Andrea chooses an unlocked card in the deck — the value written on the card is
𝑣
— and performs exactly one of the following two operations:

1. Unlock the first 𝑣
   locked cards in the deck from the top. If there are less than 𝑣
   locked cards in the deck, then unlock all the locked cards.
2. Earn 𝑣 victory points.

In both cases, after performing the operation, he removes the card from the deck.
The game ends when all the cards remaining in the deck are locked, or there are no more cards in the deck.

What is the maximum number of victory points Andrea can earn?

### thoughts

1. 操作2在后半段进行处理，不影响结果
2. 所以可以考虑，通过连续的操作1，解锁了前i个cards， 然后用操作2获得分数
3. 要解锁前i个，必须有操作1的cards的分数之和 = i
4. 第一个必须要使用v + (...) = i, 所以这里必须要知道是否存在i - v 的分数
5. 这里，只需要知道n的分数，超出n的部分可以忽略掉，n * n / 64
6. 10**10 / 64 似乎还是太大了