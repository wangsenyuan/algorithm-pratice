Today you are going to lead a group of elven archers to defend the castle that is attacked by an army of angry orcs. Three sides of the castle are protected by impassable mountains and the remaining side is occupied by a long wall that is split into n sections. At this moment there are exactly ai archers located at the i-th section of this wall. You know that archer who stands at section i can shoot orcs that attack section located at distance not exceeding r, that is all such sections j that |i - j| ≤ r. In particular, r = 0 means that archers are only capable of shooting at orcs who attack section i.

Denote as defense level of section i the total number of archers who can shoot at the orcs attacking this section. Reliability of the defense plan is the minimum value of defense level of individual wall section.

There is a little time left till the attack so you can't redistribute archers that are already located at the wall. However, there is a reserve of k archers that you can distribute among wall sections in arbitrary way. You would like to achieve maximum possible reliability of the defence plan.

### ideas
1. 考虑没有增援的情况下，一个位置的防守强度 = sum(i - r, i + r) 弓箭手的数量
2. 如果要增加，那么就应该在最远的i+r处增加