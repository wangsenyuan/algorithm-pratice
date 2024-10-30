A boy named Vasya wants to play an old Russian solitaire called "Accordion". In this solitaire, the player must observe the following rules:

A deck of n cards is carefully shuffled, then all n cards are put on the table in a line from left to right;
Before each move the table has several piles of cards lying in a line (initially there are n piles, each pile has one card). Let's number the piles from left to right, from 1 to x. During one move, a player can take the whole pile with the maximum number x (that is the rightmost of remaining) and put it on the top of pile x - 1 (if it exists) or on the top of pile x - 3 (if it exists). The player can put one pile on top of another one only if the piles' top cards have the same suits or values. Please note that if pile x goes on top of pile y, then the top card of pile x becomes the top card of the resulting pile. Also note that each move decreases the total number of piles by 1;
The solitaire is considered completed if all cards are in the same pile.
Vasya has already shuffled the cards and put them on the table, help him understand whether completing this solitaire is possible or not.

### ideas
1. n, 要么移动到n-1, 要么移动到n-3
2. 可以处理了，每个位置，记录value+suite的状态