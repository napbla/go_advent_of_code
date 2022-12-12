package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

type comparer func(next, last byte) bool

func compare(n, l byte) bool {
	return int(n)-int(l) < 2
}

func compare2(n, l byte) bool {
	return int(l)-int(n) < 2
}

func createFlags(hillMap []string) [][]int {
	flags := make([][]int, len(hillMap))
	for i := 0; i < len(hillMap); i++ {
		flags[i] = make([]int, len(hillMap[i]))
		for j := 0; j < len(hillMap[i]); j++ {
			flags[i][j] = len(hillMap) * len(hillMap[0])
		}
	}
	return flags
}

func printDebugFlags(flags [][]int, min int) {
	for i := 0; i < len(flags); i++ {
		for j := 0; j < len(flags[i]); j++ {
			if flags[i][j] == min {
				fmt.Printf("  0_")
			} else {
				fmt.Printf("%3d_", flags[i][j])
			}
		}
		fmt.Printf("\n")
	}

}

// Dijkstra using Dynamic programming (recursive call + caching), kinda similar to Flood Fill
//
//	We iterate 4 directions , caching the distance in flags array.
//	For one direction, if the distance from start to the next position which is stores in flags array
//
// smaller than current distance + 1, we will move to that direction. If not there exists some
// path with smaller distance so we can ignore it.
func explore(i int, j int, sum int, hillMap []string, flags [][]int, f comparer) {
	flags[i][j] = sum

	if i-1 > -1 && f(hillMap[i-1][j], hillMap[i][j]) {
		if flags[i-1][j]-1 > sum {
			explore(i-1, j, sum+1, hillMap, flags, f)
		}
	}

	if i+1 < len(hillMap) && f(hillMap[i+1][j], hillMap[i][j]) {
		if flags[i+1][j]-1 > sum {
			explore(i+1, j, sum+1, hillMap, flags, f)
		}
	}

	if j-1 > -1 && f(hillMap[i][j-1], hillMap[i][j]) {
		if flags[i][j-1]-1 > sum {
			explore(i, j-1, sum+1, hillMap, flags, f)
		}
	}

	if j+1 < len(hillMap[i]) && f(hillMap[i][j+1], hillMap[i][j]) {
		if flags[i][j+1]-1 > sum {
			explore(i, j+1, sum+1, hillMap, flags, f)
		}
	}
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hillMap := make([]string, 0)
	count := 0
	var start position
	var end position

	debug := true

	for scanner.Scan() {
		s := scanner.Text()
		hillMap = append(hillMap, s)
		if p := strings.IndexByte(s, 'S'); p > -1 {
			start = position{
				x: count,
				y: p,
			}
		}
		if p := strings.IndexByte(s, 'E'); p > -1 {
			end = position{
				x: count,
				y: p,
			}
		}
		count++
	}

	hillMap[start.x] = strings.Replace(hillMap[start.x], "S", "a", 1)
	hillMap[end.x] = strings.Replace(hillMap[end.x], "E", "z", 1)

	min := len(hillMap) * len(hillMap[0])
	flags := createFlags(hillMap)

	explore(start.x, start.y, 0, hillMap, flags, compare)

	if debug {
		printDebugFlags(flags, min)
	}

	fmt.Printf("part1 = %v\n", flags[end.x][end.y])

	//part2
	//reset flags
	flags = createFlags(hillMap)

	explore(end.x, end.y, 0, hillMap, flags, compare2)

	if debug {
		printDebugFlags(flags, min)
	}

	//find smallest path from E to any 'a'
	//reuse min variable
	for i := 0; i < len(hillMap); i++ {
		for j := 0; j < len(hillMap[i]); j++ {
			if hillMap[i][j] == 'a' && flags[i][j] < min {
				min = flags[i][j]
			}
		}
	}
	fmt.Printf("part2 = %v\n", min)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
