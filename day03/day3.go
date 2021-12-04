package main 

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	const length_input int = 12
	var occorrenze [2][length_input]int

	for {
		_, e := fmt.Scan(&number)

		if e != nil {
			break
		}

		for pos := 0; pos < length_input; pos++ {
			val := parseDigit(number,pos)	
			if val == 0 {
				occorrenze[0][pos]++
			} else {
				occorrenze[1][pos]++
			}
		}
	}

	fmt.Println(occorrenze)

	var gamma_rate string
	var delta_rate string 

	for i := 0; i < length_input ; i++ {
		if occorrenze[0][i] >= occorrenze[1][i] {
			gamma_rate += "0"
			delta_rate += "1"
		} else {
			gamma_rate += "1"
			delta_rate += "0"
		}
	}

	fmt.Println(gamma_rate)
	fmt.Println(delta_rate)
	fmt.Println(fromBinaryToDecimal(gamma_rate)*fromBinaryToDecimal(delta_rate))
}

func parseDigit(s string , pos int) int {
	digit, _ := strconv.Atoi(string(s[pos]))
	return digit
} 

func fromBinaryToDecimal(binary string) int {
	var decimal_value int
	for i := len(binary) - 1; i >= 0; i-- {
		cifra, _ := strconv.Atoi(string(binary[i]))
		decimal_value += cifra*pow(2,len(binary) - 1 - i)
	}
	fmt.Println(decimal_value)
	return decimal_value
}

func pow(n int ,e int ) int {
	var power = 1
	for i := 0; i < e; i++ {
		power *= n
	}
	return power
}