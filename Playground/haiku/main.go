package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	gridSize     = 20
	obstacleProb = 0.2 // 20% chance of obstacle
)

// Cell represents a cell in the grid
type Cell struct {
	X              int
	Y              int
	IsWall         bool
	Visited        bool
	IsPath         bool
	IsCurrent      bool
	Parent         *Cell
	IsFinalPath    bool
}

// Grid represents the maze/graph
type Grid struct {
	cells [gridSize][gridSize]*Cell
}

// NewGrid creates and initializes a new grid
func NewGrid() *Grid {
	g := &Grid{}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			isWall := rng.Float64() < obstacleProb
			// Ensure start and end are not walls
			if (x == 0 && y == 0) || (x == gridSize-1 && y == gridSize-1) {
				isWall = false
			}
			g.cells[y][x] = &Cell{
				X:      x,
				Y:      y,
				IsWall: isWall,
			}
		}
	}
	return g
}

// GetNeighbors returns unvisited neighbors for DFS
func (g *Grid) GetNeighbors(x, y int) []*Cell {
	var neighbors []*Cell
	// 4-directional movement: up, down, left, right
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < gridSize && ny >= 0 && ny < gridSize {
			cell := g.cells[ny][nx]
			if !cell.IsWall && !cell.Visited {
				neighbors = append(neighbors, cell)
			}
		}
	}
	return neighbors
}

// DFS performs depth-first search with visualization
func (g *Grid) DFS(startX, startY, endX, endY int, render func()) bool {
	stack := []*Cell{g.cells[startY][startX]}
	g.cells[startY][startX].Visited = true
	g.cells[startY][startX].IsCurrent = true

	render() // Initial render

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		// Check if we reached the end
		if current.X == endX && current.Y == endY {
			current.IsPath = true
			// Reconstruct the final path
			g.markFinalPath(current)
			render()
			return true
		}

		neighbors := g.GetNeighbors(current.X, current.Y)
		if len(neighbors) == 0 {
			// Backtrack
			current.IsCurrent = false
			stack = stack[:len(stack)-1]
			render()
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Pick first unvisited neighbor
		next := neighbors[0]
		next.Visited = true
		next.Parent = current
		next.IsCurrent = true
		current.IsCurrent = false
		current.IsPath = true

		stack = append(stack, next)
		render()
		time.Sleep(100 * time.Millisecond)
	}

	return false
}

// markFinalPath marks the final path from end to start
func (g *Grid) markFinalPath(endCell *Cell) {
	current := endCell
	for current != nil {
		current.IsFinalPath = true
		current.Visited = false // Clear visited to show only final path
		current.IsPath = false
		current = current.Parent
	}
}

// Render draws the grid to the terminal
func (g *Grid) Render() {
	// Move cursor to home position (top-left) to overwrite previous frame
	fmt.Print("\033[H")

	fmt.Println("DFS Path Finding Visualization (20x20)")
	fmt.Println("█ = Wall | . = Visited | @ = Current | * = Path | # = Final Path | S = Start | E = End")
	fmt.Println()

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			cell := g.cells[y][x]

			switch {
			case x == 0 && y == 0:
				fmt.Print("S")
			case x == gridSize-1 && y == gridSize-1:
				fmt.Print("E")
			case cell.IsWall:
				fmt.Print("█")
			case cell.IsFinalPath:
				fmt.Print("#")
			case cell.IsCurrent:
				fmt.Print("@")
			case cell.IsPath:
				fmt.Print("*")
			case cell.Visited:
				fmt.Print(".")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Running DFS algorithm...")
}

func main() {
	// Clear screen once at the beginning
	fmt.Print("\033[2J\033[H")

	fmt.Println("Initializing DFS Pathfinding...")
	grid := NewGrid()

	startX, startY := 0, 0
	endX, endY := gridSize-1, gridSize-1

	fmt.Println("Press Enter to start...")
	fmt.Scanln()

	found := grid.DFS(startX, startY, endX, endY, func() {
		grid.Render()
	})

	grid.Render()
	if found {
		fmt.Println("✓ Path found!")
	} else {
		fmt.Println("✗ No path exists!")
	}
	fmt.Println("\nAlgorithm completed.")
}
