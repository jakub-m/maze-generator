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

func (m *Maze) ForEachWall(fn func(a, b Pos)) {
	for ix := 0; ix < m.Width+1; ix++ {
		for iy := 0; iy < m.Height+1; iy++ {
			cRef := Cell{X: ix, Y: iy}

			cEast := Cell{X: ix + 1, Y: iy}
			if !m.passages.hasPassage(cRef, cEast) {
				fn(Pos(cRef), Pos(cEast))
			}

			cSouth := Cell{X: ix, Y: iy + 1}
			if !m.passages.hasPassage(cRef, cSouth) {
				fn(Pos(cRef), Pos(cSouth))
			}
		}
	}
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
	pair1 := cellPair{a, b}
	pair2 := cellPair{b, a}
	if _, ok := p[pair1]; ok {
		return false
	}
	if _, ok := p[pair2]; ok {
		return false
	}
	p[pair1] = true
	p[pair2] = true
	return true
}

func (p passageMap) hasPassage(a, b Cell) bool {
	_, ok := p[cellPair{a, b}]
	return ok
}

type cellPair struct {
	a, b Cell
}

type Cell Pos

func (c Cell) String() string {
	return fmt.Sprintf("[%d, %d]", c.X, c.Y)
}

type Pos struct{ X, Y int }

func (c Pos) String() string {
	return fmt.Sprintf("[%d, %d]", c.X, c.Y)
}
