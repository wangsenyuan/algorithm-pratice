Long is a huge fan of CFC (Codeforces Fried Chicken). But the price of CFC is increasing, so he decides to breed the chicken on his own farm.

His farm is presented by a rectangle grid with 𝑟
 rows and 𝑐
 columns. Some of these cells contain rice, others are empty. 𝑘
 chickens are living on his farm. The number of chickens is not greater than the number of cells with rice on the farm.

Long wants to give his chicken playgrounds by assigning these farm cells to his chickens. He would like to satisfy the following requirements:

Each cell of the farm is assigned to exactly one chicken.
Each chicken is assigned at least one cell.
The set of cells assigned to every chicken forms a connected area. More precisely, if two cells (𝑥,𝑦)
 and (𝑢,𝑣)
 are assigned to the same chicken, this chicken is able to walk from (𝑥,𝑦)
 to (𝑢,𝑣)
 by passing only its cells and moving from each cell to another cell sharing a side.
Long also wants to prevent his chickens from fighting for food. Hence he wants the difference between the maximum and the minimum number of cells with rice assigned to a chicken to be as small as possible. Please help him.



### ideas
1. 分配给同一个小鸡的区域，必须是连接在一起的
2. 小鸡的个数是x，R的个数是y (y <= x>)
3. 那么如果 x % y = 0, => 0 else 1
4. 现在剩下的就是怎么分配的问题
5. 如果在一维就很简单处理
6. 如果是2维，假设每个小鸡要分配2个，怎么连续的分配
7. 如果在上一行的后半段给小鸡分配了，下一行的前半段分配后，这两部分要有重叠的区域才行
8. 所以，要把最近的放在一起，用 r + c 作为距离，进行分配就可以了