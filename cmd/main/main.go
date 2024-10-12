package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var mainApp fyne.App

func main() {
	mainApp = app.New()
	mainWindow := mainApp.NewWindow("mainWindow")

	textGrid := widget.NewTextGrid()
	textGrid.SetText("hello world hello\n hehe hello")
	mainWindow.SetContent(container.NewVBox(
		widget.NewLabel("hello world"),
		textGrid,
	))
	go closeWindow()
	fmt.Println("running main window...")
	mainWindow.ShowAndRun()

}

func closeWindow() {
	time.Sleep(4 * time.Second)
	fmt.Println("closing main window")
	mainApp.Quit()
}
