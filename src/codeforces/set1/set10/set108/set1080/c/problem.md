Recently, Masha was presented with a chessboard with a height of 𝑛
and a width of 𝑚
.

The rows on the chessboard are numbered from 1
to 𝑛
from bottom to top. The columns are numbered from 1
to 𝑚
from left to right. Therefore, each cell can be specified with the coordinates (𝑥,𝑦)
, where 𝑥
is the column number, and 𝑦
is the row number (do not mix up).

Let us call a rectangle with coordinates (𝑎,𝑏,𝑐,𝑑)
a rectangle lower left point of which has coordinates (𝑎,𝑏)
, and the upper right one — (𝑐,𝑑)
.

The chessboard is painted black and white as follows:

An example of a chessboard.
Masha was very happy with the gift and, therefore, invited her friends Maxim and Denis to show off. The guys decided to
make her a treat — they bought her a can of white and a can of black paint, so that if the old board deteriorates, it
can be repainted. When they came to Masha, something unpleasant happened: first, Maxim went over the threshold and
spilled white paint on the rectangle (𝑥1,𝑦1,𝑥2,𝑦2)
. Then after him Denis spilled black paint on the rectangle (𝑥3,𝑦3,𝑥4,𝑦4)
.

To spill paint of color 𝑐𝑜𝑙𝑜𝑟
onto a certain rectangle means that all the cells that belong to the given rectangle become 𝑐𝑜𝑙𝑜𝑟
. The cell dyeing is superimposed on each other (if at first some cell is spilled with white paint and then with black
one, then its color will be black).

Masha was shocked! She drove away from the guests and decided to find out how spoiled the gift was. For this, she needs
to know the number of cells of white and black color. Help her find these numbers!

### ideas

1. take care the latter black will overwrite the previous white rectangle
2. 好麻烦，要考虑的情况太多了