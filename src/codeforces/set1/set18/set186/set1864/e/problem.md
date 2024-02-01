Carol has a sequence 𝑠
of 𝑛
non-negative integers. She wants to play the "Guess Game" with Alice and Bob.

To play the game, Carol will randomly select two integer indices 𝑖𝑎
and 𝑖𝑏
within the range [1,𝑛]
, and set 𝑎=𝑠𝑖𝑎
, 𝑏=𝑠𝑖𝑏
. Please note that 𝑖𝑎
and 𝑖𝑏
may coincide.

Carol will tell:

the value of 𝑎
to Alice;
the value of 𝑏
to Bob;
the value of 𝑎∣𝑏
to both Alice and Bob, where |
denotes the bitwise OR operation.
Please note that Carol will not tell any information about 𝑠
to either Alice or Bob.

Then the guessing starts. The two players take turns making guesses, with Alice starting first. The goal of both players
is to establish which of the following is true: 𝑎<𝑏
, 𝑎>𝑏
, or 𝑎=𝑏
.

In their turn, a player can do one of the following two things:

say "I don't know", and pass the turn to the other player;
say "I know", followed by the answer "𝑎<𝑏
", "𝑎>𝑏
", or "𝑎=𝑏
"; then the game ends.
Alice and Bob hear what each other says, and can use this information to deduce the answer. Both Alice and Bob are smart
enough and only say "I know" when they are completely sure.

You need to calculate the expected value of the number of turns taken by players in the game. Output the answer modulo
998244353
.

Formally, let 𝑀=998244353
. It can be shown that the answer can be expressed as an irreducible fraction 𝑝𝑞
, where 𝑝
and 𝑞
are integers and 𝑞≢0(mod𝑀)
. Output the integer equal to 𝑝⋅𝑞−1mod𝑀
. In other words, output such an integer 𝑥
that 0≤𝑥<𝑀
and 𝑥⋅𝑞≡𝑝(mod𝑀)
.

### thoughts

1. 一脸懵～～
2. let x = a | b
3. 如果a的最高位和x的最高位不一致，alice马上可以回答 知道
4. alice回答不知道，意味着a的最高位和b的最高位是一样的
5. 同样的对于b来说也是如此，所以询问的次数 = 第一个不相同的位在那哪里
6. 如果选定了b，选择不同的a的方式，决定了猜的次数