package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readBoxText(s string) []byte {
	res := make([]byte, 0)
	for i := 1; i < len(s); i += 4 {
		res = append(res, s[i])
	}
	return res
}

func reform(mat [][]byte, label []byte) map[byte][]byte {
	fmt.Printf("%v\n", mat)
	fmt.Printf("%v\n", label)
	res := make(map[byte][]byte)
	for i := 0; i < len(label); i++ {
		res2 := make([]byte, 0)
		for j := len(mat) - 1; j > -1; j-- {
			if mat[j][i] == ' ' {
				break
			}
			res2 = append(res2, mat[j][i])
		}
		res[label[i]] = res2
	}
	return res
}

func parseNumberFromString(s string, start int) (byte, int) {
	stop := 0
	val := byte(0)
	for i := start; i < len(s); i++ {
		if s[i] == ' ' {
			stop = i + 1
			break
		}
		val = val*10 + s[i] - 48
		// fmt.Printf("val=%d, s[i]=%c \n", val, s[i])
	}
	// fmt.Printf("val=%d stop=%d\n", val, stop)
	return val, stop
}

func parseMove(s string) (byte, byte, byte) {
	num, end := parseNumberFromString(s, 5)
	num2, end2 := parseNumberFromString(s, end+5)
	num3, _ := parseNumberFromString(s, end2+3)
	return num, num2, num3
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matrix := make([][]byte, 0)
	label := make([]byte, 0)

	var mm map[byte][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if len(s) > 0 {
			if strings.Contains(s, "[") {
				//readblock
				matrix = append(matrix, readBoxText(s))
				continue
			}
			if s[0] == ' ' {
				//readlabel
				for i := 1; i < len(s); i += 4 {
					label = append(label, s[i]-48)
				}
				mm = reform(matrix, label)
				fmt.Printf("%v\n", mm)
				continue
			}
			if s[0] == 'm' {
				a, b, c := parseMove(s)

				// if you need part 1, just uncomment the 9000 and comment 9001.
				// for part 2 do the reverse

				//9000
				// for i := byte(0); i < a; i++ {
				// 	mm[c] = append(mm[c], mm[b][len(mm[b])-int(i)-1])
				// }
				// mm[b] = mm[b][:len(mm[b])-int(a)]

				//9001
				for i := a; i > 0; i-- {
					mm[c] = append(mm[c], mm[b][len(mm[b])-1-int(i-1)])
				}
				mm[b] = mm[b][:len(mm[b])-int(a)]
				continue
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Printf("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(label); i++ {
		stack := mm[label[i]]
		fmt.Printf("%c", stack[len(stack)-1])
	}
}
