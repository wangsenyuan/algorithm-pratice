Innopolis University scientists continue to investigate the periodic table. There are n·m known elements and they form a periodic table: a rectangle with n rows and m columns. Each element can be described by its coordinates (r, c) (1 ≤ r ≤ n, 1 ≤ c ≤ m) in the table.

Recently scientists discovered that for every four different elements in this table that form a rectangle with sides parallel to the sides of the table, if they have samples of three of the four elements, they can produce a sample of the fourth element using nuclear fusion. So if we have elements in positions (r1, c1), (r1, c2), (r2, c1), where r1 ≠ r2 and c1 ≠ c2, then we can produce element (r2, c2).


Samples used in fusion are not wasted and can be used again in future fusions. Newly crafted elements also can be used in future fusions.

Innopolis University scientists already have samples of q elements. They want to obtain samples of all n·m elements. To achieve that, they will purchase some samples from other laboratories and then produce all remaining elements using an arbitrary number of nuclear fusions in some order. Help them to find the minimal number of elements they need to purchase.

### ideas
1. 又是个数学相关的。
2. 假设位置 (r, c) 能够被用到，那么有可能会触发更进一步的反应
3. 完全没想法呐～
4. 首先肯定不能把所有n*m个元素给表示出来
5. (r1, c), (r2, c) 如果 r1其他列上存在某个元素，那么r2对应的也可以有（所以这里是一种合并的关系）只有那些完全没有的在这种情况下是空的
6. 到r2为止，上面到底哪一行给某一列有值其实没有关系，重要的是那一列有值
7. 假设用一个数据结构知道到目前为止哪些列上有值，
8. 如果r2和它有重叠的情况，那么就可以直接把r2 union上去
9. 有点进展，但似乎还不足够