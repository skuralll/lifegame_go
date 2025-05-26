package entities

type Board struct {
	Width  int   `json:"width"`
	Height int   `json:"height"`
	Cells  []int `json:"cells"`
}
