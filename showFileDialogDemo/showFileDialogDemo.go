package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/dialog"
)

func main() {
	app := app.New()

	win := app.NewWindow("Show File Dialog Demo")
	win.SetContent(widget.NewVBox(
		widget.NewButton("Browse", func() {
			dialog.ShowFileOpen(func(file fyne.FileReadCloser, err error) {
			}, win)
		}),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	size := fyne.NewSize(500, 500)
	win.Resize(size)
	win.ShowAndRun()
}

