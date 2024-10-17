For her birthday Alice received an interesting gift from her friends – The Light Square. The Light Square game is played on an 𝑁×𝑁
 lightbulbs square board with a magical lightbulb bar of size 𝑁×1
 that has magical properties. At the start of the game some lights on the square board and magical bar are turned on. The goal of the game is to transform the starting light square board pattern into some other pattern using the magical bar without rotating the square board. The magical bar works as follows:

It can be placed on any row or column

The orientation of the magical lightbulb must be left to right or top to bottom for it to keep its magical properties

The entire bar needs to be fully placed on a board

The lights of the magical bar never change

If the light on the magical bar is the same as the light of the square it is placed on it will switch the light on the square board off, otherwise it will switch the light on

The magical bar can be used an infinite number of times

Alice has a hard time transforming her square board into the pattern Bob gave her. Can you help her transform the board or let her know it is impossible? If there are multiple solutions print any.

### ideas
1. 考察位置(i, j), 如果 s[i][j] != t[i][j], 
2. 那么必须，bar[i] = 1 or bar[j] = 1 (否则这个地方无法被调整)
3. 且只能被垂直或者水平调整一次（如果 bar[i] = 1, 就可以使用垂直调整）
4. 如果s[i][j] = t[i][j], 如果 bar[i] = 1, 那么就不能使用垂直调整，
5.  ....
6. 所以，就是行r和列c之间连线，如果 s[r][c] != t[r][c], 
7.   如果 bar[r] = 1 and bar[c] = 1, c => !r, r => !c 如果在c列进行了调整，就不能在r行调整
8.   如果 bar[r] = 1 and bar[c] = 0, （那么 c比如为true, 必须调整c列， 对r没有要求）
9.   如果 bar[r] = 0 and bar[c] = 1, 那么r必须为true， 必须被调整，对c没有要求
10.   Else => false
11. s[r][c] == t[r][c] 
12.   如果 bar[r] = 1 and bar[c] = 1, r => c, !c => !r (r和c必须同时调整、或者同时不调整)
13.   如果 bar[r] = 1, and bar[c] = 0, 那么 c必须为false
14.   。。 bar[r] = 0, and bar[c] = 1, 那么 r必须为false
15.   。。 bar[r] = 0, and bar[c] = 0, 对r, c没有要求
16. 这样构成一个图后，看是否有冲突，没有冲突就有解
17. 感觉是个很麻烦的东西～
18.  
