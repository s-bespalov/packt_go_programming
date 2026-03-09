package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows = 20
	cols = 20
)

type Point struct {
	r, c int
}

type Cell rune

const (
	Empty   Cell = ' '
	Wall    Cell = '#'
	Start   Cell = 'S'
	End     Cell = 'E'
	Visited Cell = '·'
	InQueue Cell = '+'
	Path    Cell = '*'
)

type Grid [rows][cols]Cell

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	grid := generateGrid()
	start := Point{0, 0}
	end := Point{rows - 1, cols - 1}

	// Hide cursor, clear screen once
	fmt.Print("\033[?25l")
	fmt.Print("\033[2J\033[H")

	printHeader()
	printGrid(&grid)
	printLegend()
	fmt.Println("\n  Searching...")

	time.Sleep(400 * time.Millisecond)

	path := bfs(&grid, start, end)

	// Mark final path
	if path != nil {
		for _, p := range path {
			if grid[p.r][p.c] != Start && grid[p.r][p.c] != End {
				grid[p.r][p.c] = Path
			}
		}
	}

	// Final render
	fmt.Print("\033[H")
	printHeader()
	printGrid(&grid)
	printLegend()

	if path != nil {
		fmt.Printf("\n  \033[92mPath found! Length: %d steps\033[0m\n", len(path)-1)
	} else {
		fmt.Printf("\n  \033[91mNo path found!\033[0m\n")
	}

	// Restore cursor
	fmt.Print("\033[?25h")
}

func generateGrid() Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rng.Float32() < 0.28 {
				grid[r][c] = Wall
			} else {
				grid[r][c] = Empty
			}
		}
	}

	grid[0][0] = Start
	grid[rows-1][cols-1] = End

	// Clear 3x3 area around start and end for better reachability
	clearAround := func(r, c int) {
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == Wall {
					grid[nr][nc] = Empty
				}
			}
		}
	}
	clearAround(0, 0)
	clearAround(rows-1, cols-1)

	return grid
}

func bfs(grid *Grid, start, end Point) []Point {
	var visited [rows][cols]bool
	var parent [rows][cols]Point

	visited[start.r][start.c] = true
	queue := []Point{start}

	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return reconstructPath(parent, start, end)
		}

		if grid[curr.r][curr.c] != Start {
			grid[curr.r][curr.c] = Visited
		}

		for _, d := range dirs {
			next := Point{curr.r + d.r, curr.c + d.c}
			if next.r < 0 || next.r >= rows || next.c < 0 || next.c >= cols {
				continue
			}
			if visited[next.r][next.c] || grid[next.r][next.c] == Wall {
				continue
			}
			visited[next.r][next.c] = true
			parent[next.r][next.c] = curr
			queue = append(queue, next)
			if grid[next.r][next.c] != End {
				grid[next.r][next.c] = InQueue
			}
		}

		// Move cursor to home and redraw in place
		fmt.Print("\033[H")
		printHeader()
		printGrid(grid)
		printLegend()
		fmt.Printf("\n  Exploring... queue size: %d   \n", len(queue))

		time.Sleep(20 * time.Millisecond)
	}

	return nil
}

func reconstructPath(parent [rows][cols]Point, start, end Point) []Point {
	var path []Point
	curr := end
	for curr != start {
		path = append(path, curr)
		curr = parent[curr.r][curr.c]
	}
	path = append(path, start)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func printHeader() {
	fmt.Println("  \033[1;96mBFS Pathfinding — 20×20 Grid\033[0m")
	fmt.Println()
}

func printGrid(grid *Grid) {
	fmt.Print("  ┌")
	for c := 0; c < cols; c++ {
		fmt.Print("─")
	}
	fmt.Println("┐")

	for r := 0; r < rows; r++ {
		fmt.Print("  │")
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case Wall:
				fmt.Print("\033[90m█\033[0m")
			case Empty:
				fmt.Print(" ")
			case Start:
				fmt.Print("\033[92;1mS\033[0m")
			case End:
				fmt.Print("\033[91;1mE\033[0m")
			case Visited:
				fmt.Print("\033[36m·\033[0m")
			case InQueue:
				fmt.Print("\033[34m+\033[0m")
			case Path:
				fmt.Print("\033[93;1m*\033[0m")
			}
		}
		fmt.Println("│")
	}

	fmt.Print("  └")
	for c := 0; c < cols; c++ {
		fmt.Print("─")
	}
	fmt.Println("┘")
}

func printLegend() {
	fmt.Println()
	fmt.Println("  \033[90m█\033[0m Wall  " +
		"\033[92;1mS\033[0m Start  " +
		"\033[91;1mE\033[0m End  " +
		"\033[36m·\033[0m Visited  " +
		"\033[34m+\033[0m In Queue  " +
		"\033[93;1m*\033[0m Path")
}
