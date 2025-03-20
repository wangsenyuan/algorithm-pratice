On Bertown's main street n trees are growing, the tree number i has the height of ai meters (1 ≤ i ≤ n). By the arrival of the President of Berland these trees were decided to be changed so that their heights formed a beautiful sequence. This means that the heights of trees on ends (the 1st one and the n-th one) should be equal to each other, the heights of the 2-nd and the (n - 1)-th tree must also be equal to each other, at that the height of the 2-nd tree should be larger than the height of the first tree by 1, and so on. In other words, the heights of the trees, standing at equal distance from the edge (of one end of the sequence) must be equal to each other, and with the increasing of the distance from the edge by 1 the tree height must also increase by 1. For example, the sequences "2 3 4 5 5 4 3 2" and "1 2 3 2 1" are beautiful, and '1 3 3 1" and "1 2 3 1" are not.

Changing the height of a tree is a very expensive operation, using advanced technologies invented by Berland scientists. In one operation you can choose any tree and change its height to any number, either increase or decrease. Note that even after the change the height should remain a positive integer, i. e, it can't be less than or equal to zero. Identify the smallest number of changes of the trees' height needed for the sequence of their heights to become beautiful.

### ideas
1. 肯定是 a[1], a[1] + 1, .... a[1] + i， 。。a[1] + 1, a[1]
2. 所以理论上a[1]确定后， 这个结果就确定了
3. 另外 a[i] - a[1] = i - 1 的话，它们可以和a[1]同时变化
4. 如果 a[i] != a[1], a[i] - a[1] > i - 1, 也就是比如 a[1] = 1, a[2] = 3
5. 那么如果a[1]增大的时候， a[2]其实有可能不用变化（或者变化的比较少）
6. 所以，可以根据a[i] - i 进行分组。然后排序。然后找中点