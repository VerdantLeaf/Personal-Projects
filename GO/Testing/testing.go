package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("LOL")

	var starttime time.Time = time.Now()

	fmt.Println(starttime)

	var array [100]float64

	for i := 0; i < 100; i++ {
		array[i] = math.Cos(math.Pi / float64(i))
		fmt.Printf("%0.3f\n", array[i])
	}

}
