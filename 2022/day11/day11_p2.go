package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation func(int) int

type Monkey struct {
	Items     []int
	Operation Operation
	Test      int //divisible by Test
	Throw     map[bool]int
}

func parseStartingItems(s string) []int {
	i := 0
	for ; s[i] != ':' && i < len(s); i++ {
	}
	numbers := strings.Split(s[i+2:], ", ")
	result := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		n, err := strconv.Atoi(numbers[i])
		if err != nil {
			log.Fatalf("can not parse number in %s: %v", s, err)
		}
		result[i] = n
	}
	return result
}

func parseOperation(s string) Operation {
	i := 0
	for ; s[i] != '=' && i < len(s); i++ {
	}
	//find operator pos
	j := i + 1
	for ; j < len(s) && (s[j] != '+' && s[j] != '-' && s[j] != '*' && s[j] != '/'); j++ {
	}
	op := s[j]
	right := s[j+2:]
	if right != "old" {
		n, err := strconv.Atoi(right)
		if err != nil {
			log.Fatalf("can not parse right value of operation:%v", err)
		}
		switch op {
		case '+':
			{
				return func(i int) int {
					return i + n
				}
			}

		case '-':
			{
				return func(i int) int {
					return i - n
				}
			}
		case '*':
			{
				return func(i int) int {
					return i * n
				}
			}
		case '/':
			{
				return func(i int) int {
					return i / n
				}
			}
		}
	}

	switch op {
	case '+':
		{
			return func(i int) int {
				return i + i
			}
		}

	case '-':
		{
			return func(i int) int {
				return 0
			}
		}
	case '*':
		{
			return func(i int) int {
				return i * i
			}
		}
	case '/':
		{
			return func(i int) int {
				return 1
			}
		}
	}
	return nil
}

func parseTest(s string) int {
	i := 0
	for ; s[i] != 'y' && i < len(s); i++ {
	}
	n, err := strconv.Atoi(s[i+2:])
	if err != nil {
		log.Fatalf("can not parse test value:%v", err)
	}
	return n
}

func parseTrueFalse(s string) int {
	i := len(s) - 1
	for ; s[i] != ' ' && i > -1; i-- {
	}
	n, err := strconv.Atoi(s[i+1:])
	if err != nil {
		log.Fatalf("can not parse true false value:%v", err)
	}
	return n
}

// func (m *Monkey) PrintDebug() {
// 	fmt.Printf("Starting items: %v\n", m.Items)
// 	fmt.Printf("Operation(8): %v\n%v\n", m.Operation(8), m.Operation)
// 	fmt.Printf("Divisible by: %d\n", m.Test)
// 	fmt.Printf(" . True throw to: %d\n", m.Throw[true])
// 	fmt.Printf(" . False throw to: %d\n", m.Throw[false])
// }

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	monkeys := make([]*Monkey, 0)

	monkeyInfos := make([]string, 5)
	count := 0

	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			// fmt.Printf("%v", monkeyInfos)
			monkeys = append(monkeys, &Monkey{
				Items:     parseStartingItems(monkeyInfos[0]),
				Operation: parseOperation(monkeyInfos[1]),
				Test:      parseTest(monkeyInfos[2]),
				Throw: map[bool]int{
					true:  parseTrueFalse(monkeyInfos[3]),
					false: parseTrueFalse(monkeyInfos[4]),
				},
			})
			count = 0
		} else {
			if s[:6] != "Monkey" {
				monkeyInfos[count] = s
				count++
			}
		}
	}

	if count == 5 {
		monkeys = append(monkeys, &Monkey{
			Items:     parseStartingItems(monkeyInfos[0]),
			Operation: parseOperation(monkeyInfos[1]),
			Test:      parseTest(monkeyInfos[2]),
			Throw: map[bool]int{
				true:  parseTrueFalse(monkeyInfos[3]),
				false: parseTrueFalse(monkeyInfos[4]),
			},
		})
	}

	// for _, m := range monkeys {
	// 	m.PrintDebug()
	// }

	times := make([]int, len(monkeys))

	roundNumber := 10000

	lcm := 1
	for _, m := range monkeys {
		lcm *= m.Test
	}

	for r := 0; r < roundNumber; r++ {
		for i, m := range monkeys {
			times[i] += len(m.Items)
			for _, n := range m.Items {
				worryLevel := m.Operation(n)
				next := m.Throw[worryLevel%m.Test == 0]
				monkeys[next].Items = append(monkeys[next].Items,
					worryLevel%lcm)
			}
			m.Items = []int{}
		}
	}

	// for i, t := range times {
	// 	fmt.Printf("Monkey %d plays %d times\n", i, t)
	// }

	sort.Slice(times, func(i int, j int) bool {
		return times[i] > times[j]
	})

	fmt.Printf("part2=%v\n", times[0]*times[1])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
