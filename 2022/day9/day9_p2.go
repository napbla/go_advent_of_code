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

// func (p *position) print() string {
// 	return fmt.Sprintf("%d.%d\n", p.x, p.y)
// }

// func printDebug(r []position) {
// 	minX := r[0].x
// 	maxX := r[0].x
// 	minY := r[0].y
// 	maxY := r[0].y
// 	positions := make(map[string]int)
// 	for i := 0; i < len(r); i++ {
// 		// fmt.Printf("%d\n", r[i])
// 		if r[i].x > maxX {
// 			maxX = r[i].x
// 		}
// 		if r[i].x < minX {
// 			minX = r[i].x
// 		}

// 		if r[i].y > maxY {
// 			maxY = r[i].y
// 		}
// 		if r[i].y < minY {
// 			minY = r[i].y
// 		}
// 		positions[fmt.Sprintf("%d_%d", r[i].x, r[i].y)] = i + 1

// 	}

// 	for i := maxX; i >= minX; i-- {
// 		for j := minY; j <= maxY; j++ {
// 			if positions[fmt.Sprintf("%d_%d", i, j)] > 0 {
// 				fmt.Printf("%d", positions[fmt.Sprintf("%d_%d", i, j)]-1)
// 			} else {
// 				fmt.Printf(".")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// }

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	length := 10
	rope := make([]position, length)
	for i := 0; i < length; i++ {
		rope[i] = position{
			x: 0,
			y: 0,
		}
	}
	visited := make(map[string]int)
	sum := 0

	for scanner.Scan() {
		s := scanner.Text()
		step, err := strconv.Atoi(s[2:])
		// fmt.Printf("\nCommand = %c %d\n", s[0], step)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < step; i++ {
			// fmt.Printf("\n===step %d :\n", i)
			switch s[0] {
			case 'U':
				{
					rope[0].x += 1
				}
			case 'D':
				{
					rope[0].x -= 1
				}
			case 'R':
				{
					rope[0].y += 1
				}
			case 'L':
				{
					rope[0].y -= 1
				}
			}

			// fmt.Printf("r[0]=%v", rope[0])
			for j := 1; j < length; j++ {
				dx := rope[j-1].x - rope[j].x
				dy := rope[j-1].y - rope[j].y
				distance := (dx * dx) + (dy * dy)
				// fmt.Printf("distance:%v\n", distance)
				if distance > 2 {
					if dx != 1 && dx != -1 {
						dx = dx / 2
					}
					if dy != 1 && dy != -1 {
						dy = dy / 2
					}
					rope[j].x = rope[j].x + dx
					rope[j].y = rope[j].y + dy
				}
				// fmt.Printf("r[%d]=%v", j, rope[j])
			}
			// fmt.Printf("\n")
			// printDebug(rope)

			name := fmt.Sprintf("%d-%d", rope[length-1].x, rope[length-1].y)
			visited[name]++

			// fmt.Printf("name:%s : %d\n", name, visited[name])
			if visited[name] == 1 {
				sum++
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart2=%d\n", sum)
}
