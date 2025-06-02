package game

type Board struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Cells  [][]int `json:"cells"`
}

// Boardをディープコピーして返す
func (b *Board) Copy() *Board {
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

// 指定したセルの周囲の生存セル数を取得
func (b *Board) coutAroundLives(x int, y int) int {
	cnt := 0
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			cx := x + dx
			cy := y + dy
			if 0 <= cx && cx < b.Width && 0 <= cy && cy < b.Height && b.Cells[cy][cx] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

// 次世代の盤面を取得する関数
func (b *Board) GetNext() *Board {
	newBoard := b.Copy()
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			cnt := b.coutAroundLives(x, y)
			// 誕生
			if b.Cells[y][x] == 0 && cnt == 3 {
				newBoard.Cells[y][x] = 1
			}
			// 生存
			if b.Cells[y][x] == 1 && (cnt == 2 || cnt == 3) {
				newBoard.Cells[y][x] = 1
			}
			// 過疎
			if b.Cells[y][x] == 1 && cnt <= 1 {
				newBoard.Cells[y][x] = 0
			}
			// 過密
			if b.Cells[y][x] == 1 && 4 <= cnt {
				newBoard.Cells[y][x] = 0
			}
		}
	}
	return newBoard
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
