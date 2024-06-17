Sasha gave Anna a list 𝑎
 of 𝑛
 integers for Valentine's Day. Anna doesn't need this list, so she suggests destroying it by playing a game.

Players take turns. Sasha is a gentleman, so he gives Anna the right to make the first move.

On her turn, Anna must choose an element 𝑎𝑖
 from the list and reverse the sequence of its digits. For example, if Anna chose the element with a value of 42
, it would become 24
; if Anna chose the element with a value of 1580
, it would become 851
. Note that leading zeros are removed. After such a turn, the number of elements in the list does not change.
On his turn, Sasha must extract two elements 𝑎𝑖
 and 𝑎𝑗
 (𝑖≠𝑗
) from the list, concatenate them in any order and insert the result back into the list. For example, if Sasha chose the elements equal to 2007
 and 19
, he would remove these two elements from the list and add the integer 200719
 or 192007
. After such a turn, the number of elements in the list decreases by 1
.
Players can't skip turns. The game ends when Sasha can't make a move, i.e. after Anna's move there is exactly one number left in the list. If this integer is not less than 10𝑚
 (i.e., ≥10𝑚
), Sasha wins. Otherwise, Anna wins.

It can be shown that the game will always end. Determine who will win if both players play optimally.