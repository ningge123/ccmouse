package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, y int
}

var dirs = [4]point{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.y + r.y}
}

func (p point) at(grid [][]int) (int, bool)  {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.y < 0 || p.y >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.y], true
}

func walk(maze [][]int, start, end point) [][]int {

	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	 // 队列
	 queue := []point{start}

	for len(queue) > 0 {
		// 探索队列头部
		cur := queue[0]

		if cur.i == 3 && cur.y == 4 {
			fmt.Println("asd")
		}

		queue = queue[1:]

		fmt.Println(cur, end)

		if cur == end {
			break
		}
		// 从上下左右四个方向探索
		for _, dir := range dirs {
			next := cur.add(dir)

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
			curStep, _ := cur.at(steps)
			steps[next.i][next.y] = curStep + 1

			queue = append(queue, next)
		}
	}

	return steps
}

func main()  {
	// 读取文件
	maze := readMaze("maze.in")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, step := range steps {
		for _, i2 := range step {
			fmt.Printf("%3d", i2)
		}

		fmt.Println()
	}
}
