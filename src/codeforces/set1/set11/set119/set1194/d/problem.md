Alice and Bob play a game. There is a paper strip which is divided into n + 1 cells numbered from left to right starting from 0. There is a chip placed in the n-th cell (the last one).

Players take turns, Alice is first. Each player during his or her turn has to move the chip 1, 2 or k cells to the left (so, if the chip is currently in the cell i, the player can move it into cell i - 1, i - 2 or i - k). The chip should not leave the borders of the paper strip: it is impossible, for example, to move it k cells to the left if the current cell has number i < k. The player who can't make a move loses the game.

Who wins if both participants play optimally?

Alice and Bob would like to play several games, so you should determine the winner in each game.


### ideas
1. 0... n
2. 如果Alice到达了位置1,2, k就输了(没有3...k-1)
3. 假设到达了k-1， 这时，只能使用1/2操作
4. 比如k-1 = 5
5. 如果k-1 >= 3, 似乎alice始终可以获胜？
6. 如果alice 到达了3这个位置，那么bob只能选择1/2,下一步alice就获胜了
7. 如果alice在3后面的位置，但是k的前面，那么这时候假设alice先到了位置4，那么他可以选择1，
8. 如果先到了位置5， 那么他可以使用2，如果到了6，似乎选择权就到了bob手中
9. 如果n > k， m = n / k, r = n % k
10. m-= 1, 先考虑r，有没有可能让对方一定得到k这个位置？
11. 如果r = 0, 也就是一开始就在k的倍数位置
12. 如果 m = 0, alice获胜，
13. m > 0, 如果m 是偶数的话，且 k % 3 = 0, 且 k / 3 是偶数，那么alice获胜
14. 在这种情况下， alice能保证始终在奇数倍数的k的位置
15. 如果 m是偶数，且r / 3 是偶数，那么alice获胜，其他情况下，应该是bob获胜