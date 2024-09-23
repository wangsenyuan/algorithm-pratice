When Kefa came to the restaurant and sat at a table, the waiter immediately brought him the menu. There were n dishes. Kefa knows that he needs exactly m dishes. But at that, he doesn't want to order the same dish twice to taste as many dishes as possible.

Kefa knows that the i-th dish gives him ai units of satisfaction. But some dishes do not go well together and some dishes go very well together. Kefa set to himself k rules of eating food of the following type — if he eats dish x exactly before dish y (there should be no other dishes between x and y), then his satisfaction level raises by c.

Of course, our parrot wants to get some maximal possible satisfaction from going to the restaurant. Help him in this hard task!

### ideas
1. dp[state][i] 最后一个是i的，集合state所能带来的最大值