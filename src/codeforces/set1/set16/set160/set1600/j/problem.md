https://codeforces.com/problemset/problem/1600/J

### problem

You have received data from a Bubble bot. You know your task is to make factory facilities, but before you even start,
you need to know how big the factory is and how many rooms it has. When you look at the data you see that you have the
dimensions of the construction, which is in rectangle shape: N x M.

Then in the next N lines you have M numbers. These numbers represent factory tiles and they can go from 0 to 15. Each of
these numbers should be looked in its binary form. Because from each number you know on which side the tile has walls.
For example number 10 in it's binary form is 1010, which means that it has a wall from the North side, it doesn't have a
wall from the East, it has a wall on the South side and it doesn't have a wall on the West side. So it goes North, East,
South, West.

It is guaranteed that the construction always has walls on it's edges. The input will be correct.

Your task is to print the size of the rooms from biggest to smallest.

### ideas

1. dfs? or union-find?