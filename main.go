package main

import (
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {

	game := Game{
		currentScore: NewScore(),
		highScore:    NewHighScore(),
		exitButton:   NewExitButton(),
		title:        NewTitle(),
		arena:        NewArena(),
		guide:        tview.NewBox().SetBorder(true).SetTitle(" How to Play "),
	}
	game.resetButton = NewResetButton(game.Reset)

	app.SetInputCapture(game.UserInput)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.currentScore, 7, 0, false).
			AddItem(game.highScore, 7, 0, false).
			AddItem(game.resetButton, 7, 0, false).
			AddItem(game.exitButton, 7, 0, false),
			16, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.title, 6, 1, false).
			AddItem(game.arena, gridHeight*tileHeight+4, 0, false).
			AddItem(tview.NewBox(), 0, 1, false),
			gridWidth*tileWidth+4, 0, false).
		AddItem(game.guide, 20, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
