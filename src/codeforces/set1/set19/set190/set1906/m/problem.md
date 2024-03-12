You are given a regular 𝑁
-sided polygon. Label one arbitrary side as side 1
, then label the next sides in clockwise order as side 2
, 3
, …
, 𝑁
. There are 𝐴𝑖
special points on side 𝑖
. These points are positioned such that side 𝑖
is divided into 𝐴𝑖+1
segments with equal length.

For instance, suppose that you have a regular 4
-sided polygon, i.e., a square. The following illustration shows how the special points are located within each side
when 𝐴=[3,1,4,6]
. The uppermost side is labelled as side 1
.

You want to create as many non-degenerate triangles as possible while satisfying the following requirements. Each
triangle consists of 3
distinct special points (not necessarily from different sides) as its corners. Each special point can only become the
corner of at most 1
triangle. All triangles must not intersect with each other.

Determine the maximum number of non-degenerate triangles that you can create.

A triangle is non-degenerate if it has a positive area.

### ideas

1. 理想情况下 = sum() / 3
2. 但是如果有一条边太多的话，有可能无法全部使用完