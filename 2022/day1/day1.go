package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type topN struct {
	count    int
	buf      []int
	n        int
	min      int
	minIndex int
}

func NewTopN(max int) *topN {
	return &topN{
		// count: 0,
		buf:      make([]int, max),
		n:        max,
		count:    0,
		min:      0,
		minIndex: 0,
	}
}

func (t *topN) Insert(val int) {
	//init min
	if t.count == 0 {
		t.min, t.minIndex = val, 0
	}
	//
	if t.count < t.n {
		t.buf[t.count] = val
		t.count++
		if val < t.min {
			t.min = val
			t.minIndex = t.count
		}
		return
	}
	//replace min
	if val > t.min {
		t.buf[t.minIndex] = val
		t.min = val

		//find new min
		for i := 0; i < t.n; i++ {
			if t.min > t.buf[i] {
				t.min = t.buf[i]
				t.minIndex = i
			}
		}
	}
}

func (t *topN) Sum() int {
	sum := 0
	for i := 0; i < t.n; i++ {
		sum += t.buf[i]
	}
	return sum
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	//for part 1 change to m := NewTopN(1)
	top := NewTopN(1)
	top3 := NewTopN(3)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		if s != "" {
			intVar, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("err=%v\n", err)
				return
			}
			sum += intVar
		} else {
			//reset
			top.Insert(sum)
			top3.Insert(sum)
			sum = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d, part2=%d\n", top.Sum(), top3.Sum())
}
