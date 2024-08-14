A palindrome is a string that reads the same left to right and right to left. For example, "101101" is a palindrome, while "0101" is not.

Alice and Bob are playing a game on a string 𝑠
 of length 𝑛
 consisting of the characters '0' and '1'. Both players take alternate turns with Alice going first.

In each turn, the player can perform one of the following operations:

Choose any 𝑖
 (1≤𝑖≤𝑛
), where 𝑠[𝑖]=
 '0' and change 𝑠[𝑖]
 to '1'. Pay 1 dollar.
Reverse the whole string, pay 0 dollars. This operation is only allowed if the string is currently not a palindrome, and the last operation was not reverse. That is, if Alice reverses the string, then Bob can't reverse in the next move, and vice versa.
Reversing a string means reordering its letters from the last to the first. For example, "01001" becomes "10010" after reversing.

The game ends when every character of string becomes '1'. The player who spends minimum dollars till this point wins the game and it is a draw if both spend equal dollars. If both players play optimally, output whether Alice wins, Bob wins, or if it is a draw.

### ideas
1. 首先，reverse这步貌似就是skip，表示当前用户，不操作（当然也不用付钱）
2. 如果一开始就是回文（那就不能skip）
3. 如果一开始不是回文，那就可以skip
4. 假设当前状态是个lose的状态，那么通过skip，可以翻转状态
5. 假设当前状态不是回文，那么是不是skip肯定是一个更优的策略呢？
6. 假设alice就采取skip，（bob就必须付钱，把0变成1）
7. 直到出现第一个回文，假设这时候花费了x元；
8. alice必须付钱操作，然后bobskip，alice再操作一次，得到一个回文。bob必须付钱。。。
9. 假设alice付钱操作后， bob不skip，反而花了一次钱；那么alice就可以skip了。
10. 首先，所有要花的钱数是确定的。这个数量 = 0的个数；s
11. 然后第一次出现回文时花费的最小的钱数也是确定的（中心点一会再考虑）；x
12. 然后剩余 y = s - x（如果存在中心点，先假设中心点此时为1）
13. 那么y必然能整除2， 也就是alice需要付的钱 = (y+1) / 2 * 2, 而bob要付的钱 = y/2 * 2 + x
14. 然后考虑中心点的情况，如果中心点=0， 那么就有一次把
15. 而这个决定权在bob手中。
16. 这个我们可以放到最后取。 如果 y / 2 是个奇数，那么这个中心点，还是要bob处理；如果是个偶数，那么bob可以跳过，让alice处理