Bees Alice and Alesya gave beekeeper Polina famous card game "Set" as a Christmas present. The deck consists of cards that vary in four features across three options for each kind of feature: number of shapes, shape, shading, and color. In this game, some combinations of three cards are said to make up a set. For every feature — color, number, shape, and shading — the three cards must display that feature as either all the same, or pairwise different. The picture below shows how sets look.



Polina came up with a new game called "Hyperset". In her game, there are 𝑛
 cards with 𝑘
 features, each feature has three possible values: "S", "E", or "T". The original "Set" game can be viewed as "Hyperset" with 𝑘=4
.

Similarly to the original game, three cards form a set, if all features are the same for all cards or are pairwise different. The goal of the game is to compute the number of ways to choose three cards that form a set.

Unfortunately, winter holidays have come to an end, and it's time for Polina to go to school. Help Polina find the number of sets among the cards lying on the table.


### ideas
1. 阅读理解
2. n个card，有k个feature，每个feature有3个值, 
3. 然后选择3张cards，满足条件，对于feature，要么值都相同，要么都不相同，组成一个set
4. 有多少组方式
5. 选中了(i, j)后，第三种k的选择是确定的