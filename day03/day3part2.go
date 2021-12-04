package main 

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	const length_input int = 12

	var elements[]string

	for {
		_, e := fmt.Scan(&number)
		if e != nil {
			break
		}

		elements = append(elements,number)
	}

	elements2 := make([]string,len(elements))
	copy(elements2,elements)
	gamma_rate := gamma_rate(elements)
	delta_rate := delta_rate(elements2)

	fmt.Println(gamma_rate,delta_rate)
	fmt.Println(fromBinaryToDecimal(gamma_rate)*fromBinaryToDecimal(delta_rate))

}

func createOccurrences(elements []string) [2][12]int {
	var occorrenze [2][12]int
	for _,v := range elements {
		for pos := 0; pos < 12; pos++ {
			val := parseDigit(v,pos)	
			if val == 0 {
				occorrenze[0][pos]++
			} else {
				occorrenze[1][pos]++
			}
		}
	}
	
	return occorrenze
}

func parseDigit(s string , pos int) int {
	digit, _ := strconv.Atoi(string(s[pos]))
	return digit
} 

func isMostCommon(val int, pos int, occorrenze [2][12]int) bool {
	if val == 1 {
		return occorrenze[1][pos] >= occorrenze[0][pos] 
	} else {
		return occorrenze[0][pos] > occorrenze[1][pos] 
	}
}


func gamma_rate(elements []string) string {
	for j := 0; j < 12 && len(elements) > 1; j++ {
		occorrenze := createOccurrences(elements)
		for i := len(elements) - 1; i >= 0; i-- {
			val,_ := strconv.Atoi(string(elements[i][j]))
			if !isMostCommon(val,j,occorrenze) {
				elements = append(elements[:i], elements[i+1:]...)
			}
		}
	}
	return elements[0]
}

func delta_rate(elements []string) string {
	for j := 0; j < 12 && len(elements) > 1; j++ {
		occorrenze := createOccurrences(elements)
		for i := len(elements) - 1; i >= 0; i-- {
			val,_ := strconv.Atoi(string(elements[i][j]))
			if isMostCommon(val,j,occorrenze) {
				elements = append(elements[:i], elements[i+1:]...)
			}
		}
		
	}
	return elements[0]
}

func fromBinaryToDecimal(binary string) int {
	var decimal_value int
	for i := len(binary) - 1; i >= 0; i-- {
		cifra, _ := strconv.Atoi(string(binary[i]))
		decimal_value += cifra*pow(2,len(binary) - 1 - i)
	}
	return decimal_value
}

func pow(n int ,e int ) int {
	var power = 1	
	for i := 0; i < e; i++ {
		power *= n
	}
	return power
}
