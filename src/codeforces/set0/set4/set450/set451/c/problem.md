There are n games in a football tournament. Three teams are participating in it. Currently k games had already been
played.

You are an avid football fan, but recently you missed the whole k games. Fortunately, you remember a guess of your
friend for these k games. Your friend did not tell exact number of wins of each team, instead he thought that absolute
difference between number of wins of first and second team will be d1 and that of between second and third team will be
d2.

You don't want any of team win the tournament, that is each team should have the same number of wins after n games.
That's why you want to know: does there exist a valid tournament satisfying the friend's guess such that no team will
win this tournament?

Note that outcome of a match can not be a draw, it has to be either win or loss.

### ideas

1. n % 3 = 0 否则不能满足它们有相同的获胜场次
2. 假设到目前为止，分别获胜 x, y, z 那么有 abs(x - y) = d1, abs(y - z) = d2
3. x + y + z = k
4. 假设 y > z
5. 如果 y > x => y = x + d1 = z + d2
6. y = k - (x + z)
7. 2 * y = x + z + d1 + d2
8. 3 * y = k + d1 + d2 => y
9. 如果 y < x => y = x - d1 = z + d2 => 2 * y = x + z - d1 + d2 => 3 * y = k - d1 + d2
10. 如果 y < z
11. 如果y > x =>