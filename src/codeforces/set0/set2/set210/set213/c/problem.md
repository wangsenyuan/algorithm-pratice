Furik and Rubik take part in a relay race. The race will be set up on a large square with the side of n meters. The
given square is split into n × n cells (represented as unit squares), each cell has some number.

At the beginning of the race Furik stands in a cell with coordinates (1, 1), and Rubik stands in a cell with
coordinates (n, n). Right after the start Furik runs towards Rubik, besides, if Furik stands at a cell with
coordinates (i, j), then he can move to cell (i + 1, j) or (i, j + 1). After Furik reaches Rubik, Rubik starts running
from cell with coordinates (n, n) to cell with coordinates (1, 1). If Rubik stands in cell (i, j), then he can move to
cell (i - 1, j) or (i, j - 1). Neither Furik, nor Rubik are allowed to go beyond the boundaries of the field; if a
player goes beyond the boundaries, he will be disqualified.

To win the race, Furik and Rubik must earn as many points as possible. The number of points is the sum of numbers from
the cells Furik and Rubik visited. Each cell counts only once in the sum.

Print the maximum number of points Furik and Rubik can earn on the relay race.