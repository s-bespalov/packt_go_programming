package main

import (
	"fmt"
	"strings"
)

const (
	rows            = 10
	cols            = 15
	canvasWidth     = cols*4 + 5
	startRow        = 0
	startCol        = 0
	maxQueueView    = 8
	maxVisitedPrint = 20
	maxStepsToShow  = 6
)

var (
	startCell = cell{row: startRow, col: startCol}
	goalCell  = cell{row: rows - 1, col: cols - 1}
)

var directions = []cell{
	{row: 0, col: 1},
	{row: 1, col: 0},
	{row: 0, col: -1},
	{row: -1, col: 0},
}

type cell struct {
	row int
	col int
}

func (c cell) key() string {
	return fmt.Sprintf("%02d,%02d", c.row, c.col)
}

func (c cell) String() string {
	return c.key()
}

type graph struct {
	obstacles map[string]bool
}

type searchStep struct {
	number  int
	current cell
	queue   []cell
	visited []cell
}

type searchResult struct {
	steps []searchStep
	path  []cell
	found bool
}

func main() {
	g := sampleGraph()

	fmt.Println("BFS Pathfinding / ASCII Grid Demo")
	fmt.Println(strings.Repeat("=", canvasWidth))

	result := g.search(startCell, goalCell)

	if len(result.steps) == 0 {
		fmt.Println("Поиск не стартовал.")
		return
	}

	stepCount := len(result.steps)
	for i, step := range result.steps {
		if i == maxStepsToShow && stepCount > maxStepsToShow+1 {
			fmt.Printf("[...> пропущено %d промежуточных шагов ...]\n", stepCount-maxStepsToShow-1)
		}
		if i < maxStepsToShow || i == stepCount-1 {
			fmt.Printf("\nStep %d: visiting %s\n", step.number, step.current)
			fmt.Printf("Queue: %s\n", formatCellList(step.queue, " -> ", maxQueueView))
			fmt.Printf("Visited: %s\n", formatCellList(step.visited, ", ", maxVisitedPrint))
			pathToHighlight := []cell(nil)
			showPath := false
			if i == stepCount-1 && result.found {
				pathToHighlight = result.path
				showPath = true
			}
			fmt.Println(g.renderGrid(step, pathToHighlight, showPath))
		}
	}

	if result.found {
		fmt.Println("Найден путь:")
		fmt.Println(formatCellList(result.path, " -> ", 0))
	} else {
		fmt.Printf("Маршрут из %s в %s не найден.\n", startCell, goalCell)
	}
}

func sampleGraph() graph {
	obstacles := []cell{
		{1, 3}, {1, 4}, {1, 8}, {1, 9},
		{2, 5}, {2, 6}, {2, 7},
		{3, 2}, {3, 3}, {3, 10}, {3, 11},
		{4, 8}, {4, 9}, {4, 10},
		{5, 1}, {5, 2}, {5, 5}, {5, 6}, {5, 7}, {5, 12},
		{6, 4}, {6, 5}, {6, 6}, {6, 13},
		{7, 8}, {7, 9}, {7, 10},
		{8, 11}, {8, 12}, {8, 13},
		{9, 3}, {9, 7},
	}

	obstacleSet := make(map[string]bool, len(obstacles))
	for _, o := range obstacles {
		if o == startCell || o == goalCell {
			continue
		}
		obstacleSet[o.key()] = true
	}

	return graph{obstacles: obstacleSet}
}

func (g graph) search(start, goal cell) searchResult {
	queue := []cell{start}
	visited := make(map[string]bool)
	enqueued := make(map[string]bool)
	enqueued[start.key()] = true
	parent := make(map[string]cell)
	steps := []searchStep{}
	visitedOrder := []cell{}
	found := false
	count := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.key()] {
			continue
		}

		visited[current.key()] = true
		visitedOrder = append(visitedOrder, current)
		count++

		steps = append(steps, searchStep{
			number:  count,
			current: current,
			queue:   append([]cell(nil), queue...),
			visited: append([]cell(nil), visitedOrder...),
		})

		if current == goal {
			found = true
			break
		}

		for _, dir := range directions {
			neighbor := cell{row: current.row + dir.row, col: current.col + dir.col}
			key := neighbor.key()
			if !g.inBounds(neighbor) || g.isObstacle(neighbor) {
				continue
			}
			if visited[key] || enqueued[key] {
				continue
			}
			parent[key] = current
			queue = append(queue, neighbor)
			enqueued[key] = true
		}
	}

	var path []cell
	if found {
		path = buildPath(parent, start, goal)
	}

	return searchResult{steps: steps, path: path, found: found}
}

func buildPath(parent map[string]cell, start, goal cell) []cell {
	path := []cell{}
	current := goal

	for {
		path = append([]cell{current}, path...)
		if current == start {
			break
		}
		prev, ok := parent[current.key()]
		if !ok {
			break
		}
		current = prev
	}

	return path
}

func (g graph) inBounds(c cell) bool {
	return c.row >= 0 && c.row < rows && c.col >= 0 && c.col < cols
}

func (g graph) isObstacle(c cell) bool {
	return g.obstacles[c.key()]
}

func (g graph) renderGrid(step searchStep, path []cell, highlightPath bool) string {
	visitedSet := make(map[string]bool, len(step.visited))
	for _, v := range step.visited {
		visitedSet[v.key()] = true
	}
	queueSet := make(map[string]bool, len(step.queue))
	for _, q := range step.queue {
		queueSet[q.key()] = true
	}
	pathSet := make(map[string]bool)
	if highlightPath {
		for _, p := range path {
			pathSet[p.key()] = true
		}
	}

	var builder strings.Builder
	builder.WriteString("    ")
	for c := 0; c < cols; c++ {
		builder.WriteString(fmt.Sprintf("%3d", c))
	}
	builder.WriteByte('\n')

	for r := 0; r < rows; r++ {
		builder.WriteString(fmt.Sprintf("%3d ", r))
		for c := 0; c < cols; c++ {
			cell := cell{row: r, col: c}
			builder.WriteString(g.cellSymbol(cell, step.current, visitedSet, queueSet, pathSet))
		}
		builder.WriteByte('\n')
	}

	builder.WriteString("Legend: S=start G=goal #=obstacle @=текущий *=путь .=посещён o=очередь\n")
	return builder.String()
}

func (g graph) cellSymbol(c, current cell, visited, queued, path map[string]bool) string {
	switch {
	case g.isObstacle(c):
		return " # "
	case c == startCell:
		return " S "
	case c == goalCell:
		return " G "
	case path[c.key()]:
		return " * "
	case c == current:
		return " @ "
	case queued[c.key()]:
		return " o "
	case visited[c.key()]:
		return " . "
	default:
		return "   "
	}
}

func formatCellList(cells []cell, sep string, limit int) string {
	if len(cells) == 0 {
		return "(пусто)"
	}
	display := cells
	suffix := ""
	if limit > 0 && len(cells) > limit {
		display = cells[:limit]
		suffix = " -> …"
	}
	names := make([]string, 0, len(display))
	for _, c := range display {
		names = append(names, c.String())
	}
	return strings.Join(names, sep) + suffix
}
