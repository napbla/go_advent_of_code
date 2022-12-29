package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var parseMap map[rune]string = map[rune]string{
	'R': "R",
	'L': "L",
}

type Vec2 struct {
	X int
	Y int
}

func Add(a Vec2, b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func parseCommands(s string) []string {
	res := make([]string, 0)
	numberBegin := -1
	for i, c := range s {
		if c >= '0' && c <= '9' {
			if numberBegin == -1 {
				numberBegin = i
			}
		}
		if c == 'R' || c == 'L' {
			res = append(res, s[numberBegin:i])
			res = append(res, parseMap[c])
			numberBegin = -1
		}
	}
	if numberBegin != -1 {
		res = append(res, s[numberBegin:])
	}
	return res
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rawMap := make([]string, 0)
	mapWidth := 0

	commands := make([]string, 0)

	mapEnded := false

	for scanner.Scan() {
		s := scanner.Text()

		if !mapEnded {
			rawMap = append(rawMap, s)
			if len(s) > mapWidth {
				mapWidth = len(s)
			}
		}

		if s == "" {
			mapEnded = true
			continue
		}
		if mapEnded {
			commands = parseCommands(s)
		}
	}

	// fmt.Printf("map width=%v\n", mapWidth)
	// fmt.Printf("commands=%v\n", commands)

	myMap := make([][]byte, len(rawMap))
	for i, s := range rawMap {
		myMap[i] = make([]byte, mapWidth)
		for j := 0; j < len(s); j++ {
			myMap[i][j] = s[j]
		}
	}

	//find start location
	current := Vec2{0, 0}
	for i, c := range myMap[0] {
		if c == '.' {
			current.Y = i
			break
		}
	}

	right := Vec2{0, 1}
	left := Vec2{0, -1}
	up := Vec2{-1, 0}
	down := Vec2{1, 0}

	directions := []Vec2{right, down, left, up}
	currentDirection := 0 //directions[0] = right

	for _, command := range commands {
		// fmt.Print("\n")
		// for i, m := range myMap {
		// 	for j, c := range m {
		// 		if current.X == i && current.Y == j {
		// 			fmt.Print("X ")
		// 		} else {
		// 			fmt.Printf("%c ", c)
		// 		}

		// 	}
		// 	fmt.Print("\n")
		// }
		// fmt.Printf("\n")

		if command == "R" {
			currentDirection = (currentDirection + 1) % 4
			continue
		}
		if command == "L" {
			currentDirection = currentDirection - 1
			if currentDirection < 0 {
				currentDirection += 4
			}
			continue
		}
		count, err := strconv.Atoi(command)
		if err != nil {
			log.Fatalf("can not convert %s:%v\n", command, err)
		}
		next := current
		for i := 0; i < count; i++ {

			next = Add(next, directions[currentDirection])
			if next.X < 0 {
				next.X = len(myMap) - 1
			} else {
				next.X %= len(myMap)
			}
			if next.Y < 0 {
				next.Y = mapWidth - 1
			} else {
				next.Y %= mapWidth
			}
			// fmt.Printf("%v | ", next)

			for myMap[next.X][next.Y] == ' ' || myMap[next.X][next.Y] == 0 {
				next = Add(next, directions[currentDirection])
				if next.X < 0 {
					next.X = len(myMap) - 1
				} else {
					next.X %= len(myMap)
				}
				if next.Y < 0 {
					next.Y = mapWidth - 1
				} else {
					next.Y %= mapWidth
				}
				// fmt.Printf("%v+%c| ", next, myMap[next.X][next.Y])
			}

			if myMap[next.X][next.Y] == '.' {
				current = next
				continue
			}
			if myMap[next.X][next.Y] == '#' {
				break
			}
		}
	}

	fmt.Printf("row=%d, col=%d, face=%d\n", current.X, current.Y, currentDirection)
	fmt.Printf("part1=%d\n", 1000*(current.X+1)+4*(current.Y+1)+currentDirection)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
