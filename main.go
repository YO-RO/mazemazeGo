package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mattn/go-tty"
)

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	mmg := NewMazeMazeGo(GetMaze())
	printedLine := 0
GAME_LOOP:
	for {
		// 前回の表示内容を消去してからmmgを表示する
		if printedLine > 0 {
			// 表示内容を消去する
			// \033[nF -> カーソルをn行上の先頭に移動
			//	(nが0の時、1行上の先頭に移動してしまう)
			// \033[0J -> カーソルより後ろの画面をクリア
			fmt.Printf("\033[%dF\033[0J", printedLine)
		}
		fmt.Print(mmg)
		printedLine = strings.Count(mmg.String(), "\n")

		if mmg.GameIsEnd() {
			return
		}

		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		// handle input
		switch r {
		case 'x':
			break GAME_LOOP
		case 'h':
			mmg.ToggleCorrectRouteDisplay()
		case 'w':
			mmg.MovePlayer(Up)
		case 's':
			mmg.MovePlayer(Down)
		case 'a':
			mmg.MovePlayer(Left)
		case 'd':
			mmg.MovePlayer(Right)
		}
	}
}
