package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	cols       = 20
	rows       = 20
	wallChance = 0.28
	stepDelay  = 50 * time.Millisecond
	pathDelay  = 30 * time.Millisecond
)

// Cell types for the grid.
type Cell byte

const (
	CellEmpty Cell = iota
	CellWall
	CellVisited
	CellFrontier
	CellPath
)

// Pos is a 2D coordinate on the grid.
type Pos struct{ X, Y int }

// Node is an A* search node.
type Node struct {
	Pos    Pos
	G, F   float64
	Parent *Node
	idx    int
}

// PQ implements heap.Interface for A* open set.
type PQ []*Node

func (q PQ) Len() int            { return len(q) }
func (q PQ) Less(i, j int) bool  { return q[i].F < q[j].F }
func (q PQ) Swap(i, j int)       { q[i], q[j] = q[j], q[i]; q[i].idx = i; q[j].idx = j }
func (q *PQ) Push(x interface{}) { n := x.(*Node); n.idx = len(*q); *q = append(*q, n) }
func (q *PQ) Pop() interface{} {
	old := *q
	n := old[len(old)-1]
	old[len(old)-1] = nil
	*q = old[:len(old)-1]
	return n
}

var dirs = [4]Pos{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func manhattan(a, b Pos) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

func main() {
	// Restore cursor on Ctrl+C.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Print("\033[?25h\n")
		os.Exit(0)
	}()
	defer fmt.Print("\033[?25h")

	start := Pos{0, 0}
	end := Pos{cols - 1, rows - 1}

	grid := generateGrid(start, end)

	// Hide cursor, clear screen.
	fmt.Print("\033[?25l\033[2J")

	render(grid, start, end, 0, 0, "Searching...")

	// --- A* ---
	open := &PQ{}
	heap.Init(open)
	startNode := &Node{Pos: start, G: 0, F: manhattan(start, end)}
	heap.Push(open, startNode)

	best := map[Pos]*Node{start: startNode}
	closed := map[Pos]bool{}

	var final *Node
	steps := 0

	for open.Len() > 0 {
		cur := heap.Pop(open).(*Node)

		if cur.Pos == end {
			final = cur
			break
		}
		if closed[cur.Pos] {
			continue
		}
		closed[cur.Pos] = true
		steps++

		if cur.Pos != start {
			grid[cur.Pos.Y][cur.Pos.X] = CellVisited
		}

		for _, d := range dirs {
			np := Pos{cur.Pos.X + d.X, cur.Pos.Y + d.Y}
			if np.X < 0 || np.X >= cols || np.Y < 0 || np.Y >= rows {
				continue
			}
			if grid[np.Y][np.X] == CellWall || closed[np] {
				continue
			}
			g := cur.G + 1
			if existing, ok := best[np]; ok && g >= existing.G {
				continue
			}
			node := &Node{Pos: np, G: g, F: g + manhattan(np, end), Parent: cur}
			best[np] = node
			heap.Push(open, node)

			if np != end && grid[np.Y][np.X] != CellVisited {
				grid[np.Y][np.X] = CellFrontier
			}
		}

		render(grid, start, end, steps, 0, "Searching...")
		time.Sleep(stepDelay)
	}

	if final == nil {
		render(grid, start, end, steps, 0, "\033[31mNo path found! Re-run to try a new grid.\033[0m")
		fmt.Println()
		return
	}

	// Collect path nodes.
	var path []Pos
	for n := final; n != nil; n = n.Parent {
		path = append(path, n.Pos)
	}

	// Animate path drawing from start to end.
	for i := len(path) - 1; i >= 0; i-- {
		p := path[i]
		if p != start && p != end {
			grid[p.Y][p.X] = CellPath
		}
		render(grid, start, end, steps, len(path), "Drawing path...")
		time.Sleep(pathDelay)
	}

	render(grid, start, end, steps, len(path), "\033[1;32mDone! Path found.\033[0m")
	fmt.Println()
}

// generateGrid creates a grid with random walls, ensuring start/end are clear.
func generateGrid(start, end Pos) [][]Cell {
	grid := make([][]Cell, rows)
	for y := range grid {
		grid[y] = make([]Cell, cols)
		for x := range grid[y] {
			if rand.Float64() < wallChance {
				grid[y][x] = CellWall
			}
		}
	}
	// Clear start, end, and their neighbors.
	for _, p := range []Pos{start, end} {
		grid[p.Y][p.X] = CellEmpty
		for _, d := range dirs {
			nx, ny := p.X+d.X, p.Y+d.Y
			if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
				grid[ny][nx] = CellEmpty
			}
		}
	}
	return grid
}

// render draws the entire grid to the terminal using a buffer for flicker-free output.
func render(grid [][]Cell, start, end Pos, steps, pathLen int, status string) {
	var buf strings.Builder
	buf.Grow(2048)

	buf.WriteString("\033[H") // cursor home

	// Top border.
	buf.WriteString("  ┌")
	for x := 0; x < cols; x++ {
		buf.WriteString("──")
	}
	buf.WriteString("─┐\n")

	// Grid rows.
	for y := 0; y < rows; y++ {
		buf.WriteString(fmt.Sprintf("%2d│", y))
		for x := 0; x < cols; x++ {
			p := Pos{x, y}
			switch {
			case p == start:
				buf.WriteString(" \033[1;32mS\033[0m")
			case p == end:
				buf.WriteString(" \033[1;31mE\033[0m")
			default:
				switch grid[y][x] {
				case CellWall:
					buf.WriteString("\033[47m  \033[0m")
				case CellVisited:
					buf.WriteString(" \033[36m·\033[0m")
				case CellFrontier:
					buf.WriteString(" \033[33m◦\033[0m")
				case CellPath:
					buf.WriteString(" \033[1;32m●\033[0m")
				default:
					buf.WriteString("  ")
				}
			}
		}
		buf.WriteString(" │\n")
	}

	// Bottom border.
	buf.WriteString("  └")
	for x := 0; x < cols; x++ {
		buf.WriteString("──")
	}
	buf.WriteString("─┘\n")

	// Stats & legend.
	buf.WriteString(fmt.Sprintf("  Status: %-40s Steps: %-4d Path: %d\n", status, steps, pathLen))
	buf.WriteString("  \033[1;32mS\033[0m Start  \033[1;31mE\033[0m End  \033[47m  \033[0m Wall  \033[36m·\033[0m Visited  \033[33m◦\033[0m Frontier  \033[1;32m●\033[0m Path\n")

	fmt.Print(buf.String())
}
