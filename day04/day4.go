package main 

import (
	"fmt"
)

//every cell of the matrix representing the board has a number and a boolean indicating if the cell is marked or not
type Cell struct {
	number int 
	marked bool
}

//the score of every board (moves to win, value, last extracted number that leads to winning)
type Score struct {
	moves int 
	points int 
	mark int
}

func main() {

	inputs := []int{79,9,13,43,53,51,40,47,56,27,0,14,33,60,61,36,72,48,83,42,10,86,41,75,16,80,15,93,95,45,68,96,84,11,85,63,18,31,35,74,71,91,39,88,55,6,21,12,58,29,69,37,44,98,89,78,17,64,59,76,54,30,65,82,28,50,32,77,66,24,1,70,92,23,8,49,38,73,94,26,22,34,97,25,87,19,57,7,2,3,46,67,90,62,20,5,52,99,81,4}
	//inputs := []int{7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1}
	
	//every board
	var matrix[][]Cell
	var scores []Score 

	//the boards are NxN
	const N = 5; 
	var err error

	for count := 0; ; count++ {
		// initialize matrix
		for i := 0; i < n; i++ {
			matrix = make([][]Cell,n)
			for j := 0; j < n; j++ {
				matrix[j] = make([]Cell,n)
			}
		}

		// populate matrix
		for i := 0; i < n; i++ {
			for j := 0; j<n; j++ {
				var number int 
				_, err = fmt.Scan(&number)
				if err != nil {
					break
				}
				matrix[i][j].number = number
			}	
		}

		//counter of the moves needed to win
		moves := 0
		for _,v := range inputs {
			mark(matrix, v)

			if haveIWon(matrix) {
				points := calculatePoints(matrix)
				var s Score

				//number of moves needed for that board to win 
				s.moves = moves
				s.points = points
				s.mark = v //last number extracted
				scores = append(scores,s)
				break
			}
			moves++
		}
	}

	//min := min(scores) for part 1
	max := max(scores)
	fmt.Println(max.mark*max.points)
	
}

func mark(matrix [][]Cell, value int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j].number == value {
				matrix[i][j].marked = true
			}
		}
	}
}

func haveIWon(matrix [][]Cell) bool{

	//looking for a complete row
	for i:= 0; i < len(matrix); i++ {
		var won = true
		for j := 0; j < len(matrix); j++ {
			if !matrix[i][j].marked {
				won = false
			}
		}
		if won {
			return true
		}
	}

	//looking for a complete column
	for j:= 0; j < len(matrix); j++ {
		var won = true
		for i := 0; i < len(matrix); i++ {
			if !matrix[i][j].marked {
				won = false
			}
		}
		if won {
			return true
		}
	}

	return false
}

//find the board that lets us win with the minimum number of moves
func min(scores []Score) Score {
	var min Score
	min.moves = scores[0].moves
	min.points = 0

	for _,v := range scores {
		if v.moves < min.moves {
			min = v
		}
	}
	return min 
}

//find the board that lets us win with the maximum number of moves
func max(scores []Score) Score {
	var max Score
	max.moves = scores[0].moves
	max.points = 0

	for _,v := range scores {
		if v.moves > max.moves {
			max = v
		}
	}
	return max
}

//computes the value of a board
func calculatePoints(matrix [][]Cell) int {
	var points int
	for i := 0; i< len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if !matrix[i][j].marked {
				points += matrix[i][j].number
			}
		}
	}
	return points
}