There are 𝑛
people taking part in a show about VOCALOID. They will sit in the row of seats, numbered 1
to 𝑚
from left to right.

The 𝑛
people come and sit in order. Each person occupies a seat in one of three ways:

Sit in the seat next to the left of the leftmost person who is already sitting, or if seat 1
is taken, then leave the show. If there is no one currently sitting, sit in seat 𝑚
.
Sit in the seat next to the right of the rightmost person who is already sitting, or if seat 𝑚
is taken, then leave the show. If there is no one currently sitting, sit in seat 1
.
Sit in the seat numbered 𝑥𝑖
. If this seat is taken, then leave the show.
Now you want to know what is the maximum number of people that can take a seat, if you can let people into the show in
any order?

The second line of each test case contains 𝑛
integers 𝑥1,𝑥2,…,𝑥𝑛
(−2≤𝑥𝑖≤𝑚
, 𝑥𝑖≠0
), the 𝑖
-th of which describes the way in which the 𝑖
-th person occupies a seat:

If 𝑥𝑖=−1
, then 𝑖
-th person takes the seat in the first way.
If 𝑥𝑖=−2
, then 𝑖
-th person takes the seat in the second way.
If 𝑥𝑖>0
, then the 𝑖
-th person takes a seat in the third way, i.e. he wants to sit in the seat with the number 𝑥𝑖
or leave the show if it is occupied..