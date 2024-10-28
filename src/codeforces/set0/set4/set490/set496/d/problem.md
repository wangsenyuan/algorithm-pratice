Petya and Gena love playing table tennis. A single match is played according to the following rules: a match consists of multiple sets, each set consists of multiple serves. Each serve is won by one of the players, this player scores one point. As soon as one of the players scores t points, he wins the set; then the next set starts and scores of both players are being set to 0. As soon as one of the players wins the total of s sets, he wins the match and the match is over. Here s and t are some positive integer numbers.

To spice it up, Petya and Gena choose new numbers s and t before every match. Besides, for the sake of history they keep a record of each match: that is, for each serve they write down the winner. Serve winners are recorded in the chronological order. In a record the set is over as soon as one of the players scores t points and the match is over as soon as one of the players wins s sets.

Petya and Gena have found a record of an old match. Unfortunately, the sequence of serves in the record isn't divided into sets and numbers s and t for the given match are also lost. The players now wonder what values of s and t might be. Can you determine all the possible options?

### ideas
1. s是获胜的set的数量，t是单个set中获胜的点数
2. 1, 1, 1, 1, 1 那么 s = 1， t = 5, 或者 s = 5, t = 1
3. 获胜多的那个人貌似是固定的，肯定是最后一个t = 1/2
4. 因为如果不是他，那么要么之前就结束了，要么还要继续比下去
5. 知道x谁获胜了，下一步就是倒推？
6. t肯定 = 最后一段是x占多数的一段（t可以确认）
7. 然后就是验证t是否可行
8. 但是这样子的问题是，对于每个t，如果能够在n/t内验证掉，就能pass
9. 