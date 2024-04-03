Nephren is playing a game with little leprechauns.

She gives them an infinite array of strings, f0... ∞.

f0 is "What are you doing at the end of the world? Are you busy? Will you save us?".

She wants to let more people know about it, so she defines fi ="What are you doing while sending "fi - 1"? Are you busy?
Will you send "fi - 1"?" for all i ≥ 1.

For example, f1 is

"What are you doing while sending "What are you doing at the end of the world? Are you busy? Will you save us?"? Are you
busy? Will you send "What are you doing at the end of the world? Are you busy? Will you save us?"?". Note that the
quotes in the very beginning and in the very end are for clarity and are not a part of f1.

It can be seen that the characters in fi are letters, question marks, (possibly) quotation marks and spaces.

Nephren will ask the little leprechauns q times. Each time she will let them find the k-th character of fn. The
characters are indexed starting from 1. If fn consists of less than k characters, output '.' (without quotes).

Can you answer her queries?

### ideas

1. 