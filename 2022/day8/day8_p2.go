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

	visibleScore := make([][]int, len(trees))
	for i := 1; i < len(trees)-1; i++ {
		visibleScore[i] = make([]int, len(trees[i]))
		for j := 1; j < len(trees[i])-1; j++ {
			//left
			left := 0
			for k := j - 1; k > -1; k-- {
				left++
				if trees[i][j] <= trees[i][k] {
					break
				}

			}
			//right
			right := 0
			for k := j + 1; k < len(trees[i]); k++ {
				right++
				if trees[i][j] <= trees[i][k] {
					break
				}

			}
			//up
			up := 0
			for k := i - 1; k > -1; k-- {
				up++
				if trees[i][j] <= trees[k][j] {
					break
				}

			}
			//down
			down := 0
			for k := i + 1; k < len(trees); k++ {
				down++
				if trees[i][j] <= trees[k][j] {
					break
				}
			}

			// debug
			// if i == 3 && j == 2 {
			// 	fmt.Printf("left:%d right:%d up:%d down:%d\n", left, right, up, down)
			// }

			visibleScore[i][j] = left * right * up * down
		}

	}

	//count visible
	max := 0
	for i := 0; i < len(visibleScore); i++ {
		// debug
		// fmt.Printf("%v\n", visibleScore[i])
		for j := 0; j < len(visibleScore[i]); j++ {
			if visibleScore[i][j] > max {
				max = visibleScore[i][j]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart2=%d\n", max)
}
