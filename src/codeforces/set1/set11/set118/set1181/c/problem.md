Innokenty works at a flea market and sells some random stuff rare items. Recently he found an old rectangular blanket. It turned out that the blanket is split in 𝑛⋅𝑚
 colored pieces that form a rectangle with 𝑛
 rows and 𝑚
 columns.

The colored pieces attracted Innokenty's attention so he immediately came up with the following business plan. If he cuts out a subrectangle consisting of three colored stripes, he can sell it as a flag of some country. Innokenty decided that a subrectangle is similar enough to a flag of some country if it consists of three stripes of equal heights placed one above another, where each stripe consists of cells of equal color. Of course, the color of the top stripe must be different from the color of the middle stripe; and the color of the middle stripe must be different from the color of the bottom stripe.

Innokenty has not yet decided what part he will cut out, but he is sure that the flag's boundaries should go along grid lines. Also, Innokenty won't rotate the blanket. Please help Innokenty and count the number of different subrectangles Innokenty can cut out and sell as a flag. Two subrectangles located in different places but forming the same flag are still considered different.

### ideas
1. 假设从(i, j)开始，那么下一个开始的位置 (k, j) 肯定是和a[i][j]不同颜色的第一个位置
2. 再下一个数就确定了。这个唯一的问题是需要找到宽度