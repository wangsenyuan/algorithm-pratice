There are 𝑛
 kids, numbered from 1
 to 𝑛
, dancing in a circle around the Christmas tree. Let's enumerate them in a clockwise direction as 𝑝1
, 𝑝2
, ..., 𝑝𝑛
 (all these numbers are from 1
 to 𝑛
 and are distinct, so 𝑝
 is a permutation). Let the next kid for a kid 𝑝𝑖
 be kid 𝑝𝑖+1
 if 𝑖<𝑛
 and 𝑝1
 otherwise. After the dance, each kid remembered two kids: the next kid (let's call him 𝑥
) and the next kid for 𝑥
. Each kid told you which kids he/she remembered: the kid 𝑖
 remembered kids 𝑎𝑖,1
 and 𝑎𝑖,2
. However, the order of 𝑎𝑖,1
 and 𝑎𝑖,2
 can differ from their order in the circle.

Example: 5 kids in a circle, 𝑝=[3,2,4,1,5]
 (or any cyclic shift). The information kids remembered is: 𝑎1,1=3
, 𝑎1,2=5
; 𝑎2,1=1
, 𝑎2,2=4
; 𝑎3,1=2
, 𝑎3,2=4
; 𝑎4,1=1
, 𝑎4,2=5
; 𝑎5,1=2
, 𝑎5,2=3
.
You have to restore the order of the kids in the circle using this information. If there are several answers, you may print any. It is guaranteed that at least one solution exists.