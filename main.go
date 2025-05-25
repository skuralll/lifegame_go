package main

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

const (
	width  = 10
	height = 10
)

type cells [width][height]bool

var cs, ncs cells

func main() {
	initialize()
	clearScreen()
	render()
}

// 盤面初期化
func initialize() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cs[y][x] = rand.Intn(2) == 1
		}
	}
}

// 盤面描画
func render() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if cs[y][x] {
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
