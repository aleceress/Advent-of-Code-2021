package main

import (
	"fmt"
	"math"
)

//a coordinate is a struct with x and y fields
type Coord struct {
	x int 
	y int
}

func main() {

	//coordinates of points given in input
	var x1,x2,y1,y2 int 
	/*
	we create a map to memorize lines: a line starts from a coordinate and ends into an other 
	coordinate (the map has a list of coordinates in case there are more lines that start from the same point)
	*/
	var lines = make(map[Coord][]Coord)

	var maxX = math.MinInt64
	var maxY = math.MinInt64

	for {
		_, err := fmt.Scanf("%d,%d -> %d,%d\n",&x1,&y1,&x2,&y2)

		/*meanwhile we search for the max and min y, which we'll need later
		to set up the grid
		*/
		maxX = updateMax(maxX,x1)
		maxX = updateMax(maxX,x2)
		maxY = updateMax(maxY,y1)
		maxY = updateMax(maxY,y2)
		
		if err != nil {
			break
		}

		start_coord := Coord{x: x1, y: y1}
		end_coord := Coord{x: x2, y: y2}

		lines[start_coord] = append(lines[start_coord],end_coord)
	}

	var grid[][]int 
	//initialize grid
	for i := 0; i <= maxX; i++ {
		grid = make([][]int,maxX + 1)
		for j := 0; j <= maxX; j++ {
			grid[j] = make([]int,maxX + 1)
		}
	}

	//we iterate on the map that contains lines
	for k,v := range lines {
		for _,c := range v {
			//if it's a vertical movement 
			if k.y == c.y {
				//if the starting coordinate is over the ending coordinate
				if k.x <= c.x {
					for i:= k.x; i <= c.x; i++{
						grid[i][k.y]++
					}	
				} else {
					for i:= c.x; i <= k.x; i++{
						grid[i][k.y]++
					}	
				}
				
			}

			//if it's an horizontal movement 
			if k.x == c.x {
				//if the starting coordinate is on the left of the first coordinate
				if k.y <= c.y {
					for j:= k.y; j <= c.y; j++ {
						grid[k.x][j]++
					}
				} else {
					for j:= c.y; j <= k.y; j++ {
						grid[k.x][j]++
					}
				}
				
			}

			//if it's a movement on the antidiagonal
			if  (c.x - k.x) == (k.y - c.y) {
				//if the starting coordinate is on the left of the first coordinate
				if k.x <= c.x {
					for i:= 0; i <= c.x - k.x; i++ {
						grid[k.x + i][k.y - i]++
					}
				} else {
					for i:= 0; i <= k.x - c.x; i++ {
						grid[k.x - i][k.y + i]++
					}
				}
			}

			//if it's a movement on the main diagonal
			if  (c.x - k.x) == (c.y - k.y) {
				//if the starting coordinate is on the left of the first coordinate
				if k.x <= c.x {
					for i:= 0; i <= c.x - k.x; i++ {
						grid[k.x + i][k.y + i]++
					}
				} else {
					for i:= 0; i <= k.x - c.x; i++ {
						grid[k.x - i][k.y - i]++
					}
				}
			}
		}
	}

	var count_overlaps int
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxX; j++ {
			if grid[i][j] >=2 {
				count_overlaps++
			}
		}
	}

	fmt.Println(count_overlaps)
}

func updateMax(max_till_now int, new_element int) int {
	if new_element >= max_till_now {
		max_till_now = new_element
	}
	return max_till_now
}
