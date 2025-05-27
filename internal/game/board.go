package game

type Board struct {
	Width  int
	Height int
	Cells  [][]int `json:"cells"`
}

// Boardをディープコピーして返す
func (b Board) Copy() *Board {
	newCells := make([][]int, b.Height)
	for i := range newCells {
		newCells[i] = make([]int, b.Width)
		copy(newCells[i], b.Cells[i])
	}
	return &Board{
		Width:  b.Width,
		Height: b.Height,
		Cells:  newCells,
	}
}

// 初期化された盤面を取得する(全ての要素が0)
func NewBoard(width int, height int) *Board {
	newCells := make([][]int, height)
	for i := range newCells {
		row := make([]int, width)
		for j := range row {
			row[j] = 0
		}
		newCells[i] = row
	}
	return &Board{
		Width:  width,
		Height: height,
		Cells:  newCells,
	}
}
