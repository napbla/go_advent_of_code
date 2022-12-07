package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// check if it is totally included. there are 4 cases:
//  1/                   2/
//	a----------b  and       a-----b
//	   c----d            c-------------d

//	=> (a - c) < 0 and (b - d) > 0
//  => (a-c)*(b-d) < 0                 (1)
//
//  3/                   4/
//	a---------b   and     a------b
//	c-----d                  c---d
//	=> (a-c) = 0 or b-d = 0            (2)

// (1) & (2) => (a-c)*(b-d) <= 0
func checkIsIncluded(a, b, c, d int) bool {
	return (a-c)*(b-d) <= 0
}

// we check if not included instead. there are 2 cases
// 1/                      |    2/
// a-----b     c------d    |    c------d     a-------b
// => (b < c) || (d < a)
// then not of not included is included
func checkIsIncluded2(a, b, c, d int) bool {
	return !((b < c) || (d < a))
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		splitted := strings.Split(s, ",")
		head, tail := splitted[0], splitted[1]
		headSplitted := strings.Split(head, "-")
		a, err := strconv.Atoi(headSplitted[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(headSplitted[1])
		if err != nil {
			log.Fatal(err)
		}
		tailSplitted := strings.Split(tail, "-")
		c, err := strconv.Atoi(tailSplitted[0])
		if err != nil {
			log.Fatal(err)
		}
		d, err := strconv.Atoi(tailSplitted[1])
		if err != nil {
			log.Fatal(err)
		}

		if checkIsIncluded(a, b, c, d) {
			sum++
		}
		if checkIsIncluded2(a, b, c, d) {
			sum2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\npart1=%d, part2=%d", sum, sum2)
}
