package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
)

type position struct {
	x int
	y int
}

// EHPZPJGL
func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// sum := 0
	cycle := 0
	X := 1

	currentSprite := uint64(7) // .......111
	screen := uint64(0)

	// fmt.Printf("1<<0 = %b", 1<<0)

	for scanner.Scan() {
		s := scanner.Text()
		switch s[:4] {
		case "noop":
			{
				cycle++

				mask := uint64(1) << ((cycle - 1) % 40)
				screen = (currentSprite & mask) | screen

				// fmt.Printf("crspr: %b , mask: %b\n, screen: %b\n", currentSprite, mask, screen)

				if cycle == 40 || cycle == 80 || cycle == 120 ||
					cycle == 160 || cycle == 200 || cycle == 240 {
					fmt.Printf("%39b\n", bits.Reverse64(screen))
					screen = uint64(0)
				}

			}
		case "addx":
			{
				op, err := strconv.Atoi(s[5:])
				if err != nil {
					log.Fatal(err)
				}
				//1st cycle
				cycle++

				mask := uint64(1) << ((cycle - 1) % 40)
				screen = (currentSprite & mask) | screen

				if cycle == 40 || cycle == 80 || cycle == 120 ||
					cycle == 160 || cycle == 200 || cycle == 240 {
					fmt.Printf("%39b\n", bits.Reverse64(screen))
					screen = uint64(0)
				}

				//2nd cycle

				cycle++
				mask = uint64(1) << ((cycle - 1) % 40)
				screen = (currentSprite & mask) | screen

				if cycle == 40 || cycle == 80 || cycle == 120 ||
					cycle == 160 || cycle == 200 || cycle == 240 {
					fmt.Printf("%39b\n", bits.Reverse64(screen))
					screen = uint64(0)
				}

				X += op

				switch X {
				case -1:
					{
						currentSprite = uint64(1)
					}
				case 0:
					{
						currentSprite = uint64(3)
					}
				case 1:
					{
						currentSprite = uint64(7)
					}
				default:
					{
						currentSprite = uint64(7) << (X - 1)
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
