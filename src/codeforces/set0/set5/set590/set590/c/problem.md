The famous global economic crisis is approaching rapidly, so the states of Berman, Berance and Bertaly formed an alliance and allowed the residents of all member states to freely pass through the territory of any of them. In addition, it was decided that a road between the states should be built to guarantee so that one could any point of any country can be reached from any point of any other State.

Since roads are always expensive, the governments of the states of the newly formed alliance asked you to help them assess the costs. To do this, you have been issued a map that can be represented as a rectangle table consisting of n rows and m columns. Any cell of the map either belongs to one of three states, or is an area where it is allowed to build a road, or is an area where the construction of the road is not allowed. A cell is called passable, if it belongs to one of the states, or the road was built in this cell. From any passable cells you can move up, down, right and left, if the cell that corresponds to the movement exists and is passable.

Your task is to construct a road inside a minimum number of cells, so that it would be possible to get from any cell of any state to any cell of any other state using only passable cells.

It is guaranteed that initially it is possible to reach any cell of any state from any cell of this state, moving only along its cells. It is also guaranteed that for any state there is at least one cell that belongs to it.

### ideas
1. 一开始，同一个state中的所有cell，应该是连接在一起的
2. 通过空的cell上去连接，这3个compoent，使得使用的cell最少
3. 这里有2种case，一种是，它们排成一排，1 -> 2 -> 3
4. 另外一种是，它们组成一个start, 1 -> 0, 2 -> 0, 3 -> 0
5. 第一种比较好处理（使用3次多源bfs即可）
6. 也是可以处理的