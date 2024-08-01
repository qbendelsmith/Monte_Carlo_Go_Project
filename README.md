Project 5 - Monte Carlo Simulation for Pi Implementation
Quinn Bendelsmith

This code will implement a method in Go to run a Monte Carlo Simulation to estimate the value of pie.  What is a Monte Carlo Simulation, you may ask? Essentially randomly select 2 values between 0 and 1 and treat them as coordinates on a plot. 0-1 allows for us to visually imagine a unit square with coordinates (0,0),(1,0),(1,1),(0,1). With this square you can also imagine a quarter-circle with a radius of 1. Why is this necessary? Because the area of a circle is pi * r^2. If r = 1, then the area equals pi. Now if you run this randomization an infinite amount of times and tally the amount of times the "dart", which is essentially the randomization of picking a point, hits (is in circle) compared to the amount of total darts and multiply by 4, you will have an approximation of the value of pi.

This is a very inefficient method of calculating pi, but it is a fun way to approach it.

How to run this code: Assuming you already have Go downloaded, go to command prompt or any equivalent, change the directory of where the file is. In my case it is C:\Users\qbend\project5bendelsmith. Then simply type go run monteCarlo.go and it will run the simulation.

Implementation 1: We were tasked with simply implementing a version of the monte carlo simulation using Go. So I wrote the file monteCarlo.go which contains a monteCarloPi function and a main function.

monteCarloPi randomizes two float64 numbers between 0 and 1 and treats it as a point in a unit square. If the distance is less than or equal to 1, then it is within a circle of radius 1 within the square. Now to calculate pi, you multiply it by 4 to get the total area of the circle, which is pi. I ran the simulation at increasing number of times 6 times to see how much closer to actual pi it gets with each simulation.

Iteration 2: We were tasked with finding the delta between real pi and the approximated pi from the Monte Carlo method. We were also tasked with timing how long the method takes with an increasing amount of darts thrown ranging from 1,000 to 10,000,000,000. The results from the code will be written out in a separate document. As is evident from the results, the more darts thrown the closer and closer it slowly converges to actual pi, but with the addition of more darts comes a much larger runtime.

Implementation 3: This time we will run the same experiment but with different levels of concurrency. I have implemented a new helper method and a new monteCarloPi method to work with a channel.

monteCarloPiThreaded runs the same monteCarloPi but instead of returning the pi estimate, it sends the amount of hits to a channel.

helper takes the number of threads and equally splits the darts per each thread. It runs goroutines and adds to a channel within a for loop and assigns the channel values to x which increments with the hits from each run of monteCarloThreaded. It then divides the hits by darts and multiplies it by 4.

The results from this were underwhelming in terms of shortening runtime. I think partially because I am running too many other things on my laptop and because the program probably is not optimized very well. Due to being short on time and not able to work any further on this project, I am calling it good and willing to take a deduction if the threading was done incorrectly because I have been unable to fix this issue. 
