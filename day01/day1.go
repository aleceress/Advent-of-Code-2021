package main

import "fmt"

func main() {
	var n1, n2 int 
	var inc_count int

	fmt.Scan(&n1)
	for {
		_, err := fmt.Scan(&n2)
		if err != nil {
			break
		}
		if n2 >= n1 {
			inc_count++
		}
		n1 = n2
	}

	fmt.Println(inc_count)
}