package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//part 1 variables
	sum := 0
	last := 0

	//part 2 variables
	//// cache for sliding window, size 4
	count := 0
	cache := make([]int, 4)
	////
	sum2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		a, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		//part 1
		if a > last {
			sum++
		}
		last = a
		//part 2
		cache[count%4] = a
		if count > 2 {
			//s1: previous window, s2: current window
			s1 := 0
			s2 := 0
			for i := 0; i < 3; i++ {
				s1 += cache[(count-i-1)%4]
				s2 += cache[(count-i)%4]
			}
			if s2 > s1 {
				sum2++
			}
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d, part2=%d\n", sum-1, sum2)
}
