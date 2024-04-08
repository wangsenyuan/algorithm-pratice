Ilya is working for the company that constructs robots. Ilya writes programs for entertainment robots, and his current
project is "Bob", a new-generation game robot. Ilya's boss wants to know his progress so far. Especially he is
interested if Bob is better at playing different games than the previous model, "Alice".

So now Ilya wants to compare his robots' performance in a simple game called "1-2-3". This game is similar to the "
Rock-Paper-Scissors" game: both robots secretly choose a number from the set {1, 2, 3} and say it at the same moment. If
both robots choose the same number, then it's a draw and noone gets any points. But if chosen numbers are different,
then one of the robots gets a point: 3 beats 2, 2 beats 1 and 1 beats 3.

Both robots' programs make them choose their numbers in such a way that their choice in (i + 1)-th game depends only on
the numbers chosen by them in i-th game.

Ilya knows that the robots will play k games, Alice will choose number a in the first game, and Bob will choose b in the
first game. He also knows both robots' programs and can tell what each robot will choose depending on their choices in
previous game. Ilya doesn't want to wait until robots play all k games, so he asks you to predict the number of points
they will have after the final game.

### ideas

1. 显然要找到经过多少次后，重新回到了起点
2. 这个应该是很容易找到的，就9种可能性，