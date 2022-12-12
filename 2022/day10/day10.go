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

func updateCycle(cycle int, X int) int {
	// fmt.Printf("cycle:%d, X:%d\n", cycle, X)
	if cycle == 20 || cycle == 60 || cycle == 100 ||
		cycle == 140 || cycle == 180 || cycle == 220 {
		sum := cycle * X
		fmt.Printf("cycle:%d, sum:%d\n", cycle, sum)
		return sum
	}
	return 0
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	cycle := 0
	X := 1

	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Printf("command:%s\n", s[:4])
		switch s[:4] {
		case "noop":
			{
				cycle++
				sum += updateCycle(cycle, X)
			}
		case "addx":
			{
				op, err := strconv.Atoi(s[5:])
				// fmt.Printf("op:%d\n", op)
				// fmt.Printf("\nCommand = %c %d\n", s[0], step)
				if err != nil {
					log.Fatal(err)
				}
				cycle++
				sum += updateCycle(cycle, X)

				cycle++
				sum += updateCycle(cycle, X)
				X += op
				fmt.Printf("X=%d\n", X)

			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d\n", sum)
}
