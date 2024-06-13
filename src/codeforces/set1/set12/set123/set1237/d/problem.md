Your favorite music streaming platform has formed a perfectly balanced playlist exclusively for you. The playlist consists of 𝑛
 tracks numbered from 1
 to 𝑛
. The playlist is automatic and cyclic: whenever track 𝑖
 finishes playing, track 𝑖+1
 starts playing automatically; after track 𝑛
 goes track 1
.

For each track 𝑖
, you have estimated its coolness 𝑎𝑖
. The higher 𝑎𝑖
 is, the cooler track 𝑖
 is.

Every morning, you choose a track. The playlist then starts playing from this track in its usual cyclic fashion. At any moment, you remember the maximum coolness 𝑥
 of already played tracks. Once you hear that a track with coolness strictly less than 𝑥2
 (no rounding) starts playing, you turn off the music immediately to keep yourself in a good mood.

For each track 𝑖
, find out how many tracks you will listen to before turning off the music if you start your morning with track 𝑖
, or determine that you will never turn the music off. Note that if you listen to the same track several times, every time must be counted.