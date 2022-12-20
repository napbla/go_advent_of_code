package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vec3 struct {
	X int
	Y int
	Z int
}

func (v Vec3) String() string {
	return fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z)
}

func Add(a Vec3, b Vec3) Vec3 {
	return Vec3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isAdjacent(a *Vec3, b *Vec3) bool {
	if a.X == b.X && a.Y == b.Y && (a.Z-b.Z == 1 || a.Z-b.Z == -1) {
		return true
	}
	if a.Y == b.Y && a.Z == b.Z && (a.X-b.X == 1 || a.X-b.X == -1) {
		return true
	}
	if a.X == b.X && a.Z == b.Z && (a.Y-b.Y == 1 || a.Y-b.Y == -1) {
		return true
	}
	return false
}

// 0: air, 1:obsidian, 2:visited in part 2
var cubeFlags map[Vec3]int

func surface(cubes []*Vec3) int {
	hides := make([]int, len(cubes))
	for i, c := range cubes {
		count := 0
		for j, cc := range cubes {
			if i != j {
				if isAdjacent(c, cc) {
					count++
				}
			}
		}
		hides[i] = count
	}

	count := 0
	for _, h := range hides {
		count += (6 - h)
	}
	return count
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cubes := make([]*Vec3, 0)

	scanner := bufio.NewScanner(file)

	cubeFlags = make(map[Vec3]int)

	min := Vec3{0, 0, 0}
	max := Vec3{0, 0, 0}
	first := 0
	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)
		splitted := strings.Split(s, ",")
		x, errx := strconv.Atoi(splitted[0])
		if err != nil {
			log.Fatalf("can not parse int:%v", errx)
		}
		y, erry := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatalf("can not parse int:%v", erry)
		}
		z, errz := strconv.Atoi(splitted[2])
		if err != nil {
			log.Fatalf("can not parse int:%v", errz)
		}
		cube := Vec3{x, y, z}
		cubes = append(cubes, &cube)

		if first == 0 {
			first++
			min = cube
			max = min
		}

		min.X = Min(min.X, x)
		min.Y = Min(min.Y, y)
		min.Z = Min(min.Z, z)

		max.X = Max(max.X, x)
		max.Y = Max(max.Y, y)
		max.Z = Max(max.Z, z)

		cubeFlags[cube] = 1
	}

	surfaceArea := surface(cubes)
	fmt.Printf("part1=%d\n", surfaceArea)

	min = Add(min, Vec3{-1, -1, -1})
	max = Add(max, Vec3{1, 1, 1})

	count := 0
	queue := []*Vec3{&min}

	moves := []Vec3{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, 1}, {0, 0, -1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if cubeFlags[*current] == 0 {
			for _, m := range moves {
				next := Add(*current, m)
				if next.X < min.X || next.Y < min.Y || next.Z < min.Z || next.X > max.X || next.Y > max.Y || next.Z > max.Z {
					continue
				}
				if cubeFlags[next] == 1 {
					count++
				}
				if cubeFlags[next] == 0 {
					queue = append(queue, &next)
				}
			}
			cubeFlags[*current] = 2
		}
	}

	fmt.Printf("part2=%d\n", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
