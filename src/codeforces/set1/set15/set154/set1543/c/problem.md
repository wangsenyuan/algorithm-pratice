After defeating a Blacklist Rival, you get a chance to draw 1
 reward slip out of 𝑥
 hidden valid slips. Initially, 𝑥=3
 and these hidden valid slips are Cash Slip, Impound Strike Release Marker and Pink Slip of Rival's Car. Initially, the probability of drawing these in a random guess are 𝑐
, 𝑚
, and 𝑝
, respectively. There is also a volatility factor 𝑣
. You can play any number of Rival Races as long as you don't draw a Pink Slip. Assume that you win each race and get a chance to draw a reward slip. In each draw, you draw one of the 𝑥
 valid items with their respective probabilities. Suppose you draw a particular item and its probability of drawing before the draw was 𝑎
. Then,

If the item was a Pink Slip, the quest is over, and you will not play any more races.
Otherwise,
If 𝑎≤𝑣
, the probability of the item drawn becomes 0
 and the item is no longer a valid item for all the further draws, reducing 𝑥
 by 1
. Moreover, the reduced probability 𝑎
 is distributed equally among the other remaining valid items.
If 𝑎>𝑣
, the probability of the item drawn reduces by 𝑣
 and the reduced probability is distributed equally among the other valid items.
For example,

If (𝑐,𝑚,𝑝)=(0.2,0.1,0.7)
 and 𝑣=0.1
, after drawing Cash, the new probabilities will be (0.1,0.15,0.75)
.
If (𝑐,𝑚,𝑝)=(0.1,0.2,0.7)
 and 𝑣=0.2
, after drawing Cash, the new probabilities will be (𝐼𝑛𝑣𝑎𝑙𝑖𝑑,0.25,0.75)
.
If (𝑐,𝑚,𝑝)=(0.2,𝐼𝑛𝑣𝑎𝑙𝑖𝑑,0.8)
 and 𝑣=0.1
, after drawing Cash, the new probabilities will be (0.1,𝐼𝑛𝑣𝑎𝑙𝑖𝑑,0.9)
.
If (𝑐,𝑚,𝑝)=(0.1,𝐼𝑛𝑣𝑎𝑙𝑖𝑑,0.9)
 and 𝑣=0.2
, after drawing Cash, the new probabilities will be (𝐼𝑛𝑣𝑎𝑙𝑖𝑑,𝐼𝑛𝑣𝑎𝑙𝑖𝑑,1.0)
.
You need the cars of Rivals. So, you need to find the expected number of races that you must play in order to draw a pink slip.


### ideas
1. 