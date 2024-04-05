There are
X+Y+Z people, conveniently numbered
1 through
X+Y+Z. Person
i has
A
i
​
gold coins,
B
i
​
silver coins and
C
i
​
bronze coins.

Snuke is thinking of getting gold coins from
X of those people, silver coins from
Y of the people and bronze coins from
Z of the people. It is not possible to get two or more different colors of coins from a single person. On the other
hand, a person will give all of his/her coins of the color specified by Snuke.

Snuke would like to maximize the total number of coins of all colors he gets. Find the maximum possible number of coins.

### ideas

1. 先选择gold最大的x个人，然后添加新的选项，
2. 这个人的出现，有3种选择，就是使用他的gold/silver/bronze
3. 如果是gold，那么前面某个gold的需要被选出来，但选出来的那个要变成silver/bronze
4. 那么此时的，收益是 +gold[j] - gold[i] + silver[i]/bronze[i]