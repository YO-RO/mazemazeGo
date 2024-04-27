package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	mmg := NewMazeMazeGo(GetMazeDesign())
GAME_LOOP:
	for {
		fmt.Println(mmg)

		if mmg.IsEscaped() {
			fmt.Println("Successful Escape!")
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
