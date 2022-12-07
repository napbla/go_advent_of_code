package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name     string
	Parent   *File
	Children []*File
	Size     int
}

func calculateAndUpdateSize(f *File) int {
	if f.Size == 0 {
		for i := 0; i < len(f.Children); i++ {
			f.Size += calculateAndUpdateSize(f.Children[i])
		}
	}
	return f.Size
}

func printDir(f *File, prefix string) {
	fmt.Printf("%s- %s (size=%d)\n", prefix, f.Name, f.Size)
	for i := 0; i < len(f.Children); i++ {
		printDir(f.Children[i], prefix+" ")
	}
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	root := &File{
		Name:     "/",
		Parent:   nil,
		Children: make([]*File, 0),
		Size:     0,
	}
	currentDirectory := root

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		//is command
		if s[0] == '$' {
			command := s[2:4]
			// fmt.Printf("command=%s\n", command)

			switch command {
			case "cd":
				//parse dir
				dir := s[5:]
				// fmt.Printf("name=%s\n", dir)

				switch dir {
				case "/":
					currentDirectory = root
				case "..":
					currentDirectory = currentDirectory.Parent
				default:
					for i := 0; i < len(currentDirectory.Children); i++ {
						if currentDirectory.Children[i].Name == dir {
							currentDirectory = currentDirectory.Children[i]
							break
						}
					}
				}
			case "ls":
				continue
			}
		} else {
			splitted := strings.Split(s, " ")
			name := splitted[1]
			// fmt.Printf("size/type=%s, name=%s\n", splitted[0], splitted[1])
			if splitted[0] == "dir" {

				found := false
				for i := 0; i < len(currentDirectory.Children); i++ {
					if currentDirectory.Children[i].Name == name {
						found = true
						break
					}
				}
				if !found {
					currentDirectory.Children = append(currentDirectory.Children,
						&File{
							Name:     name,
							Parent:   currentDirectory,
							Children: make([]*File, 0),
							Size:     0,
						})
				}
			} else {
				size, err := strconv.Atoi(splitted[0])
				if err != nil {
					log.Fatalf("error parse file size:%v", err)
				}

				index := -1
				for i := 0; i < len(currentDirectory.Children); i++ {
					if currentDirectory.Children[i].Name == name {
						index = i
						break
					}
				}
				if index > 1 {
					currentDirectory.Children[index].Size = size
				} else {
					currentDirectory.Children = append(currentDirectory.Children,
						&File{
							Name:     name,
							Parent:   currentDirectory,
							Children: make([]*File, 0),
							Size:     size,
						})
				}

			}
		}
	}

	//calculate size
	calculateAndUpdateSize(root)

	printDir(root, " ")

	//part 1
	sum := 0
	stack := []*File{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Size < 100000 {
			sum += current.Size
		}
		for i := 0; i < len(current.Children); i++ {
			if len(current.Children[i].Children) > 0 {
				stack = append(stack, current.Children[i])
			}
		}
	}

	//part 2
	UnusedSpace := 70000000 - root.Size
	min := 70000000
	saved := 0

	stack = []*File{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if UnusedSpace+current.Size > 30000000 &&
			UnusedSpace+current.Size-30000000 < min {
			saved = current.Size
			min = UnusedSpace + current.Size - 30000000
		}

		for i := 0; i < len(current.Children); i++ {
			if len(current.Children[i].Children) > 0 {
				stack = append(stack, current.Children[i])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\npart1=%d, part2=%d\n", sum, saved)
}
