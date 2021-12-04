package main 

import "fmt"

func main() {
	var horizontal_position int 
	var depth int 
	var aim int 
	
	for {
		var s string 
		var n int 

		fmt.Scan(&s)
		_,err := fmt.Scan(&n)
		if err != nil {
			break
		}

		switch(s) {
			case "forward": 
				horizontal_position += n
			case "down": 
				depth += n
			case "up": 
				depth -= n
		}
		
	}

	fmt.Println(depth*horizontal_position)
}