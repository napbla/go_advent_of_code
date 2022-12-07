package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findMarkerOfSize(s string, size int) int {
	pos := 0
	for i := 0; i < len(s); i++ {
		temp := make(map[byte]int)
		unique := true
		for j := 0; j < size; j++ {
			if _, ok := temp[s[i+j]]; ok || temp[s[i+j]] > 0 {
				unique = false
				break
			}
			temp[s[i+j]]++
		}
		if unique {
			pos = i
			break
		}
	}
	return pos + size
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()

		part1 := findMarkerOfSize(s, 4)
		part2 := findMarkerOfSize(s, 14)
		fmt.Printf("\npart1=%d ,part2=%d\n", part1, part2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
