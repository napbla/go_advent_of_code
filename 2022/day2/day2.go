package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resultScore := map[string]int{
		"A X": 3, //rock = rock
		"A Y": 6, //rock < paper
		"A Z": 0, //rock > scissors

		"B X": 0, //paper > rock
		"B Y": 3, //paper = paper
		"B Z": 6, //paper < scissors

		"C X": 6, //scissors < rock
		"C Y": 0, //scissors > paper
		"C Z": 3, //scissors = scissors
	}
	moveScore := map[byte]int{
		'X': 1, //rock
		'Y': 2, //paper
		'Z': 3, //scissors
	}

	// X lose
	// Y draw
	// Z win
	resultScore2 := map[string]int{
		"A X": 0 + 3, //rock , lose -> scissors
		"A Y": 3 + 1, //rock , draw -> rock
		"A Z": 6 + 2, //rock , win -> paper

		"B X": 0 + 1, //paper , lose -> rock
		"B Y": 3 + 2, //paper , draw -> paper
		"B Z": 6 + 3, //paper , win -> scissors

		"C X": 0 + 2, //scissors , lose -> paper
		"C Y": 3 + 3, //scissors , draw -> scissors
		"C Z": 6 + 1, //scissors ,win -> rock
	}
	sum := 0
	sum2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)

		sum += (resultScore[s] + moveScore[s[len(s)-1]])

		sum2 += resultScore2[s]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d, part2=%d\n", sum, sum2)
}
