## piPrime

A small app to time how fast a pi3 using Go can get all primes up to 16 million.

Made to compare with the benchmark for the Parallella SoC parallel computer as presented in this talk:

youtube.com/watch?v=BHZCCUEzK0s

The video showcase a 150USD pi-sized computer using a 2 cores ARM CPU and the Epiphany Chip with an additional 16 RISC cores, all running at 1Ghz.

In the video, they're running a C program to find primes up to 16 million, in serial and parallel. The results:

Serial  : 237.1 seconds (I'm assuming it is run on the ARM)
Parallel:  18.6 seconds (full tilt on all 16 Epiphany cores)

The Serial test is done with a simple 3 to 16 million step by 2 iteration. 

The Parallel test was done by dividing the problem into 16 chunks of 500k numbers each, for a total of 8 million - but being only odd numbers, you still get to 16 million...

Every 100k it reports back where it is in the program and time taken so far by dumping info to terminal.

Using much the same algorithms, let's find out how Go on Pi3 compares, though I've not bothered to chunk it up as it were. I've set max threads to 4, the number of cores on the Pi3. The divide and conquer method used in the C example I believe was done due to the less than stellar methods needed to achieve concurrency in that language. Don't get me wrong, I love C - used to its' strengths, it's great.

Oh, and also I did not bother with the entire start from 3 and step by 2 process. This is quick and dirty, and there is not much to save by complicating things (the output to terminal chiefly). Doing it serially, only ~2 seconds less to be approximately exact.


