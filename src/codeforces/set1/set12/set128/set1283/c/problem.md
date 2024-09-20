There are n
 friends who want to give gifts for the New Year to each other. Each friend should give exactly one gift and receive exactly one gift. The friend cannot give the gift to himself.

For each friend the value fi
 is known: it is either fi=0
 if the i
-th friend doesn't know whom he wants to give the gift to or 1≤fi≤n
 if the i
-th friend wants to give the gift to the friend fi
.

You want to fill in the unknown values (fi=0
) in such a way that each friend gives exactly one gift and receives exactly one gift and there is no friend who gives the gift to himself. It is guaranteed that the initial information isn't contradictory.

If there are several answers, you can print any.

### ideas
1. 考虑一个环，每个节点，都满足give/receive刚好一个的条件
2. 除了环，没有其他的可以满足的形式
3. 所以，先把已经存在的环，找出来，然后找出一条线，一个点的那种，把那些一个点的，都连到某条线上即可
4. 最后把所有的线，变成环，即可