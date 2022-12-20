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
func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sum2 := 0

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

	print(lines)

	for _, l := range lines {
		for _, point := range l {
			// point.X -= min.X
			fmt.Printf("%d\n", min.Y)
			point.Y -= min.Y
		}
	}

	print(lines)
	fmt.Printf("%v\n", min)
	fmt.Printf("%v\n", max)

	// max.X -= min.X
	max.Y -= min.Y
	start := 500 - min.Y

	fmt.Printf("%v\n", min)
	fmt.Printf("%v\n", max)

	scannedMap := make([][]int, max.X+1)
	for i := 0; i < len(scannedMap); i++ {
		scannedMap[i] = make([]int, max.Y+1)
		fmt.Printf("%v\n", scannedMap[i])
	}

	for _, l := range lines {
		for i := 0; i < len(l)-1; i++ {
			if l[i].X == l[i+1].X {
				fmt.Printf("%v - %v\n", l[i], l[i+1])
				if l[i].Y < l[i+1].Y {
					for k := l[i].Y; k <= l[i+1].Y; k++ {
						// fmt.Printf("%d - %d", l[i].X, k)
						scannedMap[l[i].X][k] = 1
					}
				} else {
					for k := l[i+1].Y; k <= l[i].Y; k++ {
						// fmt.Printf("%d - %d", l[i].X, k)
						scannedMap[l[i].X][k] = 1
					}
				}
			}
			if l[i].Y == l[i+1].Y {
				if l[i].X < l[i+1].X {
					for k := l[i].X; k <= l[i+1].X; k++ {
						scannedMap[k][l[i].Y] = 1
					}
				} else {
					for k := l[i+1].X; k <= l[i].X; k++ {
						scannedMap[k][l[i].Y] = 1
					}
				}
			}
		}
	}

	// scannedMap[0][start] = 2

	for i := 0; i < len(scannedMap); i++ {
		for j := 0; j < len(scannedMap[i]); j++ {
			fmt.Printf("%v ", scannedMap[i][j])
		}
		fmt.Print("\n")
	}

	noMoreMove := false
	sum := 0
	for !noMoreMove {

		pos := Vec2{0, start}
		// movable := true
		for {
			if pos.X+1 < len(scannedMap) {
				if scannedMap[pos.X+1][pos.Y] == 0 {
					pos.X++
					continue
				}
			} else {
				noMoreMove = true
				break
			}

			if pos.X+1 < len(scannedMap) && pos.Y-1 > -1 {
				if scannedMap[pos.X+1][pos.Y-1] == 0 {
					pos.X++
					pos.Y--
					continue
				}
			} else {
				noMoreMove = true
				break
			}

			if pos.X+1 < len(scannedMap) && pos.Y+1 < len(scannedMap[pos.X+1]) {
				if scannedMap[pos.X+1][pos.Y+1] == 0 {
					pos.X++
					pos.Y++
					continue
				}
			} else {
				noMoreMove = true
				break
			}

			scannedMap[pos.X][pos.Y] = 2

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
	fmt.Printf("\npart1=%d, part2=%d", sum-1, sum2)
}
