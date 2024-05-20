This is an interactive problem.

You are given a grid with 𝑛
 rows and 𝑚
 columns. The coordinates (𝑥,𝑦)
 represent the cell on the grid, where 𝑥
 (1≤𝑥≤𝑛
) is the row number counting from the top and 𝑦
 (1≤𝑦≤𝑚
) is the column number counting from the left. It is guaranteed that there are exactly 2
 mines in the grid at distinct cells, denoted as (𝑥1,𝑦1)
 and (𝑥2,𝑦2)
. You are allowed to make no more than 4
 queries to the interactor, and after these queries, you need to provide the location of one of the mines.

In each query, you can choose any grid cell (𝑥,𝑦)
, and in return, you will receive the minimum Manhattan distance from both the mines to the chosen cell, i.e., you will receive the value min(|𝑥−𝑥1|+|𝑦−𝑦1|,|𝑥−𝑥2|+|𝑦−𝑦2|)
.

Your task is to determine the location of one of the mines after making the queries.


### ideas
1. 选择 (1, 1) 会返回什么，(x1 + y1) - 2 or (x2 + y2) -2
2. 选在 (1, m) 会返回, (x1 - 1 + m - y1) or (x2 - 1 + m - y2)
3. 选择 (n, 1) 会返回 (n - x1 + y1 - 1) or (n - x2 + y2 - 1)
4. 选择 (n, m) 会返回 (n - x1 + m - y1) or (n - x2 + m - y2)
5. 如果第一次返回的是1，第4次返回的肯定是2，然后带入2/3进行检查