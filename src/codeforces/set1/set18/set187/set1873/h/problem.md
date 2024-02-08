Marcel and Valeriu are in the mad city, which is represented by 𝑛
buildings with 𝑛
two-way roads between them.

Marcel and Valeriu start at buildings 𝑎
and 𝑏
respectively. Marcel wants to catch Valeriu, in other words, be in the same building as him or meet on the same road.

During each move, they choose to go to an adjacent building of their current one or stay in the same building. Because
Valeriu knows Marcel so well, Valeriu can predict where Marcel will go in the next move. Valeriu can use this
information to make his move. They start and end the move at the same time.

It is guaranteed that any pair of buildings is connected by some path and there is at most one road between any pair of
buildings.

Assuming both players play optimally, answer if Valeriu has a strategy to indefinitely escape Marcel.

### thoughts

1. 一共n个节点，n条边，且是连通的
2. 所以，除去某条边后，剩下的形成了一棵树
3. 如果 Valeriu能够进入那个环，他肯定可以escape
4. 先找到这个环，然后判断两人到达环的时间；在相等时，判断是否进入同一个点