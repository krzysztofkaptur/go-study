package main

import (
	"bufio"
	"fmt"
	"note/note"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Note title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Note content: ")
	content, _ := reader.ReadString('\n')

	newNote := note.New(title, content)

	msg, err := newNote.SaveToFile()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(msg)
	}
}