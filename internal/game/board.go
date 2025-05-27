package game

type Board struct {
	Width  int
	Height int
	Cells  [][]int `json:"cells"`
}

// Boardをディープコピーして返す
func (b Board) Copy() *Board {
	height, width := len(b.Cells), len(b.Cells[0])
	newCells := make([][]int, height)
	for i := range newCells {
		newCells[i] = make([]int, width)
		copy(newCells[i], b.Cells[i])
	}
	return &Board{Cells: newCells}
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
	return &Board{Cells: newCells}
}
