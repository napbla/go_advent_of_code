package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var parseMap map[rune]string = map[rune]string{
	'R': "R",
	'L': "L",
}

type Vec2 struct {
	X int
	Y int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// sum range
// A:-2 -1 0 1 2
// A: -2 -1 0 1 2
// C: -1 0 1
// -5 -4 -3 ..... 3 4 5
// -5 => C=-1, S=0
// -4 = -5 + 1 => CARRY=-1, S=1
// -3 = -5 + 2 => C=-1, S=2
// 3 = 5 - 2 => C=1, S=-2
// 4 = 5 - 1 => C=1, S=-1
// 5 = 5 - 0 => C=1, S=0
func Add(a []int8, b []int8) []int8 {
	resLength := max(len(a), len(b))
	result := make([]int8, resLength)

	padding := make([]int8, 0)
	for i := 0; i < resLength-len(a); i++ {
		padding = append(padding, 0)
	}
	a = append(padding, a...)

	padding = make([]int8, 0)
	for i := 0; i < resLength-len(b); i++ {
		padding = append(padding, 0)
	}
	b = append(padding, b...)

	carry := int8(0)
	for i := resLength - 1; i > -1; i-- {

		t := a[i] + b[i] + carry
		if t > 2 {
			result[i] = t - 5
			carry = 1
			continue
		}
		if t < -2 {
			result[i] = 5 + t
			carry = -1
			continue
		}
		result[i] = t
		carry = 0
	}
	if carry != 0 {
		return append([]int8{carry}, result...)
	}
	return result
}

var SNAFUToDec map[byte]int8 = map[byte]int8{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var DecToSNAFU map[int8]byte = map[int8]byte{
	2:  '2',
	1:  '1',
	0:  '0',
	-1: '-',
	-2: '=',
}

func parseSNAFU(s string) []int8 {
	number := make([]int8, len(s))
	for i := range s {
		number[i] = SNAFUToDec[s[i]]
	}
	return number
}

func String(a []int8) string {
	var b strings.Builder
	for _, c := range a {
		fmt.Fprintf(&b, "%c", DecToSNAFU[c])
	}
	return b.String()
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := make([]int8, 0)

	for scanner.Scan() {
		s := scanner.Text()
		n := parseSNAFU(s)
		sum = Add(sum, n)
	}

	fmt.Printf("sum: %v\n", String(sum))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
