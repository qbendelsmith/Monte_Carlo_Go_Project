// monteCarlo.go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// monteCarloPi follows the Monte Carlo method of approximating the value
// of pi. Using a unit square of 1, if you randomly throw darts at the
// square, the number of darts that would fall in a quarter-circle with
// the radius 1 will divided by the total number of darts times 4 will
// approximate the area of the circle, which is conveniently pi*r^2
// so when r = 1, pi should be the area of the circle
func monteCarloPi(darts int) float64 {
	// counter for whether the randomly chose coordinate is within the
	// quarter circle
	hit := 0
	// loop over the number of darts thrown
	for i := 0; i < darts; i++ {
		//randomly choose an x coordinate
		x_coord := rand.Float64()
		//randomly choose a y coordinate
		y_coord := rand.Float64()
		// distance to a coordinate is the square root of x^2 + y^2
		distance := math.Sqrt((x_coord * x_coord) + (y_coord * y_coord))
		// since our radius is 1, if the distance is less than or equal to 1
		if distance <= 1 {
			//increment hit
			hit++
		}
	}
	// multiply 4 (number of quadrants in a circle) and hits/darts
	return 4 * (float64(hit) / float64(darts))
}

func monteCarloPiThreaded(darts int, ch chan int) {
	// counter for whether the randomly chose coordinate is within the
	// quarter circle
	hit := 0
	// loop over the number of darts thrown
	for i := 0; i < darts; i++ {
		//randomly choose an x coordinate
		x_coord := rand.Float64()
		//randomly choose a y coordinate
		y_coord := rand.Float64()
		// distance to a coordinate is the square root of x^2 + y^2
		distance := math.Sqrt((x_coord * x_coord) + (y_coord * y_coord))
		// since our radius is 1, if the distance is less than or equal to 1
		if distance <= 1 {
			//increment hit
			hit++
		}
	}
	// add total hits to channel
	ch <- hit
}

func helper(threads int) float64 {
	x := 0
	const darts = 10000000000
	dartsPerThread := darts / threads
	for i := 0; i < threads; i++ {
		// tried making the channel outside of the loop and it was less efficient
		ch := make(chan int)
		//run goroutine for each subproblem
		go monteCarloPiThreaded(dartsPerThread, ch)
		// increment x for each run of goroutine
		x += <-ch
	}
	//take the total of x and divide by total darts * 4 for pi estimate
	return 4 * float64(x) / float64(darts)
}

func main() {
	// test monteCarloPi with increasingly higher number of darts.
	// used tutorialspoint to see how to track time
	// https://www.tutorialspoint.com/how-to-measure-the-execution-time-in-golang
	start := time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(1)))
	elapsed := time.Since(start)
	fmt.Printf("monteCarloPi at 1 thread took %s\n", elapsed)
	start = time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(2)))
	elapsed = time.Since(start)
	fmt.Printf("monteCarloPi at 2 threads took %s\n", elapsed)
	start = time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(4)))
	elapsed = time.Since(start)
	fmt.Printf("monteCarloPi at 4 threads took %s\n", elapsed)
	start = time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(8)))
	elapsed = time.Since(start)
	fmt.Printf("monteCarloPi at 8 threads took %s\n", elapsed)
	start = time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(16)))
	elapsed = time.Since(start)
	fmt.Printf("monteCarloPi at 16 threads took %s\n", elapsed)
	start = time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(32)))
	elapsed = time.Since(start)
	fmt.Printf("monteCarloPi at 32 threads took %s\n", elapsed)
	start = time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(64)))
	elapsed = time.Since(start)
	fmt.Printf("monteCarloPi at 64 threads took %s\n", elapsed)
	start = time.Now()
	fmt.Printf("The delta from pi is %f\n", math.Abs(math.Pi-helper(128)))
	elapsed = time.Since(start)
	fmt.Printf("monteCarloPi at 128 threads took %s\n", elapsed)
}
