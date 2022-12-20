package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func find(x Vec2, arr []Vec2) int {
	found := -1
	for i := range arr {
		if arr[i].X == x.X && arr[i].Y == x.Y {
			found = i
			break
		}
	}
	return found
}

type Vec2 struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	operations := make([]Vec2, 0)
	arr := make([]Vec2, 0)
	zero := Vec2{0, 0}

	count := 0
	for scanner.Scan() {
		s := scanner.Text()
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("can not parse int:%v", err)
		}
		operations = append(operations, Vec2{x, count})
		arr = append(arr, Vec2{x, count})
		if x == 0 {
			zero.Y = count
		}
		count++

	}

	for _, op := range operations {
		t := op.X
		if t == 0 {
			continue
		}
		index := find(op, arr)

		t = index + t
		t = t % (len(arr) - 1)
		if t <= 0 {
			t += (len(arr) - 1)
		}

		if index < t {
			for i := index; i < t; i++ {
				arr[i] = arr[i+1]
			}
			arr[t] = op
		}

		if index > t {
			for i := index; i > t; i-- {
				arr[i] = arr[i-1]
			}
			arr[t] = op
		}
	}

	indexOf0 := find(zero, arr)

	x := arr[(indexOf0+1000)%len(arr)].X
	y := arr[(indexOf0+2000)%len(arr)].X
	z := arr[(indexOf0+3000)%len(arr)].X

	// fmt.Printf("1000th=%d\n", x)
	// fmt.Printf("2000th=%d\n", y)
	// fmt.Printf("3000th=%d\n", z)
	fmt.Printf("part1=%d\n", x+y+z)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
