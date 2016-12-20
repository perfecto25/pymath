# Math experiments

## the Freeman Dyson "Doubler"
I was reading the NY Times article on famous physicist Freeman Dyson and this paragraph caught my eye:

"A group of scientists will be sitting around the cafeteria, and one will idly wonder if there is an integer where, if you take its last digit and move it to the front, turning, say, 112 to 211, it’s possible to exactly double the value. Dyson will immediately say, “Oh, that’s not difficult,” allow two short beats to pass and then add, “but of course the smallest such number is 18 digits long.”

http://www.nytimes.com/2009/03/29/magazine/29Dyson-t.html?pagewanted=all&_r=0

I wanted to find out whether Dyson was right about an 18-digit number. This is the experiment, in py and golang (golang written for faster execution, the py script was taking days to run). 

I ran this on my Rasberry Pi since its always up and I can leave up running for weeks.

The key function is MoveLastToFirst, as the paragraph above states, the experiment is to see if a number that has its last integer moved to 1st position can be a double of the original.

### pseudocode
```bash
x = 2103
doubler = 4206
lastToFirst(x) = 3210

if lastToFirst == double
print Found number: x
```

 I wasnt sure how to make an optimal function to do this, except to convert an int into an string array and then move the last array elements to front. I cant simply reverse an int, for example int 45678 'doubler' is 84567, not 87654. Once the lastToFirst is done, I convert the string array back into a single integer, and then have another function "countIntegers" count how many numbers are in the integer (for example, 81334980, it returns a 8) - this is to see if Dyson was right about the number being 18 integers long.

The golang code is the better one to run since its compiled and runs much faster than python.

If anyone knows a more efficient method, submit a pull request