There are positive integers 
N and 
M, where 
N<M.
Alice and Bob will play a game. Each player has 
N cards with 
1,2,…,N written on them, one for each number. Starting with Alice, the two players take turns repeatedly performing this action: choose one card from their hand and play it onto the table.
Immediately after a card is played onto the table, if the sum of the numbers on the cards that have been played so far is divisible by 
M, the player who just played the card loses, and the other player wins.
If both players play all their cards without satisfying the above condition, Alice wins.
Who will win, Alice or Bob, when both play optimally?

### ideas
1. 这里面不变的是什么呢？
2. 假设当前在桌子上的排的sum (% m)， 如果当前player，手里只剩下一站牌，且 sum + i % m = 0
3. 那么当前player lose；否则有两张牌（或者更多），当前player总有不会马上输掉的办法
4. 如果 2 * n * (n + 1) / 2 % m = i
5. 如果 i <= n, 那么alice貌似肯定会输掉？因为bob只要把i拿在手里，就可以保证alice先输掉