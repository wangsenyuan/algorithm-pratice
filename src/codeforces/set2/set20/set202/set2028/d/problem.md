Alice is playing cards with the Queen of Hearts, King of Hearts, and Jack of Hearts. There are 𝑛
 different types of cards in their card game. Alice currently has a card of type 1
 and needs a card of type 𝑛
 to escape Wonderland. The other players have one of each kind of card.

In this card game, Alice can trade cards with the three other players. Each player has different preferences for the 𝑛
 types of cards, which can be described by permutations∗
 𝑞
, 𝑘
, and 𝑗
 for the Queen, King, and Jack, respectively.

A player values card 𝑎
 more than card 𝑏
 if for their permutation 𝑝
, 𝑝𝑎>𝑝𝑏
. Then, this player is willing to trade card 𝑏
 to Alice in exchange for card 𝑎
. Alice's preferences are straightforward: she values card 𝑎
 more than card 𝑏
 if 𝑎>𝑏
, and she will also only trade according to these preferences.

Determine if Alice can trade up from card 1
 to card 𝑛
 subject to these preferences, and if it is possible, give a possible set of trades to do it.