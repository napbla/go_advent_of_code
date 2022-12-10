package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type position struct {
	x int
	y int
}

func (p *position) print() string {
	return fmt.Sprintf("%d.%d\n", p.x, p.y)
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tailPositions := []*position{{
		x: 0,
		y: 0,
	}}

	headPosition := position{
		x: 0,
		y: 0,
	}

	scanner := bufio.NewScanner(file)
	max := 0
	min := 0
	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)
		step, err := strconv.Atoi(s[2:])
		// fmt.Printf("%d\n", step)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < step; i++ {
			prevHead := headPosition

			switch s[0] {
			case 'U':
				{
					headPosition.x += 1
				}
			case 'D':
				{
					headPosition.x -= 1
				}
			case 'R':
				{
					headPosition.y += 1
				}
			case 'L':
				{
					headPosition.y -= 1
				}
			}

			if headPosition.x > max {
				max = headPosition.x
			}
			if headPosition.y > max {
				max = headPosition.y
			}
			if headPosition.x < min {
				min = headPosition.x
			}
			if headPosition.y < min {
				min = headPosition.y
			}

			dx := headPosition.x - tailPositions[len(tailPositions)-1].x
			dy := headPosition.y - tailPositions[len(tailPositions)-1].y
			distance := (dx * dx) + (dy * dy)
			// fmt.Printf("distance:%v\n", distance)
			if distance > 2 {
				// fmt.Printf("add %v ; head: %v\n", prevHead.print(), headPosition.print())
				tailPositions = append(tailPositions, &prevHead)
			}
		}

	}

	sum := 0
	temp := make(map[int]int)
	// fmt.Printf("max:%d\n", max)

	// debug := make([][]int, max)
	// for i := 0; i < max; i++ {
	// 	debug[i] = make([]int, max)
	// }

	if min < 0 {
		max = max - min
	}

	for _, pos := range tailPositions {
		// debug[pos.x][pos.y] = 1
		temp[pos.x*max+pos.y]++
		if temp[pos.x*max+pos.y] == 1 {
			sum++
		}
	}

	// for i := len(debug) - 1; i > -1; i-- {
	// 	for j := 0; j < len(debug[i]); j++ {
	// 		fmt.Printf("%d ", debug[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d\n", sum)
}
