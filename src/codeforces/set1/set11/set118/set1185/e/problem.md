After a hard-working week Polycarp prefers to have fun. Polycarp's favorite entertainment is drawing snakes. He takes a rectangular checkered sheet of paper of size 𝑛×𝑚
 (where 𝑛
 is the number of rows, 𝑚
 is the number of columns) and starts to draw snakes in cells.

Polycarp draws snakes with lowercase Latin letters. He always draws the first snake with the symbol 'a', the second snake with the symbol 'b', the third snake with the symbol 'c' and so on. All snakes have their own unique symbol. There are only 26
 letters in the Latin alphabet, Polycarp is very tired and he doesn't want to invent new symbols, so the total number of drawn snakes doesn't exceed 26
.

Since by the end of the week Polycarp is very tired, he draws snakes as straight lines without bends. So each snake is positioned either vertically or horizontally. Width of any snake equals 1
, i.e. each snake has size either 1×𝑙
 or 𝑙×1
, where 𝑙
 is snake's length. Note that snakes can't bend.

When Polycarp draws a new snake, he can use already occupied cells for drawing the snake. In this situation, he draws the snake "over the top" and overwrites the previous value in the cell.

Recently when Polycarp was at work he found a checkered sheet of paper with Latin letters. He wants to know if it is possible to get this sheet of paper from an empty sheet by drawing some snakes according to the rules described above. If it is possible, he is interested in a way to draw snakes.

### ideas
1. 每个字符的第一个位置和最后一个位置要计算出来
2. 两个位置中间的数字，要 >= 字符