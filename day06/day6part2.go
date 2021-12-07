package main 

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var input string 
	var timers []int

	fmt.Scan(&input)

	for _,s := range strings.Split(input,",") {
		n,_ := strconv.Atoi(s)
		timers = append(timers,n)
	}

	var days [9]uint64

	for _,c := range timers {
		days[c]++
	}

	var days_number = 256
	var fish_number uint64
	
	for i := 0; i < days_number; i++ {
		fish_number += days[0]
		days = shift(days)
	}

	fmt.Println(fish_number)
}

func shift(days [9]uint64) [9]uint64 {
	var start = days[0]

	for i := 1; i < len(days); i++ {
		days[i-1] = days[i] 

		if i == 7 {
			days[i-1] += start
		}
	}
	days[8] = start

	return days
}

