package main

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	width  = 10
	height = 10
	rate   = 200 // 更新頻度(ミリ秒)
)

type cells [width][height]bool

var cs, ncs cells

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
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cs[y][x] = rand.Intn(2) == 1
		}
	}
}

// 盤面更新
func update() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cs[y][x] = rand.Intn(2) == 1
		}
	}
}

// 指定したセルの周囲の生存セル数を取得
func coutAroundLives(x int, y int) int {
	cnt := 0
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			cx := x + dx
			cy := y + dy
			if 0 <= cx && cx < width && 0 <= cy && cy < height && cs[cy][cx] {
				cnt++
			}
		}
	}
	return cnt
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
