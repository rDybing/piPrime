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

### Update:

My first attempt would not work entirely as expected. Using go-routines actually took a lot longer to complete (~18sec) than doing it serially (~5sec) when testing with a reduced count limit of 1.6 million.

Asked on reddit/r/golang - and got a few helpful replies. 

One, siritinga, even included some modified source code, so that is what this version is running. And much faster - using all four CPU cores as intended!

Doing it serially now takes roughly 2m19s finding all the primes up to 16 million - using Go routines in parallel now takes about a quarter of the time at roughly 58s.

Full readout:

here we go. Ol'skool serial prime-finder

*  1000000 || time: 2.905756589s || primes: 78498
*  2000000 || time: 7.553869118s || primes: 148933
*  3000000 || time: 13.256930484s || primes: 216816
*  4000000 || time: 19.799640909s || primes: 283146
*  5000000 || time: 27.101634158s || primes: 348513
*  6000000 || time: 35.01149957s || primes: 412849
*  7000000 || time: 43.475854512s || primes: 476648
*  8000000 || time: 52.389431267s || primes: 539777
*  9000000 || time: 1m1.783371376s || primes: 602489
* 10000000 || time: 1m11.600970307s || primes: 664579
* 11000000 || time: 1m21.85521713s || primes: 726517
* 12000000 || time: 1m32.498435842s || primes: 788060
* 13000000 || time: 1m43.524305601s || primes: 849252
* 14000000 || time: 1m54.942358213s || primes: 910077
* 15000000 || time: 2m6.735969736s || primes: 970704
* 16000000 || time: 2m18.833899131s || primes: 1031130

and now the concurrent way!

*  1000000 || time: 1.748451671s || primes: 78498
*  2000000 || time: 4.009315993s || primes: 148932
*  3000000 || time: 6.62305898s || primes: 216815
*  4000000 || time: 9.495219865s || primes: 283145
*  5000000 || time: 12.797978168s || primes: 348512
*  6000000 || time: 16.14136534s || primes: 412848
*  7000000 || time: 19.637838421s || primes: 476646
*  8000000 || time: 23.600505743s || primes: 539775
*  9000000 || time: 27.201350715s || primes: 602489
* 10000000 || time: 31.132531027s || primes: 664576
* 11000000 || time: 35.168005782s || primes: 726514
* 12000000 || time: 39.293489409s || primes: 788059
* 13000000 || time: 43.958292469s || primes: 849249
* 14000000 || time: 48.309189131s || primes: 910075
* 15000000 || time: 52.894934592s || primes: 970703
* 16000000 || time: 57.989160858s || primes: 1031130

As can be seen, the ARM CPU on the Pi3 is quite faster than the ARM CPU on the Parallella SoC board. Nearly twice as fast running in single thread.

Going parallel however, the 16 cores of the Parallella SoC board outshines the 4 cores of the rPi3, completing the task over three times as fast.