package main

import "fmt"

func main() {
	var measurements []int
	
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		measurements = append(measurements,n)	
	}

	var count_increasing int
	for i := 0; i< len(measurements) - 3; i++ {
		sum1 := measurements[i] + measurements[i+1] + measurements[i+2]
		sum2 := measurements[i+1] + measurements[i+2] + measurements[i+3] 
		if sum1 < sum2 {
			count_increasing++
		}
	}

	fmt.Println(count_increasing)
}