package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	gridWidth    = 20
	gridHeight   = 20
	obstacleRate = 0.28
	frameDelay   = 45 * time.Millisecond
	maxAttempts  = 1000
)

type Point struct {
	X int
	Y int
}

type node struct {
	point Point
	g     int
	h     int
	f     int
	index int
}

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].f == pq[j].f {
		return pq[i].h < pq[j].h
	}
	return pq[i].f < pq[j].f
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

type Snapshot struct {
	Current Point
	Open    map[Point]bool
	Closed  map[Point]bool
	Step    int
}

type Renderer struct {
	grid   [][]bool
	start  Point
	goal   Point
	width  int
	height int
	out    *bufio.Writer
}

func NewRenderer(grid [][]bool, start Point, goal Point, out *os.File) *Renderer {
	return &Renderer{
		grid:   grid,
		start:  start,
		goal:   goal,
		width:  len(grid[0]),
		height: len(grid),
		out:    bufio.NewWriter(out),
	}
}

func (r *Renderer) Init() {
	fmt.Fprint(r.out, "\033[2J\033[H\033[?25l")
	r.out.Flush()
}

func (r *Renderer) Close() {
	fmt.Fprint(r.out, "\033[?25h")
	r.out.Flush()
}

func (r *Renderer) Draw(snapshot *Snapshot, path map[Point]bool, status string) {
	var b strings.Builder

	b.WriteString("\033[H\033[J")
	b.WriteString("A* pathfinding, 20x20\n")
	b.WriteString("S=start G=goal #=obstacle o=open x=closed @=current *=path\n")
	if status == "" {
		b.WriteString("\n")
	} else {
		b.WriteString(status)
		b.WriteByte('\n')
	}

	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			p := Point{X: x, Y: y}
			b.WriteRune(r.symbolFor(p, snapshot, path))
		}
		b.WriteByte('\n')
	}

	if snapshot != nil {
		b.WriteString(fmt.Sprintf("step=%d open=%d closed=%d\n", snapshot.Step, len(snapshot.Open), len(snapshot.Closed)))
	} else {
		b.WriteString("\n")
	}

	fmt.Fprint(r.out, b.String())
	r.out.Flush()
}

func (r *Renderer) symbolFor(p Point, snapshot *Snapshot, path map[Point]bool) rune {
	if p == r.start {
		return 'S'
	}
	if p == r.goal {
		return 'G'
	}
	if path != nil && path[p] {
		return '*'
	}
	if snapshot != nil && p == snapshot.Current {
		return '@'
	}
	if r.grid[p.Y][p.X] {
		return '#'
	}
	if snapshot != nil {
		if snapshot.Closed[p] {
			return 'x'
		}
		if snapshot.Open[p] {
			return 'o'
		}
	}
	return '.'
}

func main() {
	start := Point{X: 0, Y: 0}
	goal := Point{X: gridWidth - 1, Y: gridHeight - 1}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	grid, attempts := buildSolvableGrid(gridWidth, gridHeight, start, goal, obstacleRate, rng)
	renderer := NewRenderer(grid, start, goal, os.Stdout)
	renderer.Init()
	defer renderer.Close()

	var last Snapshot
	path, found := aStar(grid, start, goal, func(s Snapshot) {
		last = s
		renderer.Draw(&s, nil, fmt.Sprintf("searching... generated map attempt=%d", attempts))
		time.Sleep(frameDelay)
	})

	if found {
		pathSet := make(map[Point]bool, len(path))
		for _, p := range path {
			pathSet[p] = true
		}
		renderer.Draw(&last, pathSet, fmt.Sprintf("path found, length=%d", len(path)-1))
	} else {
		renderer.Draw(&last, nil, "path not found")
	}

	fmt.Println()
}

func buildSolvableGrid(width int, height int, start Point, goal Point, rate float64, rng *rand.Rand) ([][]bool, int) {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		grid := make([][]bool, height)
		for y := 0; y < height; y++ {
			grid[y] = make([]bool, width)
			for x := 0; x < width; x++ {
				p := Point{X: x, Y: y}
				if p == start || p == goal {
					continue
				}
				if rng.Float64() < rate {
					grid[y][x] = true
				}
			}
		}
		if _, found := aStar(grid, start, goal, nil); found {
			return grid, attempt
		}
	}

	panic("failed to generate a solvable map")
}

func aStar(grid [][]bool, start Point, goal Point, onStep func(Snapshot)) ([]Point, bool) {
	openQueue := &priorityQueue{}
	heap.Init(openQueue)
	heap.Push(openQueue, &node{point: start, g: 0, h: heuristic(start, goal), f: heuristic(start, goal)})

	cameFrom := map[Point]Point{}
	gScore := map[Point]int{start: 0}
	openSet := map[Point]bool{start: true}
	closedSet := map[Point]bool{}

	step := 0

	for openQueue.Len() > 0 {
		currentNode := heap.Pop(openQueue).(*node)
		current := currentNode.point

		bestG, ok := gScore[current]
		if !ok || currentNode.g != bestG {
			continue
		}

		delete(openSet, current)
		closedSet[current] = true
		step++

		if onStep != nil {
			onStep(Snapshot{Current: current, Open: openSet, Closed: closedSet, Step: step})
		}

		if current == goal {
			return reconstructPath(cameFrom, current), true
		}

		for _, next := range neighbors(current, len(grid[0]), len(grid)) {
			if grid[next.Y][next.X] || closedSet[next] {
				continue
			}

			tentativeG := bestG + 1
			if oldG, seen := gScore[next]; seen && tentativeG >= oldG {
				continue
			}

			cameFrom[next] = current
			gScore[next] = tentativeG
			h := heuristic(next, goal)
			heap.Push(openQueue, &node{point: next, g: tentativeG, h: h, f: tentativeG + h})
			openSet[next] = true
		}
	}

	return nil, false
}

func heuristic(a Point, b Point) int {
	dx := a.X - b.X
	if dx < 0 {
		dx = -dx
	}
	dy := a.Y - b.Y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func neighbors(p Point, width int, height int) []Point {
	dirs := [4]Point{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}}
	result := make([]Point, 0, 4)

	for _, d := range dirs {
		next := Point{X: p.X + d.X, Y: p.Y + d.Y}
		if next.X < 0 || next.Y < 0 || next.X >= width || next.Y >= height {
			continue
		}
		result = append(result, next)
	}

	return result
}

func reconstructPath(cameFrom map[Point]Point, current Point) []Point {
	path := []Point{current}

	for {
		prev, ok := cameFrom[current]
		if !ok {
			break
		}
		path = append(path, prev)
		current = prev
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
