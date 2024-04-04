Initially, Ithea puts n clear sheets of paper in a line. They are numbered from 1 to n from left to right.

This game will go on for m rounds. In each round, Ithea will give Chtholly an integer between 1 and c, and Chtholly
needs to choose one of the sheets to write down this number (if there is already a number before, she will erase the
original one and replace it with the new one).

Chtholly wins if, at any time, all the sheets are filled with a number and the n numbers are in non-decreasing order
looking from left to right from sheet 1 to sheet n, and if after m rounds she still doesn't win, she loses the game.

Chtholly really wants to win the game as she wants to cook something for Willem. But she doesn't know how to win the
game. So Chtholly finds you, and your task is to write a program to receive numbers that Ithea gives Chtholly and help
her make the decision on which sheet of paper write this number.

### ideas

1. 这里有m个数，如果全部read完，当然很好处理，但问题出在不能全部读完，就要给出答案
2. 如果c=2, 那么非常简单， 遇到1放到头部，遇到2，放到尾部
3. 这里有个想法，就是把n张sheet分成 (c/2)段
4. 如果遇到某个数x，就把它放到对应的段内，可能会出现这样一个情况，比如c=4，分成2段，
5. 但是1特别多，而3只出现了一次，4根本没有出现