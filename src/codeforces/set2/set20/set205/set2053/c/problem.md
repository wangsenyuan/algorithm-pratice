There are 𝑛
stars in the sky, arranged in a row. Iris has a telescope, which she uses to look at the stars.

Initially, Iris observes stars in the segment [1,𝑛]
, and she has a lucky value of 0
. Iris wants to look for the star in the middle position for each segment [𝑙,𝑟]
that she observes. So the following recursive procedure is used:

First, she will calculate 𝑚=⌊𝑙+𝑟2⌋
.
If the length of the segment (i.e. 𝑟−𝑙+1
) is even, Iris will divide it into two equally long segments [𝑙,𝑚]
and [𝑚+1,𝑟]
for further observation.
Otherwise, Iris will aim the telescope at star 𝑚
, and her lucky value will increase by 𝑚
; subsequently, if 𝑙≠𝑟
, Iris will continue to observe two segments [𝑙,𝑚−1]
and [𝑚+1,𝑟]
.
Iris is a bit lazy. She defines her laziness by an integer 𝑘
: as the observation progresses, she will not continue to observe any segment [𝑙,𝑟]
with a length strictly less than 𝑘
. In this case, please predict her final lucky value.