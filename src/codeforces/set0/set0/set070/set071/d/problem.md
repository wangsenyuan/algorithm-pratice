Vasya has a pack of 54 cards (52 standard cards and 2 distinct jokers). That is all he has at the moment. Not to die from boredom, Vasya plays Solitaire with them.

Vasya lays out nm cards as a rectangle n × m. If there are jokers among them, then Vasya should change them with some of the rest of 54 - nm cards (which are not layed out) so that there were no jokers left. Vasya can pick the cards to replace the jokers arbitrarily. Remember, that each card presents in pack exactly once (i. e. in a single copy). Vasya tries to perform the replacements so that the solitaire was solved.

Vasya thinks that the solitaire is solved if after the jokers are replaced, there exist two non-overlapping squares 3 × 3, inside each of which all the cards either have the same suit, or pairwise different ranks.

Determine by the initial position whether the solitaire can be solved or not. If it can be solved, show the way in which it is possible.

Input
The first line contains integers n and m (3 ≤ n, m ≤ 17, n × m ≤ 52). Next n lines contain m words each. Each word consists of two letters. The jokers are defined as "J1" and "J2" correspondingly. For the rest of the cards, the first letter stands for the rank and the second one — for the suit. The possible ranks are: "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K" and "A". The possible suits are: "C", "D", "H" and "S". All the cards are different.

Output
If the Solitaire can be solved, print on the first line "Solution exists." without the quotes. On the second line print in what way the jokers can be replaced. Three variants are possible:

"There are no jokers.", if there are no jokers in the input data.
"Replace Jx with y.", if there is one joker. x is its number, and y is the card it should be replaced with.
"Replace J1 with x and J2 with y.", if both jokers are present in the input data. x and y here represent distinct cards with which one should replace the first and the second jokers correspondingly.
On the third line print the coordinates of the upper left corner of the first square 3 × 3 in the format "Put the first square to (r, c).", where r and c are the row and the column correspondingly. In the same manner print on the fourth line the coordinates of the second square 3 × 3 in the format "Put the second square to (r, c).".

If there are several solutions to that problem, print any of them.

If there are no solutions, print of the single line "No solution." without the quotes.

### ideas
1. 两个3 * 3 的区域，要么是同一个花色，要么是所有的数字不同
2. 对于位置(r, c)检查是否可以在不交换，或者交换 joker的情况下，是否是good（怎么good，似乎没有关系）
3. 但是交换了哪个是有关系的
4. dp[r][c][0] = flag 表示交换J1去flag表示的那些牌时，是否good
5. dp[r][c][1] = flag 表示交换J2。。。
6. dp[r1][c1][0/1] and dp[r2][c2][0/1] 与的结果，至少有两位