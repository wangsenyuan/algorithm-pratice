Having heard of that principle, but having not mastered the technique of logical thinking, 8 year olds Stas and Masha invented a game. There are a different boxes and b different items, and each turn a player can either add a new box or a new item. The player, after whose turn the number of ways of putting b items into a boxes becomes no less then a certain given number n, loses

### ideas
1. a 个 box， b个items， 把它们放进去
2. pow(a, b) ?
3. 增加一个box => (a, b) => (a + 1, b) = pow(a + 1, b)
4. 增加一个item => (a, b) => (a, b + 1)
5. 当前是(a, b), 它是一个lose if f(a + 1, b) win or f(a, b + 1) win
6. f(a, b) = win if f(a+1, b) lose and f(a, b + 1) lose
7. 但是当a = 1 的时候呢？至少是draw
8. 如果 f(2, b) 是lose， win， else draw 