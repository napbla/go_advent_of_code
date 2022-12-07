package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func split(s string) (string, string) {
	return s[:len(s)/2], s[len(s)/2:]
}

func finddup(a string, b string) []byte {
	aHas := make(map[byte]int)
	res := make([]byte, 0)

	for i := 0; i < len(a); i++ {
		aHas[a[i]]++
	}

	for i := 0; i < len(b); i++ {
		if aHas[b[i]] > 0 {
			aHas[b[i]] = -aHas[b[i]]
			res = append(res, b[i])
		}
	}

	return res
}

func score(b byte) int {
	if b < 91 {
		return int(b) - 65 + 27
	}
	return int(b) - 96
}

func calScore(bs []byte) int {
	s := 0
	for i := 0; i < len(bs); i++ {
		s += score(bs[i])
	}
	return s
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	sum2 := 0
	count := 0
	temp := make([]string, 3)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		//part 1
		a, b := split(s)
		sum += calScore(finddup(a, b))

		//part 2
		temp[count] = s
		count++

		if count == 3 {
			count = 0
			// fmt.Printf("%s\t%s\t%s\n", temp[0], temp[1], temp[2])
			sum2 += calScore(finddup(temp[2], string(finddup(temp[0], temp[1])[:])))
		}
		// fmt.Printf("%s\t%s\n", a, b)
		// fmt.Printf("dup=%c\n", finddup(a, b)[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d, part2=%d\n", sum, sum2)
}
