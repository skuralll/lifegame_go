package main

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com.skuralll/lifegame_go/internal/game"
)

const (
	width  = 20
	height = 10
	rate   = 200 // 更新頻度(ミリ秒)
)

var board *game.Board

func main() {
	initialize()
	for {
		clearScreen()
		render()
		time.Sleep(rate * time.Millisecond)
		update()
	}
}

// 盤面初期化
func initialize() {
	board = game.NewBoard(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			board.Cells[y][x] = rand.Intn(2)
		}
	}
}

// 盤面更新
func update() {
	board = board.GetNext()
}

// 盤面描画
func render() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board.Cells[y][x] == 1 {
				print("⬛︎")
			} else {
				print("⬜︎")
			}
		}
		println("")
	}
}

// 画面クリア
func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // Windows用
	default:
		cmd = exec.Command("clear") // Unix系 (Linux, macOS) 用
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
