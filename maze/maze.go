package main

import (
	"fmt"
	"os"
)

func readFile(filename string) [][]int {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d\n", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			if j == col-1 {
				fmt.Fscanf(file, "%d\n", &maze[i][j])
			} else {
				fmt.Fscanf(file, "%d", &maze[i][j])
			}
		}
	}

	return maze
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		current := Q[0]
		Q = Q[1:]

		if current == end {
			break
		}

		for _, dir := range dirs {
			next := current.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			currentSteps, _ := current.at(steps)
			steps[next.i][next.j] = currentSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readFile("maze/maze.txt")

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}
