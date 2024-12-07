Monocarp and Bicarp live in Berland, where every bus ticket consists of 𝑛
 digits (𝑛
 is an even number). During the evening walk Monocarp and Bicarp found a ticket where some of the digits have been erased. The number of digits that have been erased is even.

Monocarp and Bicarp have decided to play a game with this ticket. Monocarp hates happy tickets, while Bicarp collects them. A ticket is considered happy if the sum of the first 𝑛2
 digits of this ticket is equal to the sum of the last 𝑛2
 digits.

Monocarp and Bicarp take turns (and Monocarp performs the first of them). During each turn, the current player must replace any erased digit with any digit from 0
 to 9
. The game ends when there are no erased digits in the ticket.

If the ticket is happy after all erased digits are replaced with decimal digits, then Bicarp wins. Otherwise, Monocarp wins. You have to determine who will win if both players play optimally.

### ideas
1. 如果前后的？的数量不一致，貌似肯定是Monocarp获胜？
2. 考虑前面比后面多，比如前面3个，后面1个
3. 考虑前后？相同的情况
4. 如果前后当前的和相同，那么肯定是Bicarp 获胜。 （他只要复刻Monocarp的操作即可）
5. 如果前后当前的和不同，前面为a，后面为b 且 a < b, 那么肯定是Monocarp获胜，他如果对前面的进行操作，就使用0，对后面操作就使用9
6. 当前面的a < b的时候，对于M来说，前面的？都填0是更优的选择，后面的都甜9是更优的选择