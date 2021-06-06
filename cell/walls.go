package cell

import "fmt"

type Wall struct {
	Parts []Line
}

type Line struct {
	Start, End Pos
}

// InternalWalls returns non-overlapping set of walls with absolute coordinates.
func InternalWalls(c Cell) []Wall {
	zero := Pos{0, 0}
	return internalWallsRec(zero, c)
}

func internalWallsRec(absOrigin Pos, c Cell) []Wall {
	var walls []Wall
	for _, sub := range c.Subcells {
		if !isFirstSub(sub) {
			if isSplitHor(sub) {
				w := topWall(absOrigin, sub)
				walls = append(walls, w)
			} else if isSplitVer(sub) {
				w := leftWall(absOrigin, sub)
				walls = append(walls, w)
			} else {
				panic(fmt.Sprintf("BUG. Unreachable state %v", sub))
			}
		}
		recWalls := internalWallsRec(absOrigin.Add(sub.RelativePos), sub.Cell)
		walls = append(walls, recWalls...)
	}
	return walls
}

func isFirstSub(sub Subcell) bool {
	return sub.RelativePos.X == 0 && sub.RelativePos.Y == 0
}

// isSplitHor tells if the subcell was created after horizontal split
func isSplitHor(sub Subcell) bool {
	return sub.RelativePos.X == 0 && sub.RelativePos.Y != 0
}

// isSplitVer tells if the subcell was created after vertical split
func isSplitVer(sub Subcell) bool {
	return sub.RelativePos.X != 0 && sub.RelativePos.Y == 0
}

func topWall(absOrigin Pos, sub Subcell) Wall {
	parts := []Line{
		{
			Start: Pos{0, 0},
			End: Pos{
				X: sub.Cell.Dim.Width,
				Y: 0,
			},
		},
	}
	translation := absOrigin.Add(sub.RelativePos)
	translateLines(translation, parts)
	return Wall{parts}
}

func leftWall(absOrigin Pos, sub Subcell) Wall {
	parts := []Line{
		{
			Start: Pos{0, 0},
			End: Pos{
				X: 0,
				Y: sub.Cell.Dim.Height,
			},
		},
	}
	translation := absOrigin.Add(sub.RelativePos)
	translateLines(translation, parts)
	return Wall{parts}
}

func translateLines(translation Pos, lines []Line) {
	for i, l := range lines {
		lines[i] = Line{
			Start: l.Start.Add(translation),
			End:   l.End.Add(translation),
		}
	}
}
