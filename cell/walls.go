package cell

import "fmt"

type Wall struct {
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
				// consider only top wall
				w := topWall(absOrigin, sub)
				walls = append(walls, w)
			} else if isSplitVer(sub){
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
	translated := absOrigin.Add(sub.RelativePos)
 	start := Pos{
 		X: translated.X,
		Y: translated.Y,
	}
	end := Pos{
		X: translated.X + sub.Cell.Dim.Width,
		Y: translated.Y,
	}
	return Wall{
		Start: start,
		End: end,
	}
}

func leftWall(absOrigin Pos, sub Subcell) Wall {
	translated := absOrigin.Add(sub.RelativePos)
	start := Pos{
		X: translated.X,
		Y: translated.Y,
	}
	end := Pos{
		X: translated.X,
		Y: translated.Y + sub.Cell.Dim.Height,
	}
	return Wall{
		Start: start,
		End: end,
	}
}
