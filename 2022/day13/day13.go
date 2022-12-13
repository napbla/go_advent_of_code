package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type LexemeType int

const (
	OPEN_BRACKET LexemeType = iota
	CLOSE_BRACKET
	NUMBER
)

func (l LexemeType) String() string {
	switch l {
	case OPEN_BRACKET:
		return "OPEN_BRACKET"
	case CLOSE_BRACKET:
		return "CLOSE_BRACKET"
	default:
		return "NUMBER"
	}
}

type Lexeme struct {
	Type LexemeType
	// Begin int
	// End   int
	Value int
}

func lexer(s string) []Lexeme {
	count := 0
	result := make([]Lexeme, 0)
	for count < len(s) {
		switch s[count] {
		case '[':
			{
				result = append(result, Lexeme{
					Type: OPEN_BRACKET,
					// Begin: count,
					// End:   count,
					Value: int(s[count]),
				})
			}

		case ']':
			{
				result = append(result, Lexeme{
					Type: CLOSE_BRACKET,
					// Begin: count,
					// End:   count,
					Value: int(s[count]),
				})
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			{
				start := count
				for ; count < len(s) && s[count] >= '0' && s[count] <= '9'; count++ {
				}
				number, err := strconv.Atoi(s[start:count])
				if err != nil {
					log.Fatalf("error parse number from %d to %d:%v", start, count, err)
				}
				count-- //go back one character
				result = append(result, Lexeme{
					Type: NUMBER,
					// Begin: start,
					// End:   count,
					Value: number,
				})
			}
		case ' ':
			{
			}
		case ',':
			{
			}
		}
		count++
	}
	return result
}

func printDebug(l []Lexeme) {
	for i := 0; i < len(l); i++ {
		if l[i].Type == NUMBER {
			// fmt.Printf("Type:%v\tBegin:%v\tEnd:%v\tValue:%v\n", l[i].Type, l[i].Begin, l[i].End, l[i].Value)
			fmt.Printf("Type:%v\tValue:%v\n", l[i].Type, l[i].Value)
		} else {
			// fmt.Printf("Type:%v\tBegin:%v\tEnd:%v\n", l[i].Type, l[i].Begin, l[i].End)
			fmt.Printf("Type:%v\n", l[i].Type)
		}
	}
}

type ASTType int

const (
	LIST ASTType = iota
	INT
)

type ASTNode struct {
	Parent   *ASTNode
	Children []*ASTNode
	Type     ASTType
	Value    int
}

func parse(l []Lexeme) *ASTNode {
	count := 0
	root := &ASTNode{
		Parent:   nil,
		Children: make([]*ASTNode, 0),
	}
	current := root

	for count < len(l) {
		switch l[count].Type {
		case OPEN_BRACKET:
			{
				newListNode := &ASTNode{
					Parent:   current,
					Children: make([]*ASTNode, 0),
					Type:     LIST,
				}
				current.Children = append(current.Children, newListNode)
				current = newListNode
			}

		case CLOSE_BRACKET:
			{
				current = current.Parent
			}
		case NUMBER:
			{
				newNumberNode := &ASTNode{
					Parent: current,
					Type:   INT,
					Value:  l[count].Value,
				}
				current.Children = append(current.Children, newNumberNode)
			}
		}
		count++
	}

	return root
}

func printDebugAST(ast *ASTNode, prefix string) {
	//print
	switch ast.Type {
	case INT:
		{
			fmt.Printf("%sINT=%d\n", prefix, ast.Value)
		}
	case LIST:
		{
			fmt.Printf("%sLIST\n", prefix)
			for i := 0; i < len(ast.Children); i++ {
				printDebugAST(ast.Children[i], prefix+"__")
			}
		}
	}
}

func compare(l *ASTNode, r *ASTNode) int { //1 -> l < r
	if (l.Type == r.Type) && (r.Type == INT) {
		// fmt.Printf("compare number - number : %d - %d\n", l.Value, r.Value)
		if l.Value < r.Value {
			return 1
		}
		if l.Value > r.Value {
			return -1
		}
		return 0
	}

	if (l.Type == r.Type) && (r.Type == LIST) {
		// fmt.Printf("compare list - list\n")
		min := len(l.Children)
		if len(r.Children) < min {
			min = len(r.Children)
		}
		for i := 0; i < min; i++ {
			// fmt.Printf("compare child %d\n", i)

			temp := compare(l.Children[i], r.Children[i])
			if temp != 0 {
				return temp
			}
		}
		return compare(&ASTNode{
			Type:  INT,
			Value: len(l.Children)},
			&ASTNode{
				Type:  INT,
				Value: len(r.Children),
			})
	}

	if l.Type != r.Type {
		//convert
		if l.Type == INT {
			return compare(&ASTNode{
				Type: LIST,
				Children: []*ASTNode{
					{
						Type:  INT,
						Value: l.Value,
					},
				},
			}, r)
		}

		if r.Type == INT {
			return compare(l, &ASTNode{
				Type: LIST,
				Children: []*ASTNode{
					{
						Type:  INT,
						Value: r.Value,
					},
				},
			})
		}
	}
	return 0
}

func compareString(a string, b string) int {
	lexedA := lexer(a)
	parsedA := parse(lexedA)
	// printDebugAST(parsedA, "")

	lexedB := lexer(b)
	parsedB := parse(lexedB)
	// printDebugAST(parsedB, "")

	return compare(parsedA, parsedB)
}

func main() {

	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	signals := make([]string, 0)

	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			signals = append(signals, s)
		}
	}

	//part 1
	for i := 0; i < len(signals)/2; i++ {
		if compareString(signals[2*i], signals[2*i+1]) == 1 {
			sum += (i + 1)
		}
	}

	//part 2
	divider1 := "[[2]]"
	divider2 := "[[6]]"
	signals = append(signals, divider1)
	signals = append(signals, divider2)

	sort.Slice(signals, func(i, j int) bool {
		//can cache astTree for performance boost
		return compareString(signals[i], signals[j]) == 1
	})

	divider1Index := -1
	divider2Index := -1

	for i := 0; i < len(signals); i++ {
		if signals[i] == divider1 {
			divider1Index = i + 1
		}
		if signals[i] == divider2 {
			divider2Index = i + 1
		}
	}

	// for i := 0; i < len(signals); i++ {
	// 	fmt.Printf("%s\n", signals[i])
	// }

	fmt.Printf("part1=%d\npart2=%d\n", sum, divider1Index*divider2Index)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
