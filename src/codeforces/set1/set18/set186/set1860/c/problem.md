Alice and Bob are playing a game. They have a permutation 𝑝
of size 𝑛
(a permutation of size 𝑛
is an array of size 𝑛
where each element from 1
to 𝑛
occurs exactly once). They also have a chip, which can be placed on any element of the permutation.

Alice and Bob make alternating moves: Alice makes the first move, then Bob makes the second move, then Alice makes the
third move, and so on. During the first move, Alice chooses any element of the permutation and places the chip on that
element. During each of the next moves, the current player has to move the chip to any element that is simultaneously to
the left and strictly less than the current element (i. e. if the chip is on the 𝑖
-th element, it can be moved to the 𝑗
-th element if 𝑗<𝑖
and 𝑝𝑗<𝑝𝑖
). If a player cannot make a move (it is impossible to move the chip according to the rules of the game), that player
wins the game.

Let's say that the 𝑖
-th element of the permutation is lucky if the following condition holds:

if Alice places the chip on the 𝑖
-th element during her first move, she can win the game no matter how Bob plays (i. e. she has a winning strategy).
You have to calculate the number of lucky elements in the permutation.