package utils

import (
	term "github.com/buger/goterm"

	"fmt"
	"strings"
	"time"
)

func clearAndPrint(str string) {
	term.Clear()
	term.MoveCursor(1, 1)
	term.Print(str)
	term.Flush()
}

func Spinner(col int) {
	// Clear the screen by printing \x0c.
	bar := fmt.Sprintf("[%%-%vs]", col)
	for i := 0; i < col; i++ {
		msg := fmt.Sprintf(bar, strings.Repeat("=", i)+">")
		clearAndPrint(msg)
		time.Sleep(100 * time.Millisecond)
	}
	clearAndPrint("Done!! \n")
}
