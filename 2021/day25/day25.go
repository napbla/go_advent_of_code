package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Print(bs [][]byte) {
	for i := 0; i < len(bs); i++ {
		for j := 0; j < len(bs[i]); j++ {
			fmt.Printf("%c", bs[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seaMap := make([][]byte, 0)
	//map width
	n := 0
	// map height
	m := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)
		if n == 0 {
			n = len(s)
		}
		seaMap = append(seaMap, []byte(s))
		m++
	}

	Print(seaMap)
	move := 1
	step := 0
	for move > 0 {
		move = 0

		//left move
		leftMoves := make([][]int, 0)
		//first pass - find move
		for i := 0; i < m; i++ {
			moves := make([]int, 0)
			for j := 0; j < n; j++ {
				if seaMap[i][j] == '>' && seaMap[i][(j+1)%n] == '.' {
					moves = append(moves, j)
				}
			}
			leftMoves = append(leftMoves, moves)
		}
		//second pass - do move
		for i := 0; i < len(leftMoves); i++ {
			for j := 0; j < len(leftMoves[i]); j++ {
				seaMap[i][leftMoves[i][j]], seaMap[i][(leftMoves[i][j]+1)%n] = '.', '>'
				move++
			}
		}
		// fmt.Printf("step=%d.5, move=%d\n\n", step, move)
		// Print(seaMap)

		//down move
		downMoves := make([][]int, 0)
		//first pass - find move
		for i := 0; i < n; i++ {
			moves := make([]int, 0)
			for j := 0; j < m; j++ {
				if seaMap[j][i] == 'v' && seaMap[(j+1)%m][i] == '.' {
					moves = append(moves, j)
				}
			}
			downMoves = append(downMoves, moves)
		}
		//second pass - do move
		for i := 0; i < len(downMoves); i++ {
			for j := 0; j < len(downMoves[i]); j++ {
				seaMap[downMoves[i][j]][i], seaMap[(downMoves[i][j]+1)%m][i] = '.', 'v'
				move++
			}
		}

		step++
		fmt.Printf("step=%d, move=%d\n", step, move)
		// Print(seaMap)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d\n", step)
}
