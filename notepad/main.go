package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notepad/note"
	"example.com/notepad/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getTodoData()
	tod, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	tod.Display()
	err = tod.Save()
	if err != nil {
		fmt.Println("Saving the TODO failed.")
		return
	}
	fmt.Println("Saving the TODO succeeded!")
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}
func getTodoData() (content string) {
	content = getUserInput("Todo content:")
	return
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
