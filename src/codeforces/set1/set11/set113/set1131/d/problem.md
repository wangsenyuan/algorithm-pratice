Mr. Apple, a gourmet, works as editor-in-chief of a gastronomic periodical. He travels around the world, tasting new delights of famous chefs from the most fashionable restaurants. Mr. Apple has his own signature method of review  — in each restaurant Mr. Apple orders two sets of dishes on two different days. All the dishes are different, because Mr. Apple doesn't like to eat the same food. For each pair of dishes from different days he remembers exactly which was better, or that they were of the same quality. After this the gourmet evaluates each dish with a positive integer.

Once, during a revision of a restaurant of Celtic medieval cuisine named «Poisson», that serves chestnut soup with fir, warm soda bread, spicy lemon pie and other folk food, Mr. Apple was very pleasantly surprised the gourmet with its variety of menu, and hence ordered too much. Now he's confused about evaluating dishes.

The gourmet tasted a set of 𝑛
 dishes on the first day and a set of 𝑚
 dishes on the second day. He made a table 𝑎
 of size 𝑛×𝑚
, in which he described his impressions. If, according to the expert, dish 𝑖
 from the first set was better than dish 𝑗
 from the second set, then 𝑎𝑖𝑗
 is equal to ">", in the opposite case 𝑎𝑖𝑗
 is equal to "<". Dishes also may be equally good, in this case 𝑎𝑖𝑗
 is "=".

Now Mr. Apple wants you to help him to evaluate every dish. Since Mr. Apple is very strict, he will evaluate the dishes so that the maximal number used is as small as possible. But Mr. Apple also is very fair, so he never evaluates the dishes so that it goes against his feelings. In other words, if 𝑎𝑖𝑗
 is "<", then the number assigned to dish 𝑖
 from the first set should be less than the number of dish 𝑗
 from the second set, if 𝑎𝑖𝑗
 is ">", then it should be greater, and finally if 𝑎𝑖𝑗
 is "=", then the numbers should be the same.

Help Mr. Apple to evaluate each dish from both sets so that it is consistent with his feelings, or determine that this is impossible.

### ideas
1. 一共n+m个节点
2. i和j之间有一条边（表示 val[i] < val[j]) 或者反过来
3. 如何表示 val[i] = val[j] 呢？加两条边？
4. 然后看 ssc，如果一个ssc中的，它们是相等的，那么就ok，但是如果不是的，那么就没有答案