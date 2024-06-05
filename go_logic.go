package main

type cell struct {
	x int
	y int
}

func isneibors(cell1 cell, cell2 cell) int {
	if (cell1.x == cell2.x && cell1.y+1 == cell2.y) || (cell1.x+1 == cell2.x && cell1.y == cell2.y) ||
		(cell1.x+1 == cell2.x && cell1.y+1 == cell2.y) || (cell1.x == cell2.x && cell1.y-1 == cell2.y) ||
		(cell1.x-1 == cell2.x && cell1.y == cell2.y) || (cell1.x-1 == cell2.x && cell1.y-1 == cell2.y) ||
		(cell1.x+1 == cell2.x && cell1.y-1 == cell2.y) || (cell1.x-1 == cell2.x && cell1.y+1 == cell2.y) {
		return 1
	}
	return 0
}

func getNeibors(curCell cell) []cell {
	return []cell{
		cell{x: curCell.x, y: curCell.y + 1},
		cell{x: curCell.x + 1, y: curCell.y},
		cell{x: curCell.x + 1, y: curCell.y + 1},
		cell{x: curCell.x, y: curCell.y - 1},
		cell{x: curCell.x - 1, y: curCell.y},
		cell{x: curCell.x - 1, y: curCell.y - 1},
		cell{x: curCell.x - 1, y: curCell.y + 1},
		cell{x: curCell.x + 1, y: curCell.y - 1},
	}
}

func getAliveCells(grid []cell, curCell cell, neibors []cell) []cell {
	if len(neibors) == 0 {
		return nil
	}

	if !indexOf(grid, neibors[0]) {
		return getAliveCells(grid, curCell, neibors[1:])
	} else {
		return append(getAliveCells(grid, curCell, neibors[1:]), neibors[0])
	}
}

func removeDuplicates_(cels []cell) []cell {
	if len(cels) < 2 {
		return cels
	}

	if len(countIn(cels, cels[0])) == 1 {
		return append(removeDuplicates_(cels[1:]), cels[0])
	} else {
		return removeDuplicates_(cels[1:])
	}

}

func countIn(grid []cell, curCell cell) []cell {
	if len(grid) == 0 {
		return nil
	}
	if len(grid) == 1 && grid[0] == curCell {
		return []cell{curCell}
	}

	if curCell == grid[0] {
		return append(countIn(grid[1:], curCell), curCell)
	} else {
		return countIn(grid[1:], curCell)
	}
}

func getAllDeadCells(grid []cell, cels []cell) []cell {

	if len(cels) == 0 {
		return nil
	}

	return append(getDeadCells(grid, getNeibors(cels[0])), getAllDeadCells(grid, cels[1:])...)
}

func getDeadCells(grid []cell, neibors []cell) []cell {
	if len(neibors) == 0 {
		return nil
	}
	if indexOf(grid, neibors[0]) {
		return getDeadCells(grid, neibors[1:])
	} else {
		return append(getDeadCells(grid, neibors[1:]), neibors[0])
	}
}

func indexOf(grid []cell, curCell cell) bool {
	for _, element := range grid {
		if element.x == curCell.x && element.y == curCell.y {
			return true
		}
	}

	return false
}

func getAllSurv(grid []cell, cellc []cell) []cell {
	if len(cellc) == 0 {
		return nil
	}

	if len(getAliveCells(grid, cellc[0], getNeibors(cellc[0]))) < 2 || len(getAliveCells(grid, cellc[0], getNeibors(cellc[0]))) > 3 {
		return getAllSurv(grid, cellc[1:])
	} else {
		return append(getAllSurv(grid, cellc[1:]), cellc[0])
	}
}

func getAllNewCells(grid []cell, deadCells []cell) []cell {

	if len(deadCells) == 0 {
		return nil
	}

	if len(getAliveCells(grid, deadCells[0], getNeibors(deadCells[0]))) != 3 {
		return getAllNewCells(grid, deadCells[1:])
	} else {
		return append(getAllNewCells(grid, deadCells[1:]), deadCells[0])
	}
}

func updateField(grid []cell) []cell {
	return append(getAllSurv(grid, grid), getAllNewCells(grid, (getAllDeadCells(grid, grid)))...)
}
