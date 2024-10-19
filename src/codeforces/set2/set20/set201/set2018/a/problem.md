You have some cards. An integer between 1
 and 𝑛
 is written on each card: specifically, for each 𝑖
 from 1
 to 𝑛
, you have 𝑎𝑖
 cards which have the number 𝑖
 written on them.

There is also a shop which contains unlimited cards of each type. You have 𝑘
 coins, so you can buy at most 𝑘
 new cards in total, and the cards you buy can contain any integer between 1
 and 𝐧
, inclusive.

After buying the new cards, you must partition all your cards into decks, according to the following rules:

all the decks must have the same size;
there are no pairs of cards with the same value in the same deck.
Find the maximum possible size of a deck after buying cards and partitioning them optimally.

### ideas
1. 有点乱
2. 首先可以确定的是，最少要cnt = max(a[i])个set
3. 假设每个set有x个, 那么有 x * cnt <= sum(a[i]) + k
4. 首先买已有的card没有好处（除非，已有的已经全买到了），因为这样子不会增加size
5. 所以，可以考虑买到的都是新的card
6. x = (sum + k) / cnt ?
7. 且 cnt <= n
8. 且 sum + k <= n * n（不可能超过这个数）
9. 这里有个诡异的地方，就是要把所有的牌都分配掉
10. 假设存在一个 a[i] > n, 那么至少有a[i]堆
11. 