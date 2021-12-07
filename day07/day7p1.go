package main 

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

const MIN = math.MinInt64
const MAX = math.MaxInt64

func main() {
	var input string 
	var h_positions []int

	fmt.Scan(&input)

	for _,s := range strings.Split(input,",") {
		n,_ := strconv.Atoi(s)
		h_positions = append(h_positions,n)
	}	

	//detect leftmost crab
	var minPosition = min(h_positions)
	//detect rightmost crab
	var maxPosition = max(h_positions)

	fmt.Println(maxPosition)
	fmt.Println(minPosition)

	var min_fuel = MAX
	var min_fuel_position = 0

	for p := minPosition; p <= maxPosition; p++ {
		fuel := 0
		for j := 0; j < len(h_positions); j++ {
			if h_positions[j] >= p {
				fuel = fuel + cumsum(h_positions[j]-p)
			} else {
				fuel = fuel + cumsum(p-h_positions[j])
			}
		}

		if fuel <= min_fuel {
			min_fuel = fuel
			min_fuel_position = p
		}

	}

	fmt.Println(min_fuel)
	fmt.Println(min_fuel_position)	
}

func min(positions []int) int {
	var min int = MAX

	for _,p := range positions {
		if p < min {
			min = p
		}
	}

	return min 
}

func max(positions []int) int {
	var max int = MIN

	for _,p := range positions {
		if p > max {
			max = p
		}
	}
	
	return max
}

func cumsum(x int) int{
	var sum int
	for i := 1; i <= x; i++ {
		sum += i
	} 
	return sum
}