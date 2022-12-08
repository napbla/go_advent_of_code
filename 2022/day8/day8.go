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

	trees := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)
		trees = append(trees, []byte(s))
	}

	isVisible := make([][]bool, len(trees))
	for i := 0; i < len(trees); i++ {
		v := make([]bool, len(trees[i]))

		v[0] = true
		max := trees[i][0]

		//left
		for j := 1; j < len(trees[i]); j++ {
			if trees[i][j] > max {
				max = trees[i][j]
				v[j] = true
			}
		}

		v[len(trees[i])-1] = true
		max = trees[i][len(trees[i])-1]

		//right
		for j := len(trees[i]) - 2; j > 0; j-- {
			if trees[i][j] > max {
				max = trees[i][j]
				v[j] = true
			}
		}

		isVisible[i] = v
	}

	for i := 0; i < len(trees[0]); i++ {
		isVisible[0][i] = true
		max := trees[0][i]
		//top
		for j := 1; j < len(trees); j++ {
			if trees[j][i] > max {
				isVisible[j][i] = true
				max = trees[j][i]
			}
		}
		//bottom
		isVisible[len(trees)-1][i] = true
		max = trees[len(trees)-1][i]
		for j := len(trees) - 2; j > 0; j-- {
			if trees[j][i] > max {
				isVisible[j][i] = true
				max = trees[j][i]
			}
		}

	}

	//count visible
	sum := 0
	for i := 0; i < len(isVisible); i++ {
		for j := 0; j < len(isVisible[i]); j++ {
			if isVisible[i][j] {
				sum += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d, part2=%d\n", sum, sum)
}
