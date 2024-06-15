You are given an 𝑛×𝑛
 chessboard where you and the computer take turns alternatingly to place white rooks & black rooks on the board respectively. While placing rooks, you have to ensure that no two rooks attack each other. Two rooks attack each other if they share the same row or column regardless of color.

A valid move is placing a rook on a position (𝑟
, 𝑐
) such that it doesn't attack any other rook.

You start first, and when you make a valid move in your turn, placing a white rook at position (𝑟
, 𝑐
), the computer will mirror you and place a black rook at position (𝑐
, 𝑟
) in its turn. If 𝑟=𝑐
, then the computer can't mirror your move, and skips its turn.

You have already played 𝑘
 moves with the computer (the computer tries to mirror these moves too), and you must continue playing the game until there are no valid moves remaining. How many different final configurations are possible when you continue the game after the 𝑘
 moves? It is guaranteed that the 𝑘
 moves and the implied computer moves are valid. Since the answer may be large, print it modulo 109+7
.

Two configurations are considered different if there exists a coordinate (𝑟
, 𝑐
) which has a rook in one configuration, but not in the other or the color of the rook on the coordinate is different.

### ideas
1. 不知道怎么搞
2. 如果alice 放在某个 (i, i)上，那么bob相当于也放在了(i, i)上
3. 对于所有其他的位置 (i, j), bob都放在位置(j, i)上
4. 所以，考虑k=0的时候，alice如何放，首先他有n次机会，
5. 假设第一次他放在位置(1, x), 那么后续他就不能再放置在(?, x)上面
6. 首先它可以对(1, ...n) 排个序，然后第一次，他有n个选择，第二次，只有(n-1)个选择
7. 依次类推，也就是 (n!) * (n!)
8. 如果考虑首先已经
9. 不对，对角线上的，比较特殊，
10. C(n, i) 表示有多少个对角线，(n - i)!