Amr bought a new video game "Guess Your Way Out!". The goal of the game is to find an exit from the maze that looks like a perfect binary tree of height h. The player is initially standing at the root of the tree and the exit from the tree is located at some leaf node.

Let's index all the leaf nodes from the left to the right from 1 to 2h. The exit is located at some node n where 1 ≤ n ≤ 2h, the player doesn't know where the exit is so he has to guess his way out!

Amr follows simple algorithm to choose the path. Let's consider infinite command string "LRLRLRLRL..." (consisting of alternating characters 'L' and 'R'). Amr sequentially executes the characters of the string using following rules:

Character 'L' means "go to the left child of the current node";
Character 'R' means "go to the right child of the current node";
If the destination node is already visited, Amr skips current command, otherwise he moves to the destination node;
If Amr skipped two consecutive commands, he goes back to the parent of the current node before executing next command;
If he reached a leaf node that is not the exit, he returns to the parent of the current node;
If he reaches an exit, the game is finished.
Now Amr wonders, if he follows this algorithm, how many nodes he is going to visit before reaching the exit?

### ideas
1. 是一个深度优先的过程
2. 搞错了～