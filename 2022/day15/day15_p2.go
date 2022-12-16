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
	X int64
	Y int64
}

func parseCoord(s string) Vec2 {
	splited := strings.Split(s, ",")
	coords := make([]int64, 2)
	for i := 0; i < len(splited); i++ {
		indexOfEqual := strings.Index(splited[i], "=")
		if indexOfEqual < 0 {
			log.Fatal("can not parse coordinates")
		}
		coord, err := strconv.Atoi(splited[i][indexOfEqual+1:])
		if err != nil {
			log.Fatalf("can not parse coordinates:%v", err)
		}
		coords[i] = int64(coord)
	}
	return Vec2{coords[0], coords[1]}
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func distance(a Vec2, b Vec2) int64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func IsOverlap(a int64, b int64, c int64, d int64) bool {
	return (c <= b) && (a <= d) //rule 2 lines must not be near each other
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nearestBeaconOf := make([]Vec2, 0)
	sensors := make([]Vec2, 0)

	for scanner.Scan() {
		s := scanner.Text()
		s = strings.Replace(s, "Sensor at ", "", 1)
		s = strings.Replace(s, ": closest beacon is at ", "|", 1)
		splitted := strings.Split(s, "|")
		sensors = append(sensors, parseCoord(splitted[0]))
		nearestBeaconOf = append(nearestBeaconOf, parseCoord(splitted[1]))
	}

	for lineY := int64(0); lineY < 4000000; lineY++ {
		covers := make([]*Vec2, 0)
		for i, sensor := range sensors {
			d := distance(sensor, nearestBeaconOf[i])

			if sensor.Y <= lineY {
				coverRange := int64(sensor.Y + d - lineY)
				if coverRange >= 0 {
					covers = append(covers, &Vec2{sensor.X - coverRange, sensor.X + coverRange})
				}
				continue
			}
			if sensor.Y > lineY {
				coverRange := (lineY - (sensor.Y - d))
				if coverRange >= 0 {
					covers = append(covers, &Vec2{sensor.X - coverRange, sensor.X + coverRange})
				}
				continue
			}
		}

		//cut boundary
		for _, cover := range covers {
			if cover.X < 0 {
				cover.X = 0
			}
			if cover.X > 4000000 {
				cover.X = 4000000
			}
			if cover.Y < 0 {
				cover.Y = 0
			}
			if cover.Y > 4000000 {
				cover.Y = 4000000
			}
		}

		visited := make([]bool, len(covers))
		for i := 0; i < len(covers); i++ {
			if visited[i] {
				continue
			}
			for j := 0; j < len(covers); j++ {
				if i != j && !visited[j] {
					if IsOverlap(covers[i].X, covers[i].Y, covers[j].X, covers[j].Y) {
						visited[j] = true
						// fmt.Printf("Merged %v and %v", covers[i], covers[j])
						covers[i].X = min(covers[i].X, covers[j].X)
						covers[i].Y = max(covers[i].Y, covers[j].Y)
						// fmt.Printf(" into %v\n", covers[i])
					}
				}
			}
		}

		sum := int64(0)

		for i := 0; i < len(covers); i++ {
			if !visited[i] {
				sum += (covers[i].Y - covers[i].X)
			}
		}

		if sum < 4000000 {
			fmt.Printf("lineY=%d\n", lineY)
			x := int64(0)

			//properly there will only 2 ranges left since it needs to cover from 0 to 4000000
			for i, c := range covers {
				if !visited[i] {
					x = c.Y + 1
					break
				}
			}

			fmt.Printf("x=%d\n", x)
			fmt.Printf("part2=%d\n", x*4000000+lineY)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// func print(lines [][]*Vec2) {
// 	for _, line := range lines {
// 		for _, point := range line {
// 			fmt.Printf("(%d,%d) ", point.X, point.Y)
// 		}
// 		fmt.Print("\n")
// 	}
// }

// func printCovers(v []*Vec2) {
// 	for _, vv := range v {
// 		fmt.Printf("{%d, %d} ", vv.X, vv.Y)
// 	}
// 	fmt.Printf("\n")
// }
