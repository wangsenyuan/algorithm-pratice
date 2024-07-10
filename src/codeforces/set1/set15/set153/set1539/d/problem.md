Lena is the most economical girl in Moscow. So, when her dad asks her to buy some food for a trip to the country, she goes to the best store  — "PriceFixed". Here are some rules of that store:

The store has an infinite number of items of every product.
All products have the same price: 2
 rubles per item.
For every product 𝑖
 there is a discount for experienced buyers: if you buy 𝑏𝑖
 items of products (of any type, not necessarily type 𝑖
), then for all future purchases of the 𝑖
-th product there is a 50%
 discount (so you can buy an item of the 𝑖
-th product for 1
 ruble!).
Lena needs to buy 𝑛
 products: she must purchase at least 𝑎𝑖
 items of the 𝑖
-th product. Help Lena to calculate the minimum amount of money she needs to spend if she optimally chooses the order of purchasing. Note that if she wants, she can buy more items of some product than needed.


### ideas
1. lena至少要买n件，且每个至少买a[i]件
2. 假设最小的b[i], 那么在买到b[i]件商品前，lena是不会有discount的
3. 还有一个点，就是假设现在有x件，但是要购买下一件item，还不够b[i]， 如果直接去购买它，需要花费 2 * a[i]
4. 但是如果通过购买当前有折扣的部分, b[i] - x + a[i] <= 2 * a[i]
5. x >= b[i] - a[i] 也就是说，最好的方案就是就是购买当前折扣为1的部分，知道能够购买后面的
6. 此时的花费 = b[i] - x + a[i] 这种情况下，显然 b[i]越小越好
7. 为了增长更快 a[i]越大越好.