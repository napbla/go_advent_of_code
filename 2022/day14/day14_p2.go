package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vec2 struct {
	X int
	Y int
}

// 498,4 -> 498,6 -> 496,6
// 503,4 -> 502,4 -> 502,9 -> 494,9

// 4     5  5
// 9     0  0
// 4     0  3
// 0 ......+...
// 1 ..........
// 2 ..........
// 3 ..........
// 4 ....#...##
// 5 ....#...#.
// 6 ..###...#.
// 7 ........#.
// 8 ........#.
// 9 #########.

func print(lines [][]*Vec2) {
	for _, line := range lines {
		for _, point := range line {
			fmt.Printf("(%d,%d) ", point.X, point.Y)
		}
		fmt.Print("\n")
	}

}

func GetPos(sm map[Vec2]int, pos Vec2, max int) int {
	if pos.X >= max+2 {
		return 1
	}
	return sm[pos]
}
func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	min := Vec2{0, 0}
	max := Vec2{0, 0}

	lines := make([][]*Vec2, 0)
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		splitted := strings.Split(s, " -> ")

		positions := make([]*Vec2, len(splitted))
		for i, pos := range splitted {
			coords := strings.Split(pos, ",")
			x, err2 := strconv.Atoi(coords[1])
			if err2 != nil {
				log.Fatalf("%v", err2)
			}
			y, err := strconv.Atoi(coords[0])
			if err != nil {
				log.Fatalf("%v", err)
			}
			if count == 0 {
				min = Vec2{x, y}
			}

			if min.X > x {
				min.X = x
			}
			if min.Y > y {
				min.Y = y
			}
			if max.X < x {
				max.X = x
			}
			if max.Y < y {
				max.Y = y
			}
			positions[i] = &Vec2{x, y}
		}
		lines = append(lines, positions)
		// lines[count] = positions
		count++
	}

	// print(lines)
	// fmt.Printf("%v\n", min)
	// fmt.Printf("%v\n", max.X)

	scannedMap := make(map[Vec2]int)

	for _, l := range lines {
		for i := 0; i < len(l)-1; i++ {
			if l[i].X == l[i+1].X {
				if l[i].Y < l[i+1].Y {
					for k := l[i].Y; k <= l[i+1].Y; k++ {
						scannedMap[Vec2{l[i].X, k}] = 1
					}
				} else {
					for k := l[i+1].Y; k <= l[i].Y; k++ {
						scannedMap[Vec2{l[i].X, k}] = 1
					}
				}
			}
			if l[i].Y == l[i+1].Y {
				if l[i].X < l[i+1].X {
					for k := l[i].X; k <= l[i+1].X; k++ {
						scannedMap[Vec2{k, l[i].Y}] = 1
					}
				} else {
					for k := l[i+1].X; k <= l[i].X; k++ {
						scannedMap[Vec2{k, l[i].Y}] = 1
					}
				}
			}
		}
	}

	noMoreMove := false
	sum := 0
	for !noMoreMove {
		pos := Vec2{0, 500}
		// movable := true
		for {
			if GetPos(scannedMap, Vec2{pos.X + 1, pos.Y}, max.X) == 0 {
				pos.X++
				continue
			}

			if GetPos(scannedMap, Vec2{pos.X + 1, pos.Y - 1}, max.X) == 0 {
				pos.X++
				pos.Y--
				continue
			}

			if GetPos(scannedMap, Vec2{pos.X + 1, pos.Y + 1}, max.X) == 0 {
				pos.X++
				pos.Y++
				continue
			}

			scannedMap[Vec2{pos.X, pos.Y}] = 2
			if pos.X == 0 && pos.Y == 500 {
				noMoreMove = true
			}

			// fmt.Print("\n")
			// for i := 0; i < len(scannedMap); i++ {
			// 	for j := 0; j < len(scannedMap[i]); j++ {
			// 		c := '.'
			// 		if scannedMap[i][j] == 1 {
			// 			c = '#'
			// 		}
			// 		if scannedMap[i][j] == 2 {
			// 			c = 'o'
			// 		}
			// 		fmt.Printf("%c", c)
			// 	}
			// 	fmt.Print("\n")
			// }
			break
		}
		sum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\npart2=%d", sum)
}
