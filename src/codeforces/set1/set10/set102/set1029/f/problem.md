There is an infinite board of square tiles. Initially all tiles are white.

Vova has a red marker and a blue marker. Red marker can color 𝑎
tiles. Blue marker can color 𝑏
tiles. If some tile isn't white then you can't use marker of any color on it. Each marker must be drained completely, so
at the end there should be exactly 𝑎
red tiles and exactly 𝑏
blue tiles across the board.

Vova wants to color such a set of tiles that:

they would form a rectangle, consisting of exactly 𝑎+𝑏
colored tiles;
all tiles of at least one color would also form a rectangle.
Here are some examples of correct colorings:

### ideas

1. w * h = a + b
2. w1 * h1 = a , w1 <= w && h1 <= h
3. or w2 * h2 = b, w2 <= w && h2 <= h
4. brute force => get factors O(sqrt(1e14)) = 1e7
5. 然后判断就好了