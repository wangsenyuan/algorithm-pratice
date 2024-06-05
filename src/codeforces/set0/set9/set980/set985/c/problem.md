You have m = n·k wooden staves. The i-th stave has length ai. You have to assemble n barrels consisting of k staves each, you can use any k staves to construct a barrel. Each stave must belong to exactly one barrel.

Let volume vj of barrel j be equal to the length of the minimal stave in it.


You want to assemble exactly n barrels with the maximal total sum of volumes. But you have to make them equal enough, so a difference between volumes of any pair of the resulting barrels must not exceed l, i.e. |vx - vy| ≤ l for any 1 ≤ x ≤ n and 1 ≤ y ≤ n.

Print maximal total sum of volumes of equal enough barrels or 0 if it's impossible to satisfy the condition above.

