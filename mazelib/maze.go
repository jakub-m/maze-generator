package mazelib

import (
	"fmt"
	"log"
	"math/rand"
)

type Maze struct {
	Width    int // x: [0, width)
	Height   int // y: [0, height)
	visited  visitedMap
	passages passageMap
}

func NewMaze(width, height int) (*Maze, error) {
	m := &Maze{
		Width:    width,
		Height:   height,
		visited:  make(map[Cell]bool),
		passages: make(map[cellPair]bool),
	}
	trace := []Cell{m.chooseStartCell()}
	for len(trace) > 0 {
		log.Printf("trace len: %d, visited %d, passages: %d", len(trace), len(m.visited), len(m.passages)/2)
		current := trace[len(trace)-1]
		log.Printf("trace: %v", trace)
		m.visited.mark(current)
		candidates := m.unvisitedNeighbours(current)
		if next, ok := chooseCell(candidates); ok {
			if ok := m.passages.put(current, next); !ok {
				return nil, fmt.Errorf("BUG. Passage already exists: %s -> %s", current, next)
			}
			trace = append(trace, next)
		} else {
			n := len(trace)
			trace = trace[0 : n-1]
		}
	}
	if len(m.visited) != height*width {
		return m, fmt.Errorf("visited %d but should %d", len(m.visited), height*width)
	}
	return m, nil

}

func (m *Maze) chooseStartCell() Cell {
	return Cell{
		X: 0,
		Y: rand.Intn(m.Height),
	}
}

func (m *Maze) chooseEndCell() Cell {
	return Cell{
		X: m.Width - 1,
		Y: rand.Intn(m.Height),
	}
}

func (m *Maze) unvisitedNeighbours(c Cell) []Cell {
	var unvisited []Cell
	appendIfUnvisited := func(x, y int) {
		c := Cell{X: x, Y: y}
		if _, wasVisited := m.visited[c]; !wasVisited {
			unvisited = append(unvisited, c)
		}
	}
	if c.X > 0 {
		appendIfUnvisited(c.X-1, c.Y)
	}
	if c.X < m.Width-1 {
		appendIfUnvisited(c.X+1, c.Y)
	}
	if c.Y > 0 {
		appendIfUnvisited(c.X, c.Y-1)
	}
	if c.Y < m.Height-1 {
		appendIfUnvisited(c.X, c.Y+1)
	}
	return unvisited
}

func (m *Maze) HasPassage(a, b Cell) bool {
	_, ok := m.passages[newCellPair(a, b)]
	return ok
}

func chooseCell(cells []Cell) (chosen Cell, ok bool) {
	if len(cells) == 0 {
		return Cell{}, false
	}
	return cells[rand.Intn(len(cells))], true
}

type visitedMap map[Cell]bool

func (v visitedMap) mark(c Cell) {
	v[c] = true
}

type passageMap map[cellPair]bool

func (p passageMap) put(a, b Cell) (ok bool) {
	pair := newCellPair(a, b)
	if _, ok := p[pair]; ok {
		return false
	}
	p[pair] = true
	return true
}

func (p passageMap) hasPassage(a, b Cell) bool {
	_, ok := p[newCellPair(a, b)]
	return ok
}

type cellPair struct {
	a, b Cell
}

func newCellPair(a, b Cell) cellPair {
	if a.X < b.X {
		return cellPair{a, b}
	}
	if b.X < a.X {
		return cellPair{b, a}
	}
	if a.Y < b.Y {
		return cellPair{a, b}
	}
	if b.Y < a.Y {
		return cellPair{b, a}
	}
	return cellPair{a, b}
}

type Cell struct{ X, Y int }

func (c Cell) String() string {
	return fmt.Sprintf("[%d, %d]", c.X, c.Y)
}
