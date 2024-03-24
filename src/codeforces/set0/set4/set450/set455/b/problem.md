Andrew, Fedor and Alex are inventive guys. Now they invent the game with strings for two players.

Given a group of n non-empty strings. During the game two players build the word together, initially the word is empty.
The players move in turns. On his step player must add a single letter in the end of the word, the resulting word must
be prefix of at least one string from the group. A player loses if he cannot move.

Andrew and Alex decided to play this game k times. The player who is the loser of the i-th game makes the first move in
the (i + 1)-th game. Guys decided that the winner of all games is the player who wins the last (k-th) game. Andrew and
Alex already started the game. Fedor wants to know who wins the game if both players will play optimally. Help him.

### ideas

1. 考虑k=1，也就是一次决胜负
2. 考虑由字母组成的trie树，每个节点标识（对于当前player）是胜利/失败
3. 所有末端的节点都是胜利（因为当前player完成后，另外一个player无法进行）
4. 对于中间的节点，如果所有的后继节点都是胜利，那么它是失败，否则它就是胜利（只要有一个失败的后继，它就是胜利）
5. 通过这种方式，可以知道所有节点的胜利和失败情况
6. 如果k=1非常容易处理，对于k>1的情况，我们考虑第一个人的选择，
7. 如果下次不是最后一次，他希望自己能够失败掉，这样他下次还可以进行选择
8. 也就是说，对于Alice来说，她首先要有必胜的起始状态，且必须有一个必输的起始状态
9. 但是必输，就是走到它这里，后续的都是胜利节点
10. 如果全部是胜利的节点，再看k的奇偶性