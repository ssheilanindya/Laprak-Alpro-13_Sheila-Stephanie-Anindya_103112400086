package main

import (
	"fmt"
	"math"
)

func main() {
	var desimal float64

	fmt.Scan(&desimal)

	bulat := math.Ceil(desimal)

	for {
		fmt.Printf("%.1f\n", desimal)
		desimal += 0.1

		if desimal > bulat {
			break
		}
	}
}
