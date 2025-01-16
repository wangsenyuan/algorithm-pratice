Petya and Vasya arranged a game. The game runs by the following rules. Players have a directed graph consisting of n vertices and m edges. One of the vertices contains a chip. Initially the chip is located at vertex s. Players take turns moving the chip along some edge of the graph. Petya goes first. Player who can't move the chip loses. If the game lasts for 106 turns the draw is announced.

Vasya was performing big laboratory work in "Spelling and parts of speech" at night before the game, so he fell asleep at the very beginning of the game. Petya decided to take the advantage of this situation and make both Petya's and Vasya's moves.

Your task is to help Petya find out if he can win the game or at least draw a tie.

### ideas
1. 如果 Petya 可以把chip移动到某个u，u存在一个邻居v，进入v以后无法移动，那么win
2. 或者说，只要Petya能够从s访问到一个无法移动的点v（没有出边）就是win（因为他可以停在v的进入节点上）
3. 如果没有这样的边（就需要进入一个环才行，且这个环在奇数位置上不能离开）draw 
4. 否则肯定是lose