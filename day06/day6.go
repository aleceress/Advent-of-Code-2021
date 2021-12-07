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

	var days = 256 
	for i := 0; i < days; i++ {
		new := 0
		for j := 0; j < len(timers) - new; j++ {
			if timers[j] == 0 {
				new++
				timers = append(timers,8)
				timers[j] = 6
			} else {
				timers[j]--
			}
		}
	}

	fmt.Println(len(timers))
	
}