package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type VarType int

const (
	EXPRESSION VarType = iota
	VALUE
)

func Strings(t VarType) string {
	if t == EXPRESSION {
		return "EXPRESSION"
	}
	return "VALUE"
}

type Node struct {
	Name      string
	Type      VarType
	Value     int
	Operator  byte
	Evaluated bool
	Children  [2]*Node
}

func eval(n *Node) int {
	// fmt.Printf("eval node = %v\n", n)

	if n.Type == VALUE {
		return n.Value
	}

	//Expression
	if n.Evaluated {
		return n.Value
	}
	if n.Operator == '+' {
		n.Value = eval(n.Children[0]) + eval(n.Children[1])
	}
	if n.Operator == '-' {
		n.Value = eval(n.Children[0]) - eval(n.Children[1])
	}
	if n.Operator == '*' {
		n.Value = eval(n.Children[0]) * eval(n.Children[1])
	}
	if n.Operator == '/' {
		n.Value = eval(n.Children[0]) / eval(n.Children[1])
	}
	n.Evaluated = true

	return n.Value
}

func findAndClear(n *Node, name string) bool {
	if n.Type == VALUE {
		if n.Name == name {
			n.Evaluated = false
			return true
		} else {
			return false
		}
	}
	found := findAndClear(n.Children[0], name) || findAndClear(n.Children[1], name)
	n.Evaluated = !found
	return found
}

func clearAll(n *Node) {
	if n.Type == VALUE {
		return
	}
	n.Evaluated = false
	clearAll(n.Children[0])
	clearAll(n.Children[1])
}

func reverse(n *Node, result int) {

	if n.Type == VALUE && !n.Evaluated {
		fmt.Printf("part2=%d\n", result)
		return
	}

	next := n.Children[0]
	old := n.Children[1]
	if next.Evaluated {
		next = n.Children[1]
		old = n.Children[0]
	}
	if n.Operator == '+' {
		reverse(next, result-old.Value)
		return
	}
	if n.Operator == '-' {
		if next == n.Children[0] {
			reverse(next, result+old.Value)
		} else {
			reverse(next, old.Value-result)
		}
		return
	}
	if n.Operator == '*' {
		reverse(next, result/old.Value)
		return
	}
	if n.Operator == '/' {
		if next == n.Children[0] {
			reverse(next, result*old.Value)
		} else {
			reverse(next, old.Value/result)
		}
		return
	}
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodes := make(map[string]*Node)

	ops := []string{"+", "-", "*", "/"}

	for scanner.Scan() {
		s := scanner.Text()

		splitted := strings.Split(s, ": ")
		left := splitted[0]
		right := splitted[1]
		if right[0] >= '0' && right[0] <= '9' {
			x, err := strconv.Atoi(right)
			if err != nil {
				log.Fatalf("can not parse int:%v", err)
			}
			node := nodes[left]
			if node == nil {
				nodes[left] = &Node{
					Name:      left,
					Type:      VALUE,
					Value:     x,
					Evaluated: true,
				}
			} else {
				node.Name = left
				node.Type = VALUE
				node.Value = x
				node.Evaluated = true
			}

		} else {
			for _, o := range ops {
				i := strings.Index(right, o)
				if i > -1 {
					child1 := right[:i-1]
					child2 := right[i+2:]
					// fmt.Printf("child1=%s,child2=%s\n", child1, child2)

					c1node := nodes[child1]
					if c1node == nil {
						c1node = &Node{}
						nodes[child1] = c1node
					}

					c2node := nodes[child2]
					if c2node == nil {
						c2node = &Node{}
						nodes[child2] = c2node
					}

					node := nodes[left]
					if node == nil {
						nodes[left] = &Node{
							Name:      left,
							Type:      EXPRESSION,
							Evaluated: false,
							Children:  [2]*Node{c1node, c2node},
							Operator:  right[i],
						}
					} else {
						node.Name = left
						node.Type = EXPRESSION
						node.Evaluated = false
						node.Children = [2]*Node{c1node, c2node}
						node.Operator = right[i]
					}
					break
				}
			}
		}
	}

	// for k, v := range nodes {
	// 	fmt.Printf("k=%s v=%v\n", k, v)
	// }

	root := eval(nodes["root"])
	fmt.Printf("part1=%d\n", root)

	//part 2
	left := eval(nodes["root"].Children[0])
	right := eval(nodes["root"].Children[1])
	// fmt.Printf("left=%d, right=%d\n", left, right)

	findAndClear(nodes["root"], "humn")
	if !nodes["root"].Children[0].Evaluated {
		reverse(nodes["root"].Children[0], right)
	} else {
		reverse(nodes["root"].Children[1], left)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
