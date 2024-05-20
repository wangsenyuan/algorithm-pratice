Two players are playing an online card game. The game is played using a 32-card deck. Each card has a suit and a rank. There are four suits: clubs, diamonds, hearts, and spades. We will encode them with characters 'C', 'D', 'H', and 'S', respectively. And there are 8 ranks, in increasing order: '2', '3', '4', '5', '6', '7', '8', '9'.

Each card is denoted by two letters: its rank and its suit. For example, the 8 of Hearts is denoted as 8H.

At the beginning of the game, one suit is chosen as the trump suit.

In each round, players make moves like this: the first player places one of his cards on the table, and the second player must beat this card with one of their cards. After that, both cards are moved to the discard pile.

A card can beat another card if both cards have the same suit and the first card has a higher rank than the second. For example, 8S can beat 4S. Additionally, a trump card can beat any non-trump card, regardless of the rank of the cards, for example, if the trump suit is clubs ('C'), then 3C can beat 9D. Note that trump cards can be beaten only by the trump cards of higher rank.

There were 𝑛
 rounds played in the game, so the discard pile now contains 2𝑛
 cards. You want to reconstruct the rounds played in the game, but the cards in the discard pile are shuffled. Find any possible sequence of 𝑛
 rounds that might have been played in the game.

¸


### ideas
1. n <= 16， 2 * n <= 32 
2. 这些card可以组成一个图，然后要能够分成白/黑两种颜色
3. 白色先出，黑色beat it
4. 最大匹配问题
5. 不是最大匹配问题